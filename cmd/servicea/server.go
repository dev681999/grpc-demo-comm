package main

import (
	"context"
	"grpc-demo-comm/pkg/proto/servicea"
	"grpc-demo-comm/pkg/proto/serviceb"
)

type server struct {
	servicea.UnimplementedServiceAAPIServer
	servicebAPI serviceb.ServiceBAPIClient
}

var _ servicea.ServiceAAPIServer = &server{}

func (s *server) Hello(ctx context.Context, req *servicea.HelloRequest) (*servicea.HelloResponse, error) {
	respB, err := s.servicebAPI.Hello(ctx, &serviceb.HelloRequest{
		Num: req.GetNum(),
	})
	if err != nil {
		return nil, err
	}
	return &servicea.HelloResponse{
		NumA: req.GetNum() * 10,
		NumB: respB.GetResp(),
	}, nil
}
