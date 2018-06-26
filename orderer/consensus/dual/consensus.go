/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dual

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/common/flogging"
	pb "github.com/hyperledger/fabric/orderer/consensus/dual/grpc"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/op/go-logging"
)

const pkgLogID = "orderer/consensus/dual"

var logger *logging.Logger

func init() {
	logger = flogging.MustGetLogger(pkgLogID)
}

//credit for oderer peers

type ordererInfo struct {
	credit     float64
	isPrimary  bool
	seralizeID int
	port       int
	host       string
}

//NewOrdererInfo is for new an ordererinfo
/*func NewOrdererInfo(credit myCredit, isPrimary bool, seralizeID int) ordererInfo {
	return ordererInfo{credit, isPrimary, seralizeID}
}

//type isPrimary bool

//CalculateCredit to result
func CalculateCredit(credit myCredit) myCredit {
	credit++
	return credit
}*/
func SendHaltMSG(message *message) {
	//TODO
	//seq = ch.support.Sequence()
	var seq = message.configSeq
	var haltMsg = "halt" + strconv.Itoa(int(seq)) + "msg"
	//var message = message{seq, nil, nil, haltMsg}
	//msg := <-ch.sendChan
	message.haltMsg = haltMsg
	//ch.sendChan <- msg
	//GetBroadcastClient()

}
func CheckIfHalt(haltMsg *cb.Envelope) bool {
	var haltFlag = false
	//TODO
	return haltFlag
}

//calcuating the speculation time of when the message will put into chain, and send to all network
func preOnChainNotice() {
	//TODO
}

//CompareToOppsite to define who is better
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
func makeEnvelope(in *cb.Envelope) *pb.Envelope {
	return &pb.Envelope{Payload: in.GetPayload(), Signature: in.GetSignature()}
}

/*func checkEnvelope(in *cb.Envelope, c *pb.Envelope) bool {
	success := false

	if in.Payload() == c.GetPayload() && in.Signature() == c.GetSignature() {
		success = true
	}
	return success
}*/
func (ch *chain) main() {
	var timer <-chan time.Time
	var err error
	var o = orderers{credit: ch.oinfo.credit, isPrimary: ch.oinfo.isPrimary, seralizeID: ch.oinfo.seralizeID}
	var oc = newOrderChain()
	go start(":"+strconv.Itoa(ch.oinfo.port), &o, oc) //start gRPC backend
	var addr = ""
	if o.isPrimary {
		addr = ch.config.Backup.Host + ":" + strconv.Itoa(ch.config.Backup.Port)
	} else {

		addr = ch.config.Priamy.Host + ":" + strconv.Itoa(ch.config.Priamy.Port)
	}
	client := intClient(addr) //need to be sync
	for {
		seq := ch.support.Sequence()

		err = nil
		select {
		case msg := <-ch.sendChan:
			if msg.configMsg == nil && msg.normalMsg != nil {
				// NormalMsg
				if msg.configSeq < seq {
					_, err = ch.support.ProcessNormalMsg(msg.normalMsg)
					if err != nil {
						logger.Warningf("Discarding bad normal message: %s", err)
						continue
					}
				}
				//msg.configMsg.String
				//cb.Envelope()

				preOnChainNotice()
				batches, _ := ch.support.BlockCutter().Ordered(msg.normalMsg)
				if len(batches) == 0 && timer == nil {
					timer = time.After(ch.support.SharedConfig().BatchTimeout())
					continue
				}
				for _, batch := range batches {
					block := ch.support.CreateNextBlock(batch)
					ch.support.WriteBlock(block, nil)

				}

				client.SendChain(makeEnvelope(msg.normalMsg))
				//SendHaltMSG(msg)
				ch.sendChan <- msg

				if len(batches) > 0 {
					timer = nil
				}
			} else if msg.configMsg != nil {
				// ConfigMsg
				if msg.configSeq < seq {
					msg.configMsg, _, err = ch.support.ProcessConfigMsg(msg.configMsg)
					if err != nil {
						logger.Warningf("Discarding bad config message: %s", err)
						continue
					}
				}
				batch := ch.support.BlockCutter().Cut()
				if batch != nil {
					block := ch.support.CreateNextBlock(batch)
					ch.support.WriteBlock(block, nil)
				}

				block := ch.support.CreateNextBlock([]*cb.Envelope{msg.configMsg})
				ch.support.WriteConfigBlock(block, nil)
				timer = nil
			} else {
				//haltMsg
				//if CheckIfHalt(msg.haltMsg) {
				//	timer = nil //halt the next block create
				//}

			}
		case <-timer:
			//clear the timer
			timer = nil
			//primary running
			batch := ch.support.BlockCutter().Cut()
			if len(batch) == 0 {
				logger.Warningf("Batch timer expired with no pending requests, this might indicate a bug")
				continue
			}
			logger.Debugf("Batch timer expired, creating block")
			block := ch.support.CreateNextBlock(batch)
			ch.support.WriteBlock(block, nil)

			//SendHaltMSG(msg)
			//ch.sendChan <- msg
		case <-ch.exitChan:
			logger.Debugf("Exiting")
			return
		}
	}
}
