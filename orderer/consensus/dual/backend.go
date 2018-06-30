package dual

import (
	"net"

	"github.com/hyperledger/fabric/common/flogging"
	pb "github.com/hyperledger/fabric/orderer/consensus/dual/grpc"
	cb "github.com/hyperledger/fabric/protos/common"
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
type clients struct {
	c pb.BackendServiceClient
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
func (s *server) SendChainMessage(ctx context.Context, in *cb.Envelope) (*pb.SendChainMessageResponse, error) {
	var success = false
	s.oc.preOnChan <- in
	success = true
	return &pb.SendChainMessageResponse{Success: success}, nil
}
func (s *server) WrittenChainMessage(ctx context.Context, in *cb.Envelope) (*pb.WrittenChainMessageResponse, error) {
	var success = false
	s.oc.writtenChan <- in
	success = true
	return &pb.WrittenChainMessageResponse{Success: success}, nil
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
func intClient(address string) clients {
	// this should 
	for err==nil{
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			countinue;
		//logger.Fatal("did not connect: %v", err)
		//log.Fatalln("did not connect: %v", err)
		}
		defer conn.Close()
	}
	

	cl := pb.NewBackendServiceClient(conn)
	return clients{c: cl}
}
func (c *clients) SendChain(in *cb.Envelope) bool {
	r, err := c.c.SendChainMessage(context.Background(), in)
	if err != nil {
		logger.Fatal("could not greet: %v", err)
	}
	return r.GetSuccess()
}
func (c *clients) WrittenChain(in *cb.Envelope) bool {
	r, err := c.c.WrittenChainMessage(context.Background(), in)
	if err != nil {
		logger.Fatal("could not greet: %v", err)
	}
	return r.GetSuccess()
}
func (c *clients) cBePrimary(oinfo *orderers) bool {

	r, err := c.c.IwantoBePrimary(context.Background(), &pb.IwantToBePrimaryRequest{SeralizedId: int32(oinfo.seralizeID), Credit: float32(oinfo.credit)})
	if err != nil {
		logger.Fatal("could not greet: %v", err)
	}

	return r.GetSuccess()
}

/*func (c *client) _client(address string) pb.BackendServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("did not connect: %v", err)
		//log.Fatalln("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBackendServiceClient(conn)
	return c
}*/
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
