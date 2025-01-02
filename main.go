package main

import (
	"fmt"
	"mymqtt/myhttp"
	"mymqtt/mymqtt"
	"mymqtt/tools"
	"net"
	"os"
	"strings"
	"time"
)

var confPath = "./conf.json"

func main() {

	// 使用指定的配置文件
	if len(os.Args) >= 1 {
		confPath = os.Args[1]
	}

	// 读取配置文件
	conf := tools.InitConf(confPath)

	var args []string
	// 解析命令行参数
	if len(os.Args) >= 4 {
		conf.TenantID = os.Args[2]
		conf.DeviceType = os.Args[3]
		args = os.Args[1:]
	}

	var cdns *net.Dialer

	if conf.DNSResolver.Enable {
		cdns = tools.CustomDialer(conf.DNSResolver)
	}

	// 获取设备MAC地址
	keyMAC := tools.GetKeyMAC()

	mqttWillTopic, connectTopic, mqttAddress, mqttUserId, mqttPwd, mqttCliID, reportTopic, reportInterval, subscribeTopics, err := myhttp.JoinMQTT(
		conf.JoinAddress, conf.DeviceType, keyMAC, conf.TenantID, strings.Join(args, ","), cdns,
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 将设备ID写入文件
	writeKeyMAC := strings.ReplaceAll(keyMAC, ":", "")
	devID := fmt.Sprintf("%s-%s-%s", conf.TenantID, conf.DeviceType, writeKeyMAC)
	tools.UpdateTenantID2File(devID, fmt.Sprintf("%s/.id", conf.Biz.Device.IDPath))

	myMqttCli, err := mymqtt.NewMyMQTT(
		keyMAC,
		conf.TenantID,
		conf.DeviceType,
		args,
		mqttAddress, mqttUserId, mqttPwd, mqttCliID,
		reportTopic, conf.ErrorInfoTopic,
		mqttWillTopic, connectTopic,
		conf.KeepAlive, conf.PingTimeOut,
		subscribeTopics,
		cdns,
		conf.Biz.OTA,
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	myMqttCli.StartTicker(time.Duration(reportInterval) * time.Minute)
}

// $env:GOOS="linux"
// $env:GOARCH="mipsle"
// $env:GOMIPS="softfloat"

// go build -ldflags="-s -w" -o ttt

// tenantID := "d5b0cca3"
// deviceType := "tgg58a"
