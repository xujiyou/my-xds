package main

import (
	"context"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"log"
)

type MyCallbacks struct{}

func (cb *MyCallbacks) Report() {
	log.Println("Report...")
}

func (cb *MyCallbacks) OnStreamOpen(ctx context.Context, id int64, typ string) error {
	log.Println("OnStreamOpen...")
	return nil
}

func (cb *MyCallbacks) OnStreamClosed(id int64) {
	log.Println("OnStreamClosed...")
}

func (cb *MyCallbacks) OnStreamRequest(int64, *v2.DiscoveryRequest) error {
	log.Println("OnStreamRequest...")
	return nil
}

func (cb *MyCallbacks) OnStreamResponse(int64, *v2.DiscoveryRequest, *v2.DiscoveryResponse) {
	log.Println("OnStreamResponse...")
	cb.Report()
}

func (cb *MyCallbacks) OnFetchRequest(ctx context.Context, req *v2.DiscoveryRequest) error {
	log.Println("OnFetchRequest, req:", req)
	return nil
}

func (cb *MyCallbacks) OnFetchResponse(*v2.DiscoveryRequest, *v2.DiscoveryResponse) {
	log.Println("OnFetchResponse...")
}
