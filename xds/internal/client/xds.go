/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package client

import (
	"errors"
	"fmt"
	"net"
	"strconv"

	v1typepb "github.com/cncf/udpa/go/udpa/type/v1"
	v3clusterpb "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	v3corepb "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3endpointpb "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	v3listenerpb "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	v3routepb "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	v3httppb "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	v3tlspb "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	v3typepb "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"

	"google.golang.org/grpc/internal/grpclog"
	xdsinternal "google.golang.org/grpc/xds/internal"
	"google.golang.org/grpc/xds/internal/env"
	"google.golang.org/grpc/xds/internal/httpfilter"
	"google.golang.org/grpc/xds/internal/version"
)

// TransportSocket proto message has a `name` field which is expected to be set
// to this value by the management server.
const transportSocketName = "envoy.transport_sockets.tls"

// UnmarshalListener processes resources received in an LDS response, validates
// them, and transforms them into a native struct which contains only fields we
// are interested in.
func UnmarshalListener(_ string, resources []*anypb.Any, logger *grpclog.PrefixLogger) (map[string]ListenerUpdate, UpdateMetadata, error) {
	update := make(map[string]ListenerUpdate)
	for _, r := range resources {
		if !IsListenerResource(r.GetTypeUrl()) {
			return nil, UpdateMetadata{}, fmt.Errorf("xds: unexpected resource type: %q in LDS response", r.GetTypeUrl())
		}
		// TODO: Pass version.TransportAPI instead of relying upon the type URL
		v2 := r.GetTypeUrl() == version.V2ListenerURL
		lis := &v3listenerpb.Listener{}
		if err := proto.Unmarshal(r.GetValue(), lis); err != nil {
			return nil, UpdateMetadata{}, fmt.Errorf("xds: failed to unmarshal resource in LDS response: %v", err)
		}
		logger.Infof("Resource with name: %v, type: %T, contains: %v", lis.GetName(), lis, lis)

		lu, err := processListener(lis, v2)
		if err != nil {
			return nil, UpdateMetadata{}, err
		}
		update[lis.GetName()] = *lu
	}
	return update, UpdateMetadata{}, nil
}

func processListener(lis *v3listenerpb.Listener, v2 bool) (*ListenerUpdate, error) {
	if lis.GetApiListener() != nil {
		return processClientSideListener(lis, v2)
	}
	return processServerSideListener(lis)
}

// processClientSideListener checks if the provided Listener proto meets
// the expected criteria. If so, it returns a non-empty routeConfigName.
func processClientSideListener(lis *v3listenerpb.Listener, v2 bool) (*ListenerUpdate, error) {
	update := &ListenerUpdate{}

	apiLisAny := lis.GetApiListener().GetApiListener()
	if !IsHTTPConnManagerResource(apiLisAny.GetTypeUrl()) {
		return nil, fmt.Errorf("xds: unexpected resource type: %q in LDS response", apiLisAny.GetTypeUrl())
	}
	apiLis := &v3httppb.HttpConnectionManager{}
	if err := proto.Unmarshal(apiLisAny.GetValue(), apiLis); err != nil {
		return nil, fmt.Errorf("xds: failed to unmarshal api_listner in LDS response: %v", err)
	}

	switch apiLis.RouteSpecifier.(type) {
	case *v3httppb.HttpConnectionManager_Rds:
		if apiLis.GetRds().GetConfigSource().GetAds() == nil {
			return nil, fmt.Errorf("xds: ConfigSource is not ADS in LDS response: %+v", lis)
		}
		name := apiLis.GetRds().GetRouteConfigName()
		if name == "" {
			return nil, fmt.Errorf("xds: empty route_config_name in LDS response: %+v", lis)
		}
		update.RouteConfigName = name
	case *v3httppb.HttpConnectionManager_RouteConfig:
		// TODO: Add support for specifying the RouteConfiguration inline
		// in the LDS response.
		return nil, fmt.Errorf("xds: LDS response contains RDS config inline. Not supported for now: %+v", apiLis)
	case nil:
		return nil, fmt.Errorf("xds: no RouteSpecifier in received LDS response: %+v", apiLis)
	default:
		return nil, fmt.Errorf("xds: unsupported type %T for RouteSpecifier in received LDS response", apiLis.RouteSpecifier)
	}

	if v2 {
		return update, nil
	}

	// The following checks and fields only apply to xDS protocol versions v3+.

	update.MaxStreamDuration = apiLis.GetCommonHttpProtocolOptions().GetMaxStreamDuration().AsDuration()

	var err error
	if update.HTTPFilters, err = processHTTPFilters(apiLis.GetHttpFilters(), false); err != nil {
		return nil, fmt.Errorf("xds: %v", err)
	}

	return update, nil
}

