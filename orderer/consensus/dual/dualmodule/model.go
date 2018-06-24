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
		sendChan:    make(chan *message, 10),
		writtenChan: make(chan *message, 10),
		exitChan:    make(chan bool),
	}
}
func newOrderChain() *orderchain {
	return &orderchain{
		writtenChan: make(chan *message, 10),
		preOnChan:   make(chan *message, 10),
		exitChan:    make(chan bool),
	}
}

type peers struct {
	id int
}
type orderers struct {
	credit       float64
	isPrimary    bool
	seralizeID   int
	mockLag      int
	mockByzatine bool
	//mockBlockChain string
}

func peerBehavior(ch *chain, msg *message) {
	ch.sendChan <- msg
}
func (p peers) peer(ch *chain, ch2 *chain) {
	var msg *message = new(message)
	//msg.configSeq = 1
	//msg.normalMsg = "normalMsg"
	//msg.configMsg = "configMsg"
	//msg.haltMsg = "haltMsg"
	//send msg into channel

	for a := 0; a < 10; a++ {
		//time.Sleep(time.Millisecond * 150)
		msg.configSeq = a
		msg.normalMsg = "normalMsg" + strconv.Itoa(a) + "from" + strconv.Itoa(p.id)
		msg.configMsg = "configMsg" + strconv.Itoa(a) + "from" + strconv.Itoa(p.id)
		msg.haltMsg = "haltMsg"
		sleepTime := rand.Intn(100) + 100
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
		peerBehavior(ch, msg)
		peerBehavior(ch2, msg)
		/*if ch.oinfo.isPrimary {
			if ch.oinfo.credit < ch2.oinfo.credit {
				//ch.oinfo.isPrimary:= <- false
			}
		}*/
	}
	ch.exitChan <- true

}
func increase(blockheight int, ordererCredit float64) float64 {
	var newCredit float64 = 0
	var alpha = 1
	var theta = 0.5
	newCredit = newCredit + float64(alpha)*(1-(float64(ordererCredit)/float64(blockheight)))
	newCredit = float64(ordererCredit) + theta*newCredit
	return newCredit
}
func decrease(blockheight int, ordererCredit float64) float64 {
	var newCredit float64 = 0
	var alpha = 0.6
	var theta = 0.3
	//var beta = 2.0
	//var min = 0.0
	newCredit = newCredit + float64(alpha)*(1-(float64(ordererCredit)/float64(blockheight)))
	newCredit = float64(ordererCredit) + theta*newCredit
	//newCredit = float64(ordererCredit) - (1-theta)*beta*newCredit
	/*if newCredit < min {
		newCredit = min
	}*/
	return newCredit
}
func ordererBehavior(msg *message) {
	//msg := <-ch.sendChan
	fmt.Println("write to block:")
	fmt.Println(msg)
}
func CompareToOppsite(oinfoMine ordererInfo, oinfoOpposite ordererInfo) ordererInfo {
	var isPrimary = false
	if oinfoMine.credit > oinfoOpposite.credit {
		isPrimary = true
	}
	if oinfoMine.credit == oinfoOpposite.credit {
		if oinfoMine.seralizeID > oinfoOpposite.seralizeID {
			isPrimary = true
		}
	}
	oinfoMine.isPrimary = isPrimary
	oinfoOpposite.isPrimary = (!isPrimary)
	fmt.Println("Am I primary", isPrimary)
	return oinfoMine

}
func (o orderers) orderer(ch *chain, oc *orderchain) {
	var timer <-chan time.Time
	var batch []*message
	//batch.Init
	go start(":5005"+strconv.Itoa(o.seralizeID), &o)
	for {
		//var seq = 0 //ch.support.Sequence()
		if o.isPrimary {
			select {
			case msg := <-ch.sendChan:

				fmt.Println("write to block:")
				fmt.Println(msg)
				o.credit = increase(50, o.credit)
				//client("c")
				oc.writtenChan <- msg
				if o.mockByzatine {
					sleepTime := rand.Intn(o.mockLag) + o.mockLag*rand.Intn(o.mockLag)
					time.Sleep(time.Millisecond * time.Duration(sleepTime))
				} else {
					sleepTime := rand.Intn(o.mockLag) + o.mockLag
					time.Sleep(time.Millisecond * time.Duration(sleepTime))
				}

			case <-ch.exitChan:
				fmt.Println(strconv.Itoa(o.seralizeID) + "exit")
				ch.exitChan <- true
				return

			}
		} else {
			select {
			case msg := <-ch.sendChan:
				timer = time.After(time.Second * 2)
				batch = append(batch, msg)
				sleepTime := rand.Intn(o.mockLag) + o.mockLag
				time.Sleep(time.Millisecond * time.Duration(sleepTime))
			//case preonmsg := <-oc.preOnChan:
			//fmt.Println(preonmsg)
			case writtenmsg := <-oc.writtenChan:
				fmt.Println("read from oc:")

				for k, v := range batch {

					if v.normalMsg == writtenmsg.normalMsg && v.configSeq == writtenmsg.configSeq {
						fmt.Print("delete from batch")
						batch = append(batch[:k], batch[k+1:]...)
						break
					}
				}
				/*for e := batch.Front(); e != nil; e = e.Next() {
					//printe.Value
					if e.Value.(message).normalMsg == writtenmsg.normalMsg {
						batch.Remove(e)
						timer = time.After(time.Second * 2)
					}
					//fmt.Println(e.Value)

				}*/
			case <-timer:
				if len(batch) > 0 {
					for _, v := range batch {
						fmt.Println("written by backup service")
						fmt.Println(v)
						add()
						o.credit = increase(getHeight(), o.credit)
					}
				}

				/*if batch {
					for e := batch.Front(); e != nil; e = e.Next() {
						fmt.Println("written by backup service")
						fmt.Println(e.Value)
						o.credit++
					}
				}*/

			}
		}

	}
}

