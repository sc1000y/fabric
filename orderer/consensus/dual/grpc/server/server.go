package main

//package main

// server.go

import (
	"log"
	"net"

	pb "github.com/hyperledger/fabric/orderer/consensus/dual/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port    = ":50051"
	tstPort = ":50052"
)

type server struct{}

func (s *server) GetPeerInfo(ctx context.Context, in *pb.PeerRequest) (*pb.PeerInfoResponse, error) {
	return &pb.PeerInfoResponse{SeralizedId: 1, Credit: 20.0, AmIprimary: false}, nil
}
func (s *server) IwantoBePrimary(ctx context.Context, in *pb.IwantToBePrimaryRequest) (*pb.IwantToBePrimaryResponse, error) {
	return &pb.IwantToBePrimaryResponse{Success: false}, nil
}

//func (s *server)
//func
func main() {
	/*lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//pb.RegisterHelloServiceServer(s, &server{})
	pb.RegisterBackendServiceServer(s, &server{})
	s.Serve(lis)*/
	start(port)
	start(tstPort)

}
func start(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//pb.RegisterHelloServiceServer(s, &server{})
	pb.RegisterBackendServiceServer(s, &server{})
	s.Serve(lis)
}
