package main

import (
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/golang/protobuf/ptypes"
	"log"
	"time"
)

func BuildCluster() []types.Resource {
	var clusterName1 = "service_bbc"
	log.Println(">>>>>>>>>>>>>>>>>>> creating cluster ", clusterName1)

	return []types.Resource{
		&v2.Cluster{
			Name:                 clusterName1,
			ConnectTimeout:       ptypes.DurationProto(2 * time.Second),
			ClusterDiscoveryType: &v2.Cluster_Type{Type: v2.Cluster_EDS},
			LbPolicy:             v2.Cluster_ROUND_ROBIN,
			EdsClusterConfig: &v2.Cluster_EdsClusterConfig{
				ServiceName: clusterName1,
				EdsConfig: &core.ConfigSource{
					ResourceApiVersion: core.ApiVersion_V2,
					ConfigSourceSpecifier: &core.ConfigSource_ApiConfigSource{
						ApiConfigSource: &core.ApiConfigSource{
							ApiType:             core.ApiConfigSource_GRPC,
							TransportApiVersion: core.ApiVersion_V2,
							GrpcServices: []*core.GrpcService{
								{
									TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
										EnvoyGrpc: &core.GrpcService_EnvoyGrpc{
											ClusterName: "xds_cluster",
										},
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