func unwrapHTTPFilterConfig(config *anypb.Any) (proto.Message, string, error) {
	if typeURL := config.GetTypeUrl(); typeURL != "type.googleapis.com/udpa.type.v1.TypedStruct" {
		return config, typeURL, nil
	}
	// The real type name is inside the TypedStruct.
	s := new(v1typepb.TypedStruct)
	if err := ptypes.UnmarshalAny(config, s); err != nil {
		return nil, "", fmt.Errorf("error unmarshalling TypedStruct filter config: %v", err)
	}
	return s, s.GetTypeUrl(), nil
}

func validateHTTPFilterConfig(cfg *anypb.Any, lds bool) (httpfilter.Filter, httpfilter.FilterConfig, error) {
	config, typeURL, err := unwrapHTTPFilterConfig(cfg)
	if err != nil {
		return nil, nil, err
	}
	filterBuilder := httpfilter.Get(typeURL)
	if filterBuilder == nil {
		return nil, nil, fmt.Errorf("no filter implementation found for %q", typeURL)
	}
	parseFunc := filterBuilder.ParseFilterConfig
	if !lds {
		parseFunc = filterBuilder.ParseFilterConfigOverride
	}
	filterConfig, err := parseFunc(config)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing config for filter %q: %v", typeURL, err)
	}
	return filterBuilder, filterConfig, nil
}

func processHTTPFilterOverrides(cfgs map[string]*anypb.Any) (map[string]httpfilter.FilterConfig, error) {
	if !env.FaultInjectionSupport || len(cfgs) == 0 {
		return nil, nil
	}
	m := make(map[string]httpfilter.FilterConfig)
	for name, cfg := range cfgs {
		_, config, err := validateHTTPFilterConfig(cfg, false)
		if err != nil {
			return nil, err
		}
		m[name] = config
	}
	return m, nil
}

func processHTTPFilters(filters []*v3httppb.HttpFilter, server bool) ([]HTTPFilter, error) {
	if !env.FaultInjectionSupport {
		return nil, nil
	}

	ret := make([]HTTPFilter, 0, len(filters))
	seenNames := make(map[string]bool, len(filters))
	for _, filter := range filters {
		name := filter.GetName()
		if name == "" {
			return nil, errors.New("filter missing name field")
		}
		if seenNames[name] {
			return nil, fmt.Errorf("duplicate filter name %q", name)
		}
		seenNames[name] = true

		httpFilter, config, err := validateHTTPFilterConfig(filter.GetTypedConfig(), true)
		if err != nil {
			return nil, err
		}
		if server {
			if _, ok := httpFilter.(httpfilter.ServerInterceptorBuilder); !ok {
				return nil, fmt.Errorf("httpFilter %q not supported server-side", name)
			}
		} else {
			if _, ok := httpFilter.(httpfilter.ClientInterceptorBuilder); !ok {
				return nil, fmt.Errorf("httpFilter %q not supported client-side", name)
			}
		}

		// Save name/config
		ret = append(ret, HTTPFilter{Name: name, Filter: httpFilter, Config: config})
	}
	return ret, nil
}

