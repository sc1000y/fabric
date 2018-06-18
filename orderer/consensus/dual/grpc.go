package dual

import (
	"sync"
	"time"

	cb "github.com/hyperledger/fabric/protos/common"
	//pb "github.com/hyperledger/fabric/protos/orderer"
)

type broadcaster struct {
	f                int
	broadcastTimeout time.Duration
	msgChans         map[uint64]chan *sendRequest
	closed           sync.WaitGroup
	closedCh         chan struct{}
}

type sendRequest struct {
	msg  *cb.Envelope
	done chan bool
}

func sendMsg() {
	//pb.AtomicBroadcastServer()
}
