package dual

import pb "github.com/hyperledger/fabric/orderer/consensus/dual/grpc"

type orderers struct {
	credit     float64
	isPrimary  bool
	seralizeID int
	//mockLag      int
	//mockByzatine bool
	//mockBlockChain string
}
type orderchain struct {
	writtenChan chan *pb.Envelope
	preOnChan   chan *pb.Envelope
	exitChan    chan bool
}

func newOrderChain() *orderchain {
	return &orderchain{
		writtenChan: make(chan *pb.Envelope, 10),
		preOnChan:   make(chan *pb.Envelope, 10),
		exitChan:    make(chan bool),
	}
}
