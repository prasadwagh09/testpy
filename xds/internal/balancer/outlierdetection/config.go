/*
 *
 * Copyright 2022 gRPC authors.
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
 */

package outlierdetection

import (
	"time"

	"github.com/google/go-cmp/cmp"
	internalserviceconfig "google.golang.org/grpc/internal/serviceconfig"
	"google.golang.org/grpc/serviceconfig"
)

// SuccessRateEjection is parameters for the success rate ejection algorithm.
// This algorithm monitors the request success rate for all endpoints and ejects
// individual endpoints whose success rates are statistical outliers.
type SuccessRateEjection struct {
	// StddevFactor is used to determine the ejection threshold for
	// success rate outlier ejection. The ejection threshold is the difference
	// between the mean success rate, and the product of this factor and the
	// standard deviation of the mean success rate: mean - (stdev *
	// success_rate_stdev_factor). This factor is divided by a thousand to get a
	// double. That is, if the desired factor is 1.9, the runtime value should
	// be 1900. Defaults to 1900.
	StdevFactor uint32 `json:"stdevFactor,omitempty"`
	// EnforcingSuccessRate is the % chance that a host will be actually ejected
	// when an outlier status is detected through success rate statistics. This
	// setting can be used to disable ejection or to ramp it up slowly. Defaults
	// to 100.
	EnforcingSuccessRate uint32 `json:"enforcingSuccessRate,omitempty"`
	// MinimumHosts is the number of hosts in a cluster that must have enough
	// request volume to detect success rate outliers. If the number of hosts is
	// less than this setting, outlier detection via success rate statistics is
	// not performed for any host in the cluster. Defaults to 5.
	MinimumHosts uint32 `json:"minimumHosts,omitempty"`
	// RequestVolume is the minimum number of total requests that must be
	// collected in one interval (as defined by the interval duration above) to
	// include this host in success rate based outlier detection. If the volume
	// is lower than this setting, outlier detection via success rate statistics
	// is not performed for that host. Defaults to 100.
	RequestVolume uint32 `json:"requestVolume,omitempty"`
}

// Equal returns whether the SuccessRateEjection is the same with the parameter.
func (sre *SuccessRateEjection) Equal(sre2 *SuccessRateEjection) bool {
	if sre == nil && sre2 == nil {
		return true
	}
	if (sre != nil) != (sre2 != nil) {
		return false
	}
	if sre.StdevFactor != sre2.StdevFactor {
		return false
	}
	if sre.EnforcingSuccessRate != sre2.EnforcingSuccessRate {
		return false
	}
	if sre.MinimumHosts != sre2.MinimumHosts {
		return false
	}
	return sre.RequestVolume == sre2.RequestVolume
}

// FailurePercentageEjection is parameters for the failure percentage algorithm.
// This algorithm ejects individual endpoints whose failure rate is greater than
// some threshold, independently of any other endpoint.
type FailurePercentageEjection struct {
	// Threshold is the failure percentage to use when determining failure
	// percentage-based outlier detection. If the failure percentage of a given
	// host is greater than or equal to this value, it will be ejected. Defaults
	// to 85.
	Threshold uint32 `json:"threshold,omitempty"`
	// EnforcingFailurePercentage is the % chance that a host will be actually
	// ejected when an outlier status is detected through failure percentage
	// statistics. This setting can be used to disable ejection or to ramp it up
	// slowly. Defaults to 0.
	EnforcingFailurePercentage uint32 `json:"enforcingFailurePercentage,omitempty"`
	// MinimumHosts is the minimum number of hosts in a cluster in order to
	// perform failure percentage-based ejection. If the total number of hosts
	// in the cluster is less than this value, failure percentage-based ejection
	// will not be performed. Defaults to 5.
	MinimumHosts uint32 `json:"minimumHosts,omitempty"`
	// RequestVolume is the minimum number of total requests that must be
	// collected in one interval (as defined by the interval duration above) to
	// perform failure percentage-based ejection for this host. If the volume is
	// lower than this setting, failure percentage-based ejection will not be
	// performed for this host. Defaults to 50.
	RequestVolume uint32 `json:"requestVolume,omitempty"`
}

// Equal returns whether the FailurePercentageEjection is the same with the
// parameter.
func (fpe *FailurePercentageEjection) Equal(fpe2 *FailurePercentageEjection) bool {
	if fpe == nil && fpe2 == nil {
		return true
	}
	if (fpe != nil) != (fpe2 != nil) {
		return false
	}
	if fpe.Threshold != fpe2.Threshold {
		return false
	}
	if fpe.EnforcingFailurePercentage != fpe2.EnforcingFailurePercentage {
		return false
	}
	if fpe.MinimumHosts != fpe2.MinimumHosts {
		return false
	}
	return fpe.RequestVolume == fpe2.RequestVolume
}

// LBConfig is the config for the outlier detection balancer.
type LBConfig struct {
	serviceconfig.LoadBalancingConfig `json:"-"`
	// Interval is the time interval between ejection analysis sweeps. This can
	// result in both new ejections as well as addresses being returned to
	// service. Defaults to 10s.
	Interval time.Duration `json:"interval,omitempty"`
	// BaseEjectionTime is the base time that a host is ejected for. The real
	// time is equal to the base time multiplied by the number of times the host
	// has been ejected and is capped by MaxEjectionTime. Defaults to 30s.
	BaseEjectionTime time.Duration `json:"baseEjectionTime,omitempty"`
	// MaxEjectionTime is the maximum time that an address is ejected for. If
	// not specified, the default value (300s) or the BaseEjectionTime value is
	// applied, whichever is larger.
	MaxEjectionTime time.Duration `json:"maxEjectionTime,omitempty"`
	// MaxEjectionPercent is the maximum % of an upstream cluster that can be
	// ejected due to outlier detection. Defaults to 10% but will eject at least
	// one host regardless of the value.
	MaxEjectionPercent uint32 `json:"maxEjectionPercent,omitempty"`
	// SuccessRateEjection is the parameters for the success rate ejection
	// algorithm. If set, success rate ejections will be performed.
	SuccessRateEjection *SuccessRateEjection `json:"successRateEjection,omitempty"`
	// FailurePercentageEjection is the parameters for the failure percentage
	// algorithm. If set, failure rate ejections will be performed.
	FailurePercentageEjection *FailurePercentageEjection `json:"failurePercentageEjection,omitempty"`
	// ChildPolicy is the config for the child policy.
	ChildPolicy/*[] why is this repeated in design?*/ *internalserviceconfig.BalancerConfig `json:"childPolicy,omitempty"`
}

// Equal returns whether the LBConfig is the same with the parameter.
func (lbc *LBConfig) Equal(lbc2 *LBConfig) bool {
	if lbc == nil && lbc2 == nil {
		return true
	}
	if (lbc != nil) != (lbc2 != nil) {
		return false
	}
	if lbc.Interval != lbc2.Interval {
		return false
	}
	if lbc.BaseEjectionTime != lbc2.BaseEjectionTime {
		return false
	}
	if lbc.MaxEjectionTime != lbc2.MaxEjectionTime {
		return false
	}
	if lbc.MaxEjectionPercent != lbc2.MaxEjectionPercent {
		return false
	}
	if !lbc.SuccessRateEjection.Equal(lbc2.SuccessRateEjection) {
		return false
	}
	if !lbc.FailurePercentageEjection.Equal(lbc2.FailurePercentageEjection) {
		return false
	}
	return cmp.Equal(lbc.ChildPolicy, lbc2.ChildPolicy)
}
