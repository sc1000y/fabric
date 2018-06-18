package main

import "testing"

func Test1(t *testing.T) {
	var chain = newChain()
	var waitQ = make(chan bool)
	const producerCount int = 5           // 生产者数量
	const consumerCount int = 2           // 消费者数量
	var producers [producerCount]peers    // 生产者数组
	var consumers [consumerCount]orderers // 消费者数组
	for i := 0; i < producerCount; i++ {
		producers[i] = peers{i} // 初始化生产者id和Level
		go producers[i].peer(chain, waitQ)
	}
	for i := 0; i < consumerCount; i++ {
		consumers[i] = orderers{i} // 初始化生产者id和Level
		go consumers[i].orderer(chain, waitQ)
	}
	for i := 0; i < (producerCount + consumerCount); i++ {
		<-waitQ // 等待所有生产者和消费者结束退出
	}
}
