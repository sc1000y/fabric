package dual

type orderers struct {
	credit     float64
	isPrimary  bool
	seralizeID int
	//mockLag      int
	//mockByzatine bool
	//mockBlockChain string
}
type orderchain struct {
	writtenChan chan *message
	preOnChan   chan *message
	exitChan    chan bool
}

func newOrderChain() *orderchain {
	return &orderchain{
		writtenChan: make(chan *message, 10),
		preOnChan:   make(chan *message, 10),
		exitChan:    make(chan bool),
	}
}