/* Global Variables */
var curCookies int

//do some init for crifanLib
func init() {
	fmt.Println("init something")
	curCookies = 0
	return
}
func add() {
	curCookies++
}
func getHeight() int {
	return curCookies
}

/*func add(Set-Cookie) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "vanyar",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	//http.SetCookie
	//curCookies[0] = curCookies[0] + 1
}*/

/*func (o orderers) orderer(ch *chain, oc *orderchain) {
	var timer <-chan time.Time
	var batch map[int]*message
	fmt.Println("seralize " + strconv.Itoa(o.seralizeID) + " is online")
	//var err error
	for {
		select {

		case msg := <-ch.sendChan:
			//fmt.Println("ok")
			//fmt.Println(msg)
			if o.isPrimary == true { //primary orderer service behavior
				if msg.haltMsg == "haltMsg" {
					//mockWriteToBlock(msg)
					fmt.Println("write to block by " + strconv.Itoa(o.seralizeID))
					fmt.Println(msg)
					msg.haltMsg = "halt this!"
					//ch.writtenChan <- msg
					oc.writtenChan <- msg
					//o.credit++
					sleepTime := rand.Intn(o.mockLag) + o.mockLag
					time.Sleep(time.Millisecond * time.Duration(sleepTime))
				} else {
					fmt.Println("eject msg by " + strconv.Itoa(o.seralizeID))
					fmt.Println(msg)
				}

			}
			if o.isPrimary == false { // backup service behavior
				if msg.haltMsg == "haltMsg" {
					//batch[] += msg
					batch[msg.configSeq] = msg
					timer := time.After(time.Second * 5)
					fmt.Println("timer out start on " + strconv.Itoa(o.seralizeID))
					fmt.Println(msg)
					select {
					case <-timer:

						fmt.Println("write to block by backup service:")
						fmt.Println(msg)
					}
					/*select {
					case <-timer.C: //time out
					case hltMsg := <-ch.writtenChan:
						fmt.println(hltMsg.configSeq)
						fmt.Println("write to block by backup service:")
						fmt.Println(msg)
						msg.haltMsg = "halt this!"
						ch.sendChan <- msg
						//sleepTime := rand.Intn(o.mockLag) + o.mockLag
						//time.Sleep(time.Millisecond * time.Duration(sleepTime))
					}

					//mockWriteToBlock(msg)

					//sleepTime := rand.Intn(100) + 200
					//time.Sleep(time.Millisecond * time.Duration(sleepTime))
				} else {
					fmt.Println("eject msg by backup service:")
					fmt.Println(msg)
				}
			}
		case hltMsg := <-ch.writtenChan:
			fmt.Println("msg is already written on chain")
			fmt.Println(hltMsg)
			timer = nil
		case <-timer:
			fmt.Println("time out")
		case <-ch.exitChan:
			fmt.Println(strconv.Itoa(o.seralizeID) + "exit")
			ch.exitChan <- true
			return

			//fmt.Print(".")
		}
	}
}*/
func mockWriteToBlock(msg *message) {
	fmt.Println(msg)
}

type orderchain struct {
	writtenChan chan *message
	preOnChan   chan *message
	exitChan    chan bool
	blockheight int
}
type message struct {
	configSeq int
	normalMsg string // *cb.Envelope
	configMsg string // *cb.Envelope
	haltMsg   string // *cb.Envelope //dual message
}
type chain struct {
	//support  consensus.ConsenterSupport
	sendChan    chan *message
	writtenChan chan *message
	exitChan    chan bool //struct{}
	oinfo       *ordererInfo
}
type myCredit float64
type ordererInfo struct {
	credit     int
	isPrimary  bool
	seralizeID int
}