func processServerSideListener(lis *v3listenerpb.Listener) (*ListenerUpdate, error) {
	addr := lis.GetAddress()
	if addr == nil {
		return nil, fmt.Errorf("xds: no address field in LDS response: %+v", lis)
	}
	sockAddr := addr.GetSocketAddress()
	if sockAddr == nil {
		return nil, fmt.Errorf("xds: no socket_address field in LDS response: %+v", lis)
	}
	lu := &ListenerUpdate{
		InboundListenerCfg: &InboundListenerConfig{
			Address: sockAddr.GetAddress(),
			Port:    strconv.Itoa(int(sockAddr.GetPortValue())),
		},
	}

	// Make sure the listener resource contains a single filter chain. We do not
	// support multiple filter chains and picking the best match from the list.
	fcs := lis.GetFilterChains()
	if n := len(fcs); n != 1 {
		return nil, fmt.Errorf("xds: filter chains count in LDS response does not match expected. Got %d, want 1", n)
	}
	fc := fcs[0]

	// If the transport_socket field is not specified, it means that the control
	// plane has not sent us any security config. This is fine and the server
	// will use the fallback credentials configured as part of the
	// xdsCredentials.
	ts := fc.GetTransportSocket()
	if ts == nil {
		return lu, nil
	}
	if name := ts.GetName(); name != transportSocketName {
		return nil, fmt.Errorf("xds: transport_socket field has unexpected name: %s", name)
	}
	any := ts.GetTypedConfig()
	if any == nil || any.TypeUrl != version.V3DownstreamTLSContextURL {
		return nil, fmt.Errorf("xds: transport_socket field has unexpected typeURL: %s", any.TypeUrl)
	}
	downstreamCtx := &v3tlspb.DownstreamTlsContext{}
	if err := proto.Unmarshal(any.GetValue(), downstreamCtx); err != nil {
		return nil, fmt.Errorf("xds: failed to unmarshal DownstreamTlsContext in LDS response: %v", err)
	}
	if downstreamCtx.GetCommonTlsContext() == nil {
		return nil, errors.New("xds: DownstreamTlsContext in LDS response does not contain a CommonTlsContext")
	}
	sc, err := securityConfigFromCommonTLSContext(downstreamCtx.GetCommonTlsContext())
	if err != nil {
		return nil, err
	}
	if sc.IdentityInstanceName == "" {
		return nil, errors.New("security configuration on the server-side does not contain identity certificate provider instance name")
	}
	sc.RequireClientCert = downstreamCtx.GetRequireClientCertificate().GetValue()
	if sc.RequireClientCert && sc.RootInstanceName == "" {
		return nil, errors.New("security configuration on the server-side does not contain root certificate provider instance name, but require_client_cert field is set")
	}
	lu.SecurityCfg = sc
	return lu, nil
}

// UnmarshalRouteConfig processes resources received in an RDS response,
// validates them, and transforms them into a native struct which contains only
// fields we are interested in. The provided hostname determines the route
// configuration resources of interest.
func UnmarshalRouteConfig(_ string, resources []*anypb.Any, logger *grpclog.PrefixLogger) (map[string]RouteConfigUpdate, UpdateMetadata, error) {
	update := make(map[string]RouteConfigUpdate)
	for _, r := range resources {
		if !IsRouteConfigResource(r.GetTypeUrl()) {
			return nil, UpdateMetadata{}, fmt.Errorf("xds: unexpected resource type: %q in RDS response", r.GetTypeUrl())
		}
		rc := &v3routepb.RouteConfiguration{}
		if err := proto.Unmarshal(r.GetValue(), rc); err != nil {
			return nil, UpdateMetadata{}, fmt.Errorf("xds: failed to unmarshal resource in RDS response: %v", err)
		}
		logger.Infof("Resource with name: %v, type: %T, contains: %v.", rc.GetName(), rc, rc)

		// TODO: Pass version.TransportAPI instead of relying upon the type URL
		v2 := r.GetTypeUrl() == version.V2RouteConfigURL
		// Use the hostname (resourceName for LDS) to find the routes.
		u, err := generateRDSUpdateFromRouteConfiguration(rc, logger, v2)
		if err != nil {
			return nil, UpdateMetadata{}, fmt.Errorf("xds: received invalid RouteConfiguration in RDS response: %+v with err: %v", rc, err)
		}
		update[rc.GetName()] = u
	}
	return update, UpdateMetadata{}, nil
}

// generateRDSUpdateFromRouteConfiguration checks if the provided
// RouteConfiguration meets the expected criteria. If so, it returns a
// RouteConfigUpdate with nil error.
//
// A RouteConfiguration resource is considered valid when only if it contains a
// VirtualHost whose domain field matches the server name from the URI passed
// to the gRPC channel, and it contains a clusterName or a weighted cluster.
//
// The RouteConfiguration includes a list of VirtualHosts, which may have zero
// or more elements. We are interested in the element whose domains field
// matches the server name specified in the "xds:" URI. The only field in the
// VirtualHost proto that the we are interested in is the list of routes. We
// only look at the last route in the list (the default route), whose match
// field must be empty and whose route field must be set.  Inside that route
// message, the cluster field will contain the clusterName or weighted clusters
// we are looking for.
func generateRDSUpdateFromRouteConfiguration(rc *v3routepb.RouteConfiguration, logger *grpclog.PrefixLogger, v2 bool) (RouteConfigUpdate, error) {
	var vhs []*VirtualHost
	for _, vh := range rc.GetVirtualHosts() {
		routes, err := routesProtoToSlice(vh.Routes, logger, v2)
		if err != nil {
			return RouteConfigUpdate{}, fmt.Errorf("received route is invalid: %v", err)
		}
		vhOut := &VirtualHost{
			Domains: vh.GetDomains(),
			Routes:  routes,
		}
		if !v2 {
			cfgs, err := processHTTPFilterOverrides(vh.GetTypedPerFilterConfig())
			if err != nil {
				return RouteConfigUpdate{}, fmt.Errorf("virtual host %+v: %v", vh, err)
			}
			vhOut.HTTPFilterConfigOverride = cfgs
		}
		vhs = append(vhs, vhOut)
	}
	return RouteConfigUpdate{VirtualHosts: vhs}, nil
}

