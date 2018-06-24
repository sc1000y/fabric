package main

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	pb "github.com/hyperledger/fabric/orderer/consensus/dual/grpc"
	"google.golang.org/grpc"
)

//client.go

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func TestStart(t *testing.T) {
	var mockLag = 100
	go start(port)

	sleepTime := rand.Intn(mockLag)
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
	mockClient()
}

func mockClient() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("did not connect: %v", err)
		//log.Fatalln("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBackendServiceClient(conn)
	r, err := c.GetPeerInfo(context.Background(), &pb.PeerRequest{Greeting: "1"})
	if err != nil {
		logger.Fatal("could not greet: %v", err)
	}
	fmt.Printf("Greeting: %f", r.GetCredit())

}
