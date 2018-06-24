package main

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/hyperledger/fabric/orderer/consensus/dual/grpc"
	"google.golang.org/grpc"
)

/*func Test1(t *testing.T) {
	var chain = newChain()
	//var waitQ = make(chan bool)
	const producerCount int = 5           // 生产者数量
	const consumerCount int = 2           // 消费者数量
	var producers [producerCount]peers    // 生产者数组
	var consumers [consumerCount]orderers // 消费者数组
	for i := 0; i < producerCount; i++ {
		producers[i] = peers{i} // 初始化生产者id和Level
		go producers[i].peer(chain)
	}
	consumers[0] = orderers{10, true, 1, 10, false} //
	go consumers[0].orderer(chain)
	consumers[1] = orderers{10, false, 2, 5, false} //
	go consumers[1].orderer(chain)
	/*for i := 0; i < consumerCount; i++ {
		var primaryFlag = false
		if i == 0 {
			primaryFlag = true
		}
		consumers[i] = orderers{10, primaryFlag, i} //
		go consumers[i].orderer(chain)
	}
	for i := 0; i < (producerCount); i++ {
		print(i)
		<-chain.exitChan // 等待所有生产者和消费者结束退出
	}

}/*/
func Test1(t *testing.T) {
	var prim = orderers{16, true, 1, 4, false}
	var backup = orderers{10, false, 2, 5, false}
	behavior(5, prim, backup)
	res, err := client(address)
	if err != nil {
		logger.Fatal("could not greet: %v", err)
	}
	fmt.Printf("Greeting from 1: %f", res.GetCredit())
	res2, err2 := client("localhost:50052")
	if err2 != nil {
		logger.Fatal("could not greet: %v", err)
	}
	fmt.Printf("Greeting from 2: %f", res2.GetCredit())
	//mockClient(address)
}
func Test2(t *testing.T) {
	var prim = orderers{10, true, 1, 400, false}
	var backup = orderers{10, false, 2, 5, false}
	behavior(5, prim, backup)
}
func Test3(t *testing.T) {
	var prim = orderers{10, true, 1, 100, true} //byzantine peer
	var backup = orderers{10, false, 2, 5, false}

	behavior(5, prim, backup)
	res, err := client(address)
	if err != nil {
		logger.Fatal("could not greet: %v", err)
	}
	fmt.Printf("Greeting from 1: %f and primary is %t", res.GetCredit(), res.GetAmIprimary())
	res2, err2 := client("localhost:50052")
	if err2 != nil {
		logger.Fatal("could not greet: %v", err)
	}
	fmt.Printf("Greeting from 2: %f and primary is %t", res2.GetCredit(), res2.GetAmIprimary())
}
func TestCalculate(t *testing.T) {
	var height = 15
	var credit = 10.0
	fmt.Println(increase(height, credit))
	height = 20
	credit = 10.0
	fmt.Println(increase(height, credit))
	height = 50
	credit = 15.0
	fmt.Println(increase(height, credit))
	height = 80
	credit = 15.0
	fmt.Println(increase(height, credit))
	height = 150
	credit = 15.0
	fmt.Println(increase(height, credit))
	height = 500
	credit = 20.0
	fmt.Println(increase(height, credit))
}
func TestDecrease(t *testing.T) {
	var height = 150
	var credit = 50.0
	fmt.Println(decrease(height, credit))
	height = 500
	credit = 80.0
	fmt.Println(decrease(height, credit))
	height = 500
	credit = 105.0
	fmt.Println(decrease(height, credit))
}
func behavior(peerNum int, prim orderers, backup orderers) {
	var chain = newChain()
	var chain2 = newChain()
	var oc = newOrderChain()
	//curCookies
	//init()
	//init()
	var producers [10]peers // 生产者数组S
	//var consumers [2]orderers // 消费者数组
	for i := 0; i < peerNum; i++ {
		producers[i] = peers{i} // 初始化生产者id和Level
		go producers[i].peer(chain, chain2)
	}
	go prim.orderer(chain, oc)
	go backup.orderer(chain2, oc)
	for i := 0; i < (peerNum); i++ {
		println("curcookie is", curCookies)
		println(i)
		<-chain.exitChan // 等待所有生产者和消费者结束退出
	}

}

const (
	address = "localhost:50051"
)

func mockClient(address string) {

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
	fmt.Printf("Greeting: %f", r.GetCredit())

}
