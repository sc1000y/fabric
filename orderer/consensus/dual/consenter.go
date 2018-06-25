package dual

import (
	"fmt"

	"github.com/hyperledger/fabric/orderer/consensus"
	cb "github.com/hyperledger/fabric/protos/common"
)

type consenter struct {
	config ChannelCfg
}

// New creates a new consenter for the solo consensus scheme.
// The solo consensus scheme is very simple, and allows only one consenter for a given chain (this process).
// It accepts messages being delivered via Order/Configure, orders them, and then uses the blockcutter to form the messages
// into blocks before writing to the given ledger
func New() consensus.Consenter {
	return &consenter{config: getConfig()}
}

type chain struct {
	support  consensus.ConsenterSupport
	sendChan chan *message
	exitChan chan struct{}
	oinfo    ordererInfo
	config   ChannelCfg
}

type message struct {
	configSeq uint64
	normalMsg *cb.Envelope "github.com/hyperledger/fabric/peer/common/broadcastclient"
	configMsg *cb.Envelope
	haltMsg   string // *cb.Envelope //dual message
}

func (dual *consenter) HandleChain(support consensus.ConsenterSupport, metadata *cb.Metadata) (consensus.Chain, error) {

	return newChain(dual, support), nil
}

func newChain(dual *consenter, support consensus.ConsenterSupport) *chain {
	ch := &chain{
		support:  support,
		sendChan: make(chan *message),
		exitChan: make(chan struct{}),
		config:   dual.config,
	}
	if dual.config.IsPrimary {
		ch.oinfo.credit = dual.config.Priamy.Credit
		ch.oinfo.isPrimary = dual.config.IsPrimary
		ch.oinfo.seralizeID = dual.config.Priamy.SeralizeID
	} else {
		ch.oinfo.credit = dual.config.Backup.Credit
		ch.oinfo.isPrimary = dual.config.IsPrimary
		ch.oinfo.seralizeID = dual.config.Backup.SeralizeID
	}
	return ch
}

func (ch *chain) Start() {
	go ch.main()
}

func (ch *chain) Halt() {
	select {
	case <-ch.exitChan:
		// Allow multiple halts without panic
	default:
		close(ch.exitChan)
	}
}

func (ch *chain) WaitReady() error {
	return nil
}

// Order accepts normal messages for ordering
func (ch *chain) Order(env *cb.Envelope, configSeq uint64) error {
	select {
	case ch.sendChan <- &message{
		configSeq: configSeq,
		normalMsg: env,
	}:
		return nil
	case <-ch.exitChan:
		return fmt.Errorf("Exiting")
	}
}

// Configure accepts configuration update messages for ordering
func (ch *chain) Configure(config *cb.Envelope, configSeq uint64) error {
	select {
	case ch.sendChan <- &message{
		configSeq: configSeq,
		configMsg: config,
	}:
		return nil
	case <-ch.exitChan:
		return fmt.Errorf("Exiting")
	}
}

// Errored only closes on exit
func (ch *chain) Errored() <-chan struct{} {
	return ch.exitChan
}
