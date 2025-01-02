package mymqtt

import (
	"mymqtt/tools"
	"time"
)

func (mmq *MyMQTT)StartTicker(delay time.Duration) {
	time.Sleep(time.Second * 3)
	for {
		devInfo := tools.GetDEVInfo()
		mmq.PublishDevInfo(devInfo)
		time.Sleep(delay)
	}
}
