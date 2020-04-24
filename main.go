package main

import (
	"context"
	"fmt"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	xds "github.com/envoyproxy/go-control-plane/pkg/server/v2"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ctx := context.Background()
	log.Printf("Starting control plane")

	snapshotCache := cache.NewSnapshotCache(false, cache.IDHash{}, nil)
	snapshot := cache.NewSnapshot("1.0", BuildEndpoint(), BuildCluster(), BuildRouter(), BuildListener(), BuildRuntime())
	_ = snapshotCache.SetSnapshot("node1", snapshot)

	myCallbacks := MyCallbacks{}
	srv := xds.NewServer(ctx, snapshotCache, &myCallbacks)

	//RunManagementGateway(ctx, srv, 9001)
	RunManagementServer(ctx, srv, 9002)

	<-ctx.Done()
}

func RunManagementServer(ctx context.Context, server xds.Server, port uint) {
	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	v2.RegisterEndpointDiscoveryServiceServer(grpcServer, server)
	v2.RegisterClusterDiscoveryServiceServer(grpcServer, server)
	v2.RegisterRouteDiscoveryServiceServer(grpcServer, server)
	v2.RegisterListenerDiscoveryServiceServer(grpcServer, server)

	log.Println("management server listening: ", port)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}

//func RunManagementGateway(ctx context.Context, srv xds.Server, port uint) {
//	log.Println("gateway listening HTTP/1.1 :", port)
//	server := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: &xds.HTTPGateway{Server: srv}}
//	go func() {
//		if err := server.ListenAndServe(); err != nil {
//			log.Fatal(err)
//		}
//	}()
//}
