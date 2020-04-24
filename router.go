package main

import (
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
)

func BuildRouter() []types.Resource {

	return []types.Resource{
		&v2.RouteConfiguration{
			Name: "my-route",
			VirtualHosts: []*route.VirtualHost{
				{
					Name:    "my-virtual-host",
					Domains: []string{"*"},
					Routes: []*route.Route{
						{
							Match: &route.RouteMatch{
								PathSpecifier: &route.RouteMatch_Prefix{
									Prefix: "/",
								},
							},
							Action: &route.Route_Route{
								Route: &route.RouteAction{
									ClusterSpecifier: &route.RouteAction_Cluster{
										Cluster: "service_bbc",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
