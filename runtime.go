package main

import (
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	pstruct "github.com/golang/protobuf/ptypes/struct"
)

func BuildRuntime() []types.Resource {
	return []types.Resource{
		&discovery.Runtime{
			Name: "my-layer",
			Layer: &pstruct.Struct{
				Fields: map[string]*pstruct.Value{
					"tracing.client_enabled": {
						Kind: &pstruct.Value_NumberValue{NumberValue: 20},
					},
					"field-1": {
						Kind: &pstruct.Value_StringValue{StringValue: "foobar"},
					},
				},
			},
		},
	}
}
