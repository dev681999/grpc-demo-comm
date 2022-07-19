package main

import (
	"context"
	"grpc-demo-comm/pkg/proto/serviceb"
)

type server struct {
	serviceb.UnimplementedServiceBAPIServer
}

var _ serviceb.ServiceBAPIServer = &server{}

func (s *server) Hello(ctx context.Context, req *serviceb.HelloRequest) (*serviceb.HelloResponse, error) {
	return &serviceb.HelloResponse{
		Resp: req.GetNum() * 2,
	}, nil
}
