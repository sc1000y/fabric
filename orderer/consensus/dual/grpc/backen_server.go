package grpc

import (
	"net"

	"github.com/hyperledger/fabric/common/flogging"
	//pb "github.com/hyperledger/fabric/orderer/consensus/dual/grpc"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)
const pkgLogID = "orderer/common/orderer/consensus/dual/grpc/server"

var logger *logging.Logger

func init() {
	logger = flogging.MustGetLogger(pkgLogID)
}

type server struct{}

func (s *server) GetPeerInfo(ctx context.Context, in *PeerRequest) (*PeerInfoResponse, error) {
	return &PeerInfoResponse{SeralizedId: 1, Credit: 20.0, AmIprimary: false}, nil
}
func (s *server) IwantoBePrimary(ctx context.Context, in *IwantToBePrimaryRequest) (*IwantToBePrimaryResponse, error) {
	return &IwantToBePrimaryResponse{Success: false}, nil
}

/*func (s *server) SendChainMessage(ctx context.Context, in *pb.Envelope) (*SendChainMessageResponse, error) {
	var success = false
	s.oc.preOnChan <- in
	success = true
	return &SendChainMessageResponse{Success: success}, nil
}
func (s *server) WrittenChainMessage(ctx context.Context, in *pb.Envelope) (*WrittenChainMessageResponse, error) {
	var success = false
	s.oc.writtenChan <- in
	success = true
	return &WrittenChainMessageResponse{Success: success}, nil
}*/
func start(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//pb.RegisterHelloServiceServer(s, &server{})
	//RegisterBackendServiceServer(s, &server{})
	s.Serve(lis)
}