func routesProtoToSlice(routes []*v3routepb.Route, logger *grpclog.PrefixLogger, v2 bool) ([]*Route, error) {
	var routesRet []*Route

	for _, r := range routes {
		match := r.GetMatch()
		if match == nil {
			return nil, fmt.Errorf("route %+v doesn't have a match", r)
		}

		if len(match.GetQueryParameters()) != 0 {
			// Ignore route with query parameters.
			logger.Warningf("route %+v has query parameter matchers, the route will be ignored", r)
			continue
		}

		pathSp := match.GetPathSpecifier()
		if pathSp == nil {
			return nil, fmt.Errorf("route %+v doesn't have a path specifier", r)
		}

		var route Route
		switch pt := pathSp.(type) {
		case *v3routepb.RouteMatch_Prefix:
			route.Prefix = &pt.Prefix
		case *v3routepb.RouteMatch_Path:
			route.Path = &pt.Path
		case *v3routepb.RouteMatch_SafeRegex:
			route.Regex = &pt.SafeRegex.Regex
		default:
			return nil, fmt.Errorf("route %+v has an unrecognized path specifier: %+v", r, pt)
		}

		if caseSensitive := match.GetCaseSensitive(); caseSensitive != nil {
			route.CaseInsensitive = !caseSensitive.Value
		}

		for _, h := range match.GetHeaders() {
			var header HeaderMatcher
			switch ht := h.GetHeaderMatchSpecifier().(type) {
			case *v3routepb.HeaderMatcher_ExactMatch:
				header.ExactMatch = &ht.ExactMatch
			case *v3routepb.HeaderMatcher_SafeRegexMatch:
				header.RegexMatch = &ht.SafeRegexMatch.Regex
			case *v3routepb.HeaderMatcher_RangeMatch:
				header.RangeMatch = &Int64Range{
					Start: ht.RangeMatch.Start,
					End:   ht.RangeMatch.End,
				}
			case *v3routepb.HeaderMatcher_PresentMatch:
				header.PresentMatch = &ht.PresentMatch
			case *v3routepb.HeaderMatcher_PrefixMatch:
				header.PrefixMatch = &ht.PrefixMatch
			case *v3routepb.HeaderMatcher_SuffixMatch:
				header.SuffixMatch = &ht.SuffixMatch
			default:
				return nil, fmt.Errorf("route %+v has an unrecognized header matcher: %+v", r, ht)
			}
			header.Name = h.GetName()
			invert := h.GetInvertMatch()
			header.InvertMatch = &invert
			route.Headers = append(route.Headers, &header)
		}

		if fr := match.GetRuntimeFraction(); fr != nil {
			d := fr.GetDefaultValue()
			n := d.GetNumerator()
			switch d.GetDenominator() {
			case v3typepb.FractionalPercent_HUNDRED:
				n *= 10000
			case v3typepb.FractionalPercent_TEN_THOUSAND:
				n *= 100
			case v3typepb.FractionalPercent_MILLION:
			}
			route.Fraction = &n
		}

		route.WeightedClusters = make(map[string]WeightedCluster)
		action := r.GetRoute()
		switch a := action.GetClusterSpecifier().(type) {
		case *v3routepb.RouteAction_Cluster:
			route.WeightedClusters[a.Cluster] = WeightedCluster{Weight: 1}
		case *v3routepb.RouteAction_WeightedClusters:
			wcs := a.WeightedClusters
			var totalWeight uint32
			for _, c := range wcs.Clusters {
				w := c.GetWeight().GetValue()
				if w == 0 {
					continue
				}
				wc := WeightedCluster{Weight: w}
				if !v2 {
					cfgs, err := processHTTPFilterOverrides(c.GetTypedPerFilterConfig())
					if err != nil {
						return nil, fmt.Errorf("route %+v, action %+v: %v", r, a, err)
					}
					wc.HTTPFilterConfigOverride = cfgs
				}
				route.WeightedClusters[c.GetName()] = wc
				totalWeight += w
			}
			if totalWeight != wcs.GetTotalWeight().GetValue() {
				return nil, fmt.Errorf("route %+v, action %+v, weights of clusters do not add up to total total weight, got: %v, want %v", r, a, wcs.GetTotalWeight().GetValue(), totalWeight)
			}
			if totalWeight == 0 {
				return nil, fmt.Errorf("route %+v, action %+v, has no valid cluster in WeightedCluster action", r, a)
			}
		case *v3routepb.RouteAction_ClusterHeader:
			continue
		}

		msd := action.GetMaxStreamDuration()
		// Prefer grpc_timeout_header_max, if set.
		dur := msd.GetGrpcTimeoutHeaderMax()
		if dur == nil {
			dur = msd.GetMaxStreamDuration()
		}
		if dur != nil {
			d := dur.AsDuration()
			route.MaxStreamDuration = &d
		}

		if !v2 {
			cfgs, err := processHTTPFilterOverrides(r.GetTypedPerFilterConfig())
			if err != nil {
				return nil, fmt.Errorf("route %+v: %v", r, err)
			}
			route.HTTPFilterConfigOverride = cfgs
		}
		routesRet = append(routesRet, &route)
	}
	return routesRet, nil
}

