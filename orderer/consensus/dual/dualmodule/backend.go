package main

import (
	"net"

	"github.com/hyperledger/fabric/common/flogging"
	pb "github.com/hyperledger/fabric/orderer/consensus/dual/grpc"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)
const pkgLogID = "orderer/common/orderer/consensus/dual/dualmodule/server"

var logger *logging.Logger

func init() {
	logger = flogging.MustGetLogger(pkgLogID)
}

type server struct{ oinfo *orderers }

func (s *server) GetPeerInfo(ctx context.Context, in *pb.PeerRequest) (*pb.PeerInfoResponse, error) {
	var credit = float32(s.oinfo.credit)

	return &pb.PeerInfoResponse{SeralizedId: 1, Credit: credit, AmIprimary: s.oinfo.isPrimary}, nil
}
func (s *server) IwantoBePrimary(ctx context.Context, in *pb.IwantToBePrimaryRequest) (*pb.IwantToBePrimaryResponse, error) {
	var suc = false
	if in.Credit > float32(s.oinfo.credit) {
		suc = true
	}
	if in.Credit == float32(s.oinfo.credit) && int(in.SeralizedId) < s.oinfo.seralizeID {
		suc = true
	}
	if suc {
		s.oinfo.isPrimary = !suc
	}
	return &pb.IwantToBePrimaryResponse{Success: suc}, nil
}

func start(port string, oinfo *orderers) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//pb.RegisterHelloServiceServer(s, &server{})
	pb.RegisterBackendServiceServer(s, &server{oinfo})
	s.Serve(lis)
}
func client(address string) (*pb.PeerInfoResponse, error) {

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
	return r, err
}
func bePrimary(address string, oinfo *orderers) (*pb.IwantToBePrimaryResponse, error) {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("did not connect: %v", err)
		//log.Fatalln("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBackendServiceClient(conn)
	r, err := c.IwantoBePrimary(context.Background(), &pb.IwantToBePrimaryRequest{SeralizedId: int32(oinfo.seralizeID), Credit: float32(oinfo.credit)})
	if err != nil {
		logger.Fatal("could not greet: %v", err)
	}
	return r, err
}
