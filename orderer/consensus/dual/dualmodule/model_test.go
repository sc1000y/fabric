package main

import (
	"testing"
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
func Test2(t *testing.T) {
	var prim = orderers{10, true, 1, 10, false}
	var backup = orderers{10, false, 2, 5, false}
	behavior(5, prim, backup)
}
func behavior(peerNum int, prim orderers, backup orderers) {
	var chain = newChain()
	var chain2 = newChain()
	var oc = newOrderChain()
	var producers [10]peers // 生产者数组S
	//var consumers [2]orderers // 消费者数组
	for i := 0; i < peerNum; i++ {
		producers[i] = peers{i} // 初始化生产者id和Level
		go producers[i].peer(chain, chain2)
	}
	go prim.orderer(chain, oc)
	go backup.orderer(chain2, oc)
	for i := 0; i < (peerNum); i++ {
		print(i)
		<-chain.exitChan // 等待所有生产者和消费者结束退出
	}
}
