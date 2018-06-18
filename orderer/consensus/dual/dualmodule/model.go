package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("hello world . 你好 世界！")
	//var chain = newChain()
	//var waitQ = make(chan bool)
	//go peer(chain)
	//go orderer(chain)

}

func newChain() *chain {
	return &chain{
		//support:  support,
		sendChan: make(chan *message, 10),
		exitChan: make(chan struct{}),
	}
}

type peers struct {
	id int
}
type orderers struct {
	id int
}

func peerBehavior(ch *chain, msg *message) {
	ch.sendChan <- msg
}
func (p peers) peer(ch *chain, waitQ chan bool) {
	var msg *message = new(message)
	//msg.configSeq = 1
	msg.normalMsg = "normalMsg"
	msg.configMsg = "configMsg"
	msg.haltMsg = "haltMsg"
	//send msg into channel
	for a := 0; a < 10; a++ {
		time.Sleep(time.Millisecond * 150)
		msg.configSeq = a
		msg.normalMsg = msg.normalMsg + strconv.Itoa(a)
		msg.configMsg = msg.configMsg + strconv.Itoa(a)
		sleepTime := rand.Intn(1000) + 1000
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
		ch.sendChan <- msg
	}
	//ch.exitChan <-
	//ch.sendChan <- msg
	waitQ <- true
}
func ordererBehavior(msg *message) {
	//msg := <-ch.sendChan
	fmt.Println(msg)
}
func (o orderers) orderer(ch *chain, waitQ chan bool) {
	for {
		select {

		case msg := <-ch.sendChan:
			fmt.Println("ok")
			ordererBehavior(msg)
			waitQ <- true
			sleepTime := rand.Intn(1000) + 2000
			time.Sleep(time.Millisecond * time.Duration(sleepTime))
		default:
			fmt.Print("nothing in channel")
		}
	}
}

type message struct {
	configSeq int
	normalMsg string
	configMsg string
	haltMsg   string // *cb.Envelope //dual message
}
type chain struct {
	//support  consensus.ConsenterSupport
	sendChan chan *message
	exitChan chan struct{}
	oinfo    ordererInfo
}
type myCredit int
type ordererInfo struct {
	credit     myCredit
	isPrimary  bool
	seralizeID int
}
