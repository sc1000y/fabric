package main

//package main

// server.go

import (
	"net"

	"github.com/hyperledger/fabric/common/flogging"
	pb "github.com/hyperledger/fabric/orderer/consensus/dual/grpc"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port    = ":50051"
	tstPort = ":50052"
)
const pkgLogID = "orderer/common/orderer/consensus/dual/grpc/server"

var logger *logging.Logger

func init() {
	logger = flogging.MustGetLogger(pkgLogID)
}

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
		logger.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//pb.RegisterHelloServiceServer(s, &server{})
	pb.RegisterBackendServiceServer(s, &server{})
	s.Serve(lis)
}