// UnmarshalCluster processes resources received in an CDS response, validates
// them, and transforms them into a native struct which contains only fields we
// are interested in.
func UnmarshalCluster(version string, resources []*anypb.Any, logger *grpclog.PrefixLogger) (map[string]ClusterUpdate, UpdateMetadata, error) {
	update := make(map[string]ClusterUpdate)
	for _, r := range resources {
		if !IsClusterResource(r.GetTypeUrl()) {
			return nil, UpdateMetadata{}, fmt.Errorf("xds: unexpected resource type: %q in CDS response", r.GetTypeUrl())
		}

		cluster := &v3clusterpb.Cluster{}
		if err := proto.Unmarshal(r.GetValue(), cluster); err != nil {
			return nil, UpdateMetadata{}, fmt.Errorf("xds: failed to unmarshal resource in CDS response: %v", err)
		}
		logger.Infof("Resource with name: %v, type: %T, contains: %v", cluster.GetName(), cluster, cluster)
		cu, err := validateCluster(cluster)
		if err != nil {
			return nil, UpdateMetadata{}, err
		}

		// If the Cluster message in the CDS response did not contain a
		// serviceName, we will just use the clusterName for EDS.
		if cu.ServiceName == "" {
			cu.ServiceName = cluster.GetName()
		}
		logger.Debugf("Resource with name %v, value %+v added to cache", cluster.GetName(), cu)
		update[cluster.GetName()] = cu
	}
	return update, UpdateMetadata{}, nil
}

func validateCluster(cluster *v3clusterpb.Cluster) (ClusterUpdate, error) {
	emptyUpdate := ClusterUpdate{ServiceName: "", EnableLRS: false}
	switch {
	case cluster.GetType() != v3clusterpb.Cluster_EDS:
		return emptyUpdate, fmt.Errorf("xds: unexpected cluster type %v in response: %+v", cluster.GetType(), cluster)
	case cluster.GetEdsClusterConfig().GetEdsConfig().GetAds() == nil:
		return emptyUpdate, fmt.Errorf("xds: unexpected edsConfig in response: %+v", cluster)
	case cluster.GetLbPolicy() != v3clusterpb.Cluster_ROUND_ROBIN:
		return emptyUpdate, fmt.Errorf("xds: unexpected lbPolicy %v in response: %+v", cluster.GetLbPolicy(), cluster)
	}

	sc, err := securityConfigFromCluster(cluster)
	if err != nil {
		return emptyUpdate, err
	}
	return ClusterUpdate{
		ServiceName: cluster.GetEdsClusterConfig().GetServiceName(),
		EnableLRS:   cluster.GetLrsServer().GetSelf() != nil,
		SecurityCfg: sc,
		MaxRequests: circuitBreakersFromCluster(cluster),
	}, nil
}

