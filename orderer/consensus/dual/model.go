package dual

import (
	cb "github.com/hyperledger/fabric/protos/common"
)

type orderers struct {
	credit     float64
	isPrimary  bool
	seralizeID int
	//mockLag      int
	//mockByzatine bool
	//mockBlockChain string
}
type orderchain struct {
	writtenChan chan *cb.Envelope
	preOnChan   chan *cb.Envelope
	exitChan    chan bool
}

func newOrderChain() *orderchain {
	return &orderchain{
		writtenChan: make(chan *cb.Envelope, 10),
		preOnChan:   make(chan *cb.Envelope, 10),
		exitChan:    make(chan bool),
	}
}
