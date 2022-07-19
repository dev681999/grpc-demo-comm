package main

import (
	"flag"
	"grpc-demo-comm/pkg/proto/serviceb"
	"log"
	"net"

	"google.golang.org/grpc"
)

var flagAddr = flag.String("addr", "localhost:6000", "addr to run serviceb on")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *flagAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	serviceb.RegisterServiceBAPIServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