// securityConfigFromCluster extracts the relevant security configuration from
// the received Cluster resource.
func securityConfigFromCluster(cluster *v3clusterpb.Cluster) (*SecurityConfig, error) {
	// The Cluster resource contains a `transport_socket` field, which contains
	// a oneof `typed_config` field of type `protobuf.Any`. The any proto
	// contains a marshaled representation of an `UpstreamTlsContext` message.
	ts := cluster.GetTransportSocket()
	if ts == nil {
		return nil, nil
	}
	if name := ts.GetName(); name != transportSocketName {
		return nil, fmt.Errorf("xds: transport_socket field has unexpected name: %s", name)
	}
	any := ts.GetTypedConfig()
	if any == nil || any.TypeUrl != version.V3UpstreamTLSContextURL {
		return nil, fmt.Errorf("xds: transport_socket field has unexpected typeURL: %s", any.TypeUrl)
	}
	upstreamCtx := &v3tlspb.UpstreamTlsContext{}
	if err := proto.Unmarshal(any.GetValue(), upstreamCtx); err != nil {
		return nil, fmt.Errorf("xds: failed to unmarshal UpstreamTlsContext in CDS response: %v", err)
	}
	if upstreamCtx.GetCommonTlsContext() == nil {
		return nil, errors.New("xds: UpstreamTlsContext in CDS response does not contain a CommonTlsContext")
	}

	sc, err := securityConfigFromCommonTLSContext(upstreamCtx.GetCommonTlsContext())
	if err != nil {
		return nil, err
	}
	if sc.RootInstanceName == "" {
		return nil, errors.New("security configuration on the client-side does not contain root certificate provider instance name")
	}
	return sc, nil
}

// common is expected to be not nil.
func securityConfigFromCommonTLSContext(common *v3tlspb.CommonTlsContext) (*SecurityConfig, error) {
	// The `CommonTlsContext` contains a
	// `tls_certificate_certificate_provider_instance` field of type
	// `CertificateProviderInstance`, which contains the provider instance name
	// and the certificate name to fetch identity certs.
	sc := &SecurityConfig{}
	if identity := common.GetTlsCertificateCertificateProviderInstance(); identity != nil {
		sc.IdentityInstanceName = identity.GetInstanceName()
		sc.IdentityCertName = identity.GetCertificateName()
	}

	// The `CommonTlsContext` contains a `validation_context_type` field which
	// is a oneof. We can get the values that we are interested in from two of
	// those possible values:
	//  - combined validation context:
	//    - contains a default validation context which holds the list of
	//      accepted SANs.
	//    - contains certificate provider instance configuration
	//  - certificate provider instance configuration
	//    - in this case, we do not get a list of accepted SANs.
	switch t := common.GetValidationContextType().(type) {
	case *v3tlspb.CommonTlsContext_CombinedValidationContext:
		combined := common.GetCombinedValidationContext()
		if def := combined.GetDefaultValidationContext(); def != nil {
			for _, matcher := range def.GetMatchSubjectAltNames() {
				// We only support exact matches for now.
				if exact := matcher.GetExact(); exact != "" {
					sc.AcceptedSANs = append(sc.AcceptedSANs, exact)
				}
			}
		}
		if pi := combined.GetValidationContextCertificateProviderInstance(); pi != nil {
			sc.RootInstanceName = pi.GetInstanceName()
			sc.RootCertName = pi.GetCertificateName()
		}
	case *v3tlspb.CommonTlsContext_ValidationContextCertificateProviderInstance:
		pi := common.GetValidationContextCertificateProviderInstance()
		sc.RootInstanceName = pi.GetInstanceName()
		sc.RootCertName = pi.GetCertificateName()
	case nil:
		// It is valid for the validation context to be nil on the server side.
	default:
		return nil, fmt.Errorf("xds: validation context contains unexpected type: %T", t)
	}
	return sc, nil
}

// circuitBreakersFromCluster extracts the circuit breakers configuration from
// the received cluster resource. Returns nil if no CircuitBreakers or no
// Thresholds in CircuitBreakers.
func circuitBreakersFromCluster(cluster *v3clusterpb.Cluster) *uint32 {
	if !env.CircuitBreakingSupport {
		return nil
	}
	for _, threshold := range cluster.GetCircuitBreakers().GetThresholds() {
		if threshold.GetPriority() != v3corepb.RoutingPriority_DEFAULT {
			continue
		}
		maxRequestsPb := threshold.GetMaxRequests()
		if maxRequestsPb == nil {
			return nil
		}
		maxRequests := maxRequestsPb.GetValue()
		return &maxRequests
	}
	return nil
}

