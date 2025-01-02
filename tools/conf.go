package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const baseNum = 1000000000

var conf Conf

type Conf struct {
	TenantID         string
	DeviceType       string
	JoinAddress      string
	ErrorInfoTopic   string
	KeepAlive        time.Duration
	PingTimeOut      time.Duration
	Biz              Biz         `json:"biz"`
	DNSResolver      DNSResolver `json:"dns,omitempty"`
	NetworkInterface string      `json:"network_interface,omitempty"`
}

type Biz struct {
	OTA    OTA `json:"ota"`
	Device struct {
		IDPath string `json:"id_path"`
	} `json:"device"`
}

type OTA struct {
	EnableLocal bool   `json:"enable_local_path"`
	LocalPath   string `json:"local_path"`
}

func InitConf(confPath string) Conf {
	data, err := os.ReadFile(confPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := json.Unmarshal(data, &conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conf.KeepAlive *= baseNum
	conf.PingTimeOut *= baseNum

	return conf
}

func GetConf() Conf {
	return conf
}
