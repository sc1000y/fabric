package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("hello world . 你好 世界！")
	/*var chain = newChain()
	//var waitQ = make(chan bool)
	const producerCount int = 5           // 生产者数量
	const consumerCount int = 2           // 消费者数量
	var producers [producerCount]peers    // 生产者数组
	var consumers [consumerCount]orderers // 消费者数组
	for i := 0; i < producerCount; i++ {
		producers[i] = peers{i} // 初始化生产者id和Level
		go producers[i].peer(chain)
	}
	for i := 0; i < consumerCount; i++ {
		var primaryFlag = false
		if i == 0 {
			primaryFlag = true
		}
		consumers[i] = orderers{10, primaryFlag, i} // 初始化生产者id和Level
		go consumers[i].orderer(chain)
	}*/

}

func newChain() *chain {
	return &chain{
		//support:  support,
		sendChan: make(chan *message, 10),
		exitChan: make(chan bool),
	}
}

type peers struct {
	id int
}
type orderers struct {
	credit     int
	isPrimary  bool
	seralizeID int
	//mockBlockChain string
}

func peerBehavior(ch *chain, msg *message) {
	ch.sendChan <- msg
}
func (p peers) peer(ch *chain) {
	var msg *message = new(message)
	//msg.configSeq = 1
	msg.normalMsg = "normalMsg"
	msg.configMsg = "configMsg"
	msg.haltMsg = "haltMsg"
	//send msg into channel
	for a := 0; a < 10; a++ {
		//time.Sleep(time.Millisecond * 150)
		msg.configSeq = a
		msg.normalMsg = "normalMsg" + strconv.Itoa(a) + "from" + strconv.Itoa(p.id)
		msg.configMsg = "configMsg" + strconv.Itoa(a) + "from" + strconv.Itoa(p.id)
		sleepTime := rand.Intn(100) + 100
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
		peerBehavior(ch, msg)
	}
	ch.exitChan <- true

}

func ordererBehavior(msg *message) {
	//msg := <-ch.sendChan
	fmt.Println("write to block:")
	fmt.Println(msg)
}
func (o orderers) orderer(ch *chain) {
	for {
		select {

		case msg := <-ch.sendChan:
			//fmt.Println("ok")
			//fmt.Println(msg)
			if msg.haltMsg == "haltMsg" {
				mockWriteToBlock(msg)
				msg.haltMsg = "halt this!"
				ch.sendChan <- msg
				//sleepTime := rand.Intn(100) + 200
				//time.Sleep(time.Millisecond * time.Duration(sleepTime))
			} else {
				fmt.Println("eject msg:")
				fmt.Println(msg)
			}

		case <-ch.exitChan:
			fmt.Println(strconv.Itoa(o.seralizeID) + "exit")
			ch.exitChan <- true
			return

			//fmt.Print(".")
		}
	}
}
func mockWriteToBlock(msg *message) {
	fmt.Println(msg)
}

type message struct {
	configSeq int
	normalMsg string // *cb.Envelope
	configMsg string // *cb.Envelope
	haltMsg   string // *cb.Envelope //dual message
}
type chain struct {
	//support  consensus.ConsenterSupport
	sendChan chan *message
	exitChan chan bool //struct{}
	oinfo    ordererInfo
}
type myCredit int
type ordererInfo struct {
	credit     int
	isPrimary  bool
	seralizeID int
}