// UnmarshalEndpoints processes resources received in an EDS response,
// validates them, and transforms them into a native struct which contains only
// fields we are interested in.
func UnmarshalEndpoints(version string, resources []*anypb.Any, logger *grpclog.PrefixLogger) (map[string]EndpointsUpdate, UpdateMetadata, error) {
	update := make(map[string]EndpointsUpdate)
	for _, r := range resources {
		if !IsEndpointsResource(r.GetTypeUrl()) {
			return nil, UpdateMetadata{}, fmt.Errorf("xds: unexpected resource type: %q in EDS response", r.GetTypeUrl())
		}

		cla := &v3endpointpb.ClusterLoadAssignment{}
		if err := proto.Unmarshal(r.GetValue(), cla); err != nil {
			return nil, UpdateMetadata{}, fmt.Errorf("xds: failed to unmarshal resource in EDS response: %v", err)
		}
		logger.Infof("Resource with name: %v, type: %T, contains: %v", cla.GetClusterName(), cla, cla)

		u, err := parseEDSRespProto(cla)
		if err != nil {
			return nil, UpdateMetadata{}, err
		}
		update[cla.GetClusterName()] = u
	}
	return update, UpdateMetadata{}, nil
}

func parseAddress(socketAddress *v3corepb.SocketAddress) string {
	return net.JoinHostPort(socketAddress.GetAddress(), strconv.Itoa(int(socketAddress.GetPortValue())))
}

func parseDropPolicy(dropPolicy *v3endpointpb.ClusterLoadAssignment_Policy_DropOverload) OverloadDropConfig {
	percentage := dropPolicy.GetDropPercentage()
	var (
		numerator   = percentage.GetNumerator()
		denominator uint32
	)
	switch percentage.GetDenominator() {
	case v3typepb.FractionalPercent_HUNDRED:
		denominator = 100
	case v3typepb.FractionalPercent_TEN_THOUSAND:
		denominator = 10000
	case v3typepb.FractionalPercent_MILLION:
		denominator = 1000000
	}
	return OverloadDropConfig{
		Category:    dropPolicy.GetCategory(),
		Numerator:   numerator,
		Denominator: denominator,
	}
}

func parseEndpoints(lbEndpoints []*v3endpointpb.LbEndpoint) []Endpoint {
	endpoints := make([]Endpoint, 0, len(lbEndpoints))
	for _, lbEndpoint := range lbEndpoints {
		endpoints = append(endpoints, Endpoint{
			HealthStatus: EndpointHealthStatus(lbEndpoint.GetHealthStatus()),
			Address:      parseAddress(lbEndpoint.GetEndpoint().GetAddress().GetSocketAddress()),
			Weight:       lbEndpoint.GetLoadBalancingWeight().GetValue(),
		})
	}
	return endpoints
}

func parseEDSRespProto(m *v3endpointpb.ClusterLoadAssignment) (EndpointsUpdate, error) {
	ret := EndpointsUpdate{}
	for _, dropPolicy := range m.GetPolicy().GetDropOverloads() {
		ret.Drops = append(ret.Drops, parseDropPolicy(dropPolicy))
	}
	priorities := make(map[uint32]struct{})
	for _, locality := range m.Endpoints {
		l := locality.GetLocality()
		if l == nil {
			return EndpointsUpdate{}, fmt.Errorf("EDS response contains a locality without ID, locality: %+v", locality)
		}
		lid := xdsinternal.LocalityID{
			Region:  l.Region,
			Zone:    l.Zone,
			SubZone: l.SubZone,
		}
		priority := locality.GetPriority()
		priorities[priority] = struct{}{}
		ret.Localities = append(ret.Localities, Locality{
			ID:        lid,
			Endpoints: parseEndpoints(locality.GetLbEndpoints()),
			Weight:    locality.GetLoadBalancingWeight().GetValue(),
			Priority:  priority,
		})
	}
	for i := 0; i < len(priorities); i++ {
		if _, ok := priorities[uint32(i)]; !ok {
			return EndpointsUpdate{}, fmt.Errorf("priority %v missing (with different priorities %v received)", i, priorities)
		}
	}
	return ret, nil
}
