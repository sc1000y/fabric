package dual

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//"github.com/hyperledger/fabric/core/util" struct 结构体必须大写
type oinfoCfg struct {
	Credit     float64 `yaml:"credit"`
	SeralizeID int     `yaml:"seralizeID"`
	Host       string  `yaml:"host"`
	Port       int     `yaml:"port"`
}

//ChannelCfg is config from config.yaml
type ChannelCfg struct {
	IsPrimary bool     `yaml:"isPrimary"`
	Priamy    oinfoCfg `yaml:"primary"`
	Backup    oinfoCfg `yaml:"backup"`
}

func getConfig() ChannelCfg {
	//var err error
	var configFile, err = ioutil.ReadFile("config.yaml")
	//fmt.Println(string(configFile))
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}
	channelCfg := ChannelCfg{}
	//var oinfoCfg oinfoCfg
	err = yaml.Unmarshal(configFile, &channelCfg)
	fmt.Println(channelCfg)
	//fmt.Println(err)
	//ChannelCfg.priamy
	return channelCfg
}
func round(val float64) uint64 {
	if val < 0 {
		return uint64(val - 0.5)
	}
	return uint64(val + 0.5)
}
