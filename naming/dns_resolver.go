package naming

import (
	"fmt"
	"google.golang.org/grpc/grpclog"
	"net"
	"sort"
	"strconv"
	"strings"
)

type DNSResolver struct {
}

func (r *DNSResolver) Resolve(target string) (DNSWatcher, error) {
	return DNSWatcher{
		hostname: target,
	}, nil
}

type DNSWatcher struct {
	// hostname to watch address Update
	hostname string
	// The latest resolved address list
	curAddrs []*Update
}

type AddressType uint8

const (
	// Backend indicates the server is a backend server.
	Backend AddressType = iota
	// GRPCLB indicates the server is a grpclb load balancer.
	GRPCLB
)

type AddrMetadataGRPCLB struct {
	// AddrType is the type of server (grpc load balancer or backend).
	AddrType AddressType
	// ServerName is the name of the grpc load balancer. Used for authentication.
	ServerName string
}

func compileUpdate(oldAddrs []*Update, newAddrs []*Update) []*Update {
	result := make([]*Update, 0, len(oldAddrs)+len(newAddrs))
	idx1, idx2 := 0, 0
	for idx1 < len(oldAddrs) || idx2 < len(newAddrs) {
		if idx1 == len(oldAddrs) {
			// add all adrress left in addrs
			for _, addr := range newAddrs[idx2:] {
				u := *addr
				u.Op = Add
				result = append(result, &u)
			}
			return result
		}
		if idx2 == len(newAddrs) {
			// remove all address left in cur addrs
			for _, addr := range oldAddrs[idx1:] {
				u := *addr
				u.Op = Delete
				result = append(result, &u)
			}
			return result
		}
		switch strings.Compare(oldAddrs[idx1].Addr, newAddrs[idx2].Addr) {
		case 0:
			if oldAddrs[idx1].Metadata != newAddrs[idx2].Metadata {
				uDel := *oldAddrs[idx1]
				uDel.Op = Delete
				result = append(result, &uDel)
				uAdd := *newAddrs[idx2]
				uAdd.Op = Add
				result = append(result, &uAdd)
			}
			idx1++
			idx2++
		case -1:
			u := *oldAddrs[idx1]
			u.Op = Delete
			result = append(result, &u)
			idx1++
		case 1:
			u := *newAddrs[idx2]
			u.Op = Add
			result = append(result, &u)
			idx2++
		}
	}
	return result
}

func (w *DNSWatcher) Next() ([]*Update, error) {
	// TODO(yuxuanli): handling port?
	cname, srvs, err := net.LookupSRV("grpclb", "tcp", w.hostname)
	if err != nil {
		grpclog.Printf("grpc: failed dns srv lookup due to %v.\n", err)
	}
	fmt.Println(cname)
	for _, rc := range srvs {
		fmt.Printf("%s %d %d %d\n", rc.Target, rc.Port, rc.Priority, rc.Weight)
	}
	// target has SRV records associated with it
	if len(srvs) > 0 {
		newAddrs := make([]*Update, 0, 1000 /* TODO: decide the number here*/)
		for _, r := range srvs {
			lbAddrs, err := net.LookupHost(r.Target)
			if err != nil {
				grpclog.Printf("grpc: failed dns srv load banlacer address lookup due to %v.\n", err)
			}
			for _, a := range lbAddrs {
				newAddrs = append(newAddrs, &Update{Addr: a + ":" + strconv.Itoa(int(r.Port)),
					Metadata: AddrMetadataGRPCLB{AddrType: GRPCLB, ServerName: r.Target},
				})
			}
		}
		sort.SliceStable(newAddrs, func(i, j int) bool { return strings.Compare(newAddrs[i].Addr, newAddrs[j].Addr) < 0 })
		result := compileUpdate(w.curAddrs, newAddrs)
		w.curAddrs = newAddrs
		return result, nil
	}

	addrs, err := net.LookupHost(w.hostname)
	if err != nil {
		grpclog.Printf("grpc: failed dns resolution due to %v.\n", err)
	}
	sort.Strings(addrs)
	newAddrs := make([]*Update, len(addrs))
	for i, a := range addrs {
		newAddrs[i] = &Update{Addr: a}
	}
	result := compileUpdate(w.curAddrs, newAddrs)
	w.curAddrs = newAddrs
	return result, nil
}

func (w *DNSWatcher) Close() {
}
