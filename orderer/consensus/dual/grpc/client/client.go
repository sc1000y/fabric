package main

import (
	"context"
	"log"

	pb "github.com/hyperledger/fabric/orderer/consensus/dual/grpc"
	"google.golang.org/grpc"
)

//client.go

const (
	address     = "localhost:50051"
	defaultName = "world"
)

/*func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHelloServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Greeting: name})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Reply)
}*/
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBackendServiceClient(conn)

	//name := defaultName
	/*if len(os.Args) > 1 {
		name = os.Args[1]
	}*/
	r, err := c.GetPeerInfo(context.Background(), &pb.PeerRequest{Greeting: "1"})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("response is: %f", r.Credit)
}
