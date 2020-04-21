package main

import (
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	listener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	v2route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/golang/protobuf/ptypes"

	hcm "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"log"
)

func BuildListener() []cache.Resource {
	var clusterName = "service_bbc"
	var listenerName = "listener_0"
	var targetPrefix = "/hello"
	var virtualHostName = "service"
	var routeConfigName = "local_route"

	log.Println(">>>>>>>>>>>>>>>>>>> creating listener ", listenerName)

	virtualHost := v2route.VirtualHost{
		Name:    virtualHostName,
		Domains: []string{"*"},

		Routes: []*v2route.Route{{
			Match: &v2route.RouteMatch{
				PathSpecifier: &v2route.RouteMatch_Prefix{
					Prefix: targetPrefix,
				},
			},

			Action: &v2route.Route_Route{
				Route: &v2route.RouteAction{
					ClusterSpecifier: &v2route.RouteAction_Cluster{
						Cluster: clusterName,
					},
				},
			},
		}}}

	manager := &hcm.HttpConnectionManager{
		CodecType:  hcm.HttpConnectionManager_AUTO,
		StatPrefix: "ingress_http",
		RouteSpecifier: &hcm.HttpConnectionManager_RouteConfig{
			RouteConfig: &v2.RouteConfiguration{
				Name:         routeConfigName,
				VirtualHosts: []*v2route.VirtualHost{&virtualHost},
			},
		},
		HttpFilters: []*hcm.HttpFilter{{
			Name: "envoy.filters.http.router",
		}},
	}

	pbst, err := ptypes.MarshalAny(manager)
	if err != nil {
		panic(err)
	}

	return []cache.Resource{
		&v2.Listener{
			Name: listenerName,
			Address: &core.Address{
				Address: &core.Address_SocketAddress{
					SocketAddress: &core.SocketAddress{
						Protocol: core.SocketAddress_TCP,
						Address:  "0.0.0.0",
						PortSpecifier: &core.SocketAddress_PortValue{
							PortValue: 80,
						},
					},
				},
			},
			FilterChains: []*listener.FilterChain{{
				Filters: []*listener.Filter{{
					Name: "envoy.filters.network.http_connection_manager",
					ConfigType: &listener.Filter_TypedConfig{
						TypedConfig: pbst,
					},
				}},
			}},
		}}
}
