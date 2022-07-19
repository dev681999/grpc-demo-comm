package main

import (
	"flag"
	"grpc-demo-comm/pkg/proto/servicea"
	"grpc-demo-comm/pkg/proto/serviceb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var flagServiceBAddr = flag.String("serviceb-addr", "localhost:6000", "addr of serviceb")
var flagAddr = flag.String("addr", "localhost:5000", "addr to run servicea on")

func main() {
	flag.Parse()

	// first create serviceb client
	conn, err := grpc.Dial(*flagServiceBAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := serviceb.NewServiceBAPIClient(conn)

	// start a listener for servicea grpc server to listen to
	lis, err := net.Listen("tcp", *flagAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	servicea.RegisterServiceAAPIServer(s, &server{
		servicebAPI: c,
	})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
