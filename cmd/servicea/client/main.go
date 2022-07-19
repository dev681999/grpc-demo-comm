package main

import (
	"context"
	"flag"
	"grpc-demo-comm/pkg/proto/servicea"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var flagServiceAAddr = flag.String("addr", "localhost:5000", "addr of servicea")

func main() {
	flag.Parse()

	// first create servicea client
	conn, err := grpc.Dial(*flagServiceAAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := servicea.NewServiceAAPIClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Hello(ctx, &servicea.HelloRequest{
		Num: 43,
	})
	if err != nil {
		log.Fatalf("could not make hello call: %v", err)
	}
	log.Printf("a: %d, b: %d", r.GetNumA(), r.GetNumB())
}
