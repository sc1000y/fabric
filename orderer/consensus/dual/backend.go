package dual

import (
	"net"

	"github.com/hyperledger/fabric/common/flogging"
	pb "github.com/hyperledger/fabric/orderer/consensus/dual/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func init() {
	logger = flogging.MustGetLogger(pkgLogID)
}

type server struct {
	oinfo *orderers
	oc    *orderchain
}

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

func start(port string, oinfo *orderers, oc *orderchain) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//oinfo orderers:={Credit:oinfoCfg.Credit,isPrimary:oinfoCfg.}
	//pb.RegisterHelloServiceServer(s, &server{})
	pb.RegisterBackendServiceServer(s, &server{oinfo, oc})
	s.Serve(lis)

}
func _client(address string) pb.BackendServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("did not connect: %v", err)
		//log.Fatalln("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBackendServiceClient(conn)
	return c
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
