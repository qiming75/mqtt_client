package mymqtt

import (
	"encoding/json"
	"fmt"
	"mymqtt/tools"
	"time"
)

func (mmq *MyMQTT) canPublish() error {
	if !mmq.cli.IsConnectionOpen() {
		fmt.Println("-----未处于连接状态，暂不推送数据")
		return fmt.Errorf("connect not open")
	}
	return nil
}

func (mmq *MyMQTT) getBaseInfo() BaseInfo {
	baseInfo := BaseInfo{
		MAC:        mmq.keyMAC,
		MsgType:    "default",
		ClientID:   mmq.clientID,
		TenantID:   mmq.tenantID,
		DeviceType: mmq.deviceType,
		ExecArgs:   mmq.args,
		// ErrorInfo: "",
		TimeStamp: time.Now().Unix(),
	}
	return baseInfo
}

// 推送错误信息
func (mmq *MyMQTT) PublishErrInfo(err error) error {
	if e := mmq.canPublish(); e != nil {
		return e
	}
	fmt.Println("[推送错误信息]")
	baseInfo := mmq.getBaseInfo()
	baseInfo.MsgType = msgTypeOfErrInfo
	info := ErrorInfo{
		BaseInfo: baseInfo,
		Content:  err.Error(),
	}
	jsonData, e := json.Marshal(info)
	if e != nil {
		return e
	}
	if token := mmq.cli.Publish(mmq.errInfoTopic, 0, false, jsonData); token.Wait() && token.Error() != nil {
		fmt.Println("[!!!push fail:] ", token.Error())
		return token.Error()
	}
	fmt.Println("[PUSH TOPIC:]\n", mmq.errInfoTopic)
	fmt.Println("[PUSH MSG:]\n", string(jsonData))
	return nil
}

// 推送设备基本信息
func (mmq *MyMQTT) PublishDevInfo(devInfo tools.DEVInfo) error {
	if e := mmq.canPublish(); e != nil {
		return e
	}
	fmt.Println("[推送设备基本信息]")
	baseInfo := mmq.getBaseInfo()
	baseInfo.MsgType = msgTypeOfDevInfo
	info := DevInfo{
		BaseInfo: baseInfo,
		Content:  devInfo,
	}
	jsonData, e := json.Marshal(info)
	if e != nil {
		return e
	}
	if token := mmq.cli.Publish(mmq.devInfoTopic, 0, false, jsonData); token.Wait() && token.Error() != nil {
		fmt.Println("[!!!push fail:] ", token.Error())
		return token.Error()
	}
	fmt.Println("[PUSH TOPIC:]\n", mmq.devInfoTopic)
	fmt.Println("[PUSH MSG:]\n", string(jsonData))
	return nil
}

// 推送脚本执行结果信息
func (mmq *MyMQTT) PublishScriptExecRet(instruct InstructInfo) error {
	if e := mmq.canPublish(); e != nil {
		return e
	}
	fmt.Println("[推送脚本执行结果信息]")
	topic := instruct.ResultTopic
	baseInfo := mmq.getBaseInfo()
	baseInfo.MsgType = msgTypeOfInstructRet
	info := Instruct{
		BaseInfo: baseInfo,
		Content:  instruct,
	}
	jsonData, e := json.Marshal(info)
	if e != nil {
		return e
	}
	if token := mmq.cli.Publish(topic, 0, false, jsonData); token.Wait() && token.Error() != nil {
		fmt.Println("[!!!push fail:] ", token.Error())
		return token.Error()
	}
	fmt.Println("[PUSH TOPIC:]\n", topic)
	fmt.Println("[PUSH MSG:]\n", string(jsonData))
	return nil
}

// 推送设备连接上线信息
func (mmq *MyMQTT) PublishConnectInfo() error {
	if e := mmq.canPublish(); e != nil {
		fmt.Println("推送设备上线信息失败")
		return e
	}
	fmt.Println("推送设备上线信息")

	baseInfo := mmq.getBaseInfo()
	baseInfo.MsgType = msgTypeOfOnline
	jsonData, e := json.Marshal(baseInfo)
	if e != nil {
		return e
	}

	// onLineMsg := fmt.Sprintf(`{"mac":"%s"}`, mmq.keyMAC)
	// if token := mmq.cli.Publish(mmq.connectInfoTopic, 0, false, []byte(onLineMsg)); token.Wait() && token.Error() != nil {
	if token := mmq.cli.Publish(mmq.connectInfoTopic, 0, false, jsonData); token.Wait() && token.Error() != nil {
		fmt.Println("[!!!push fail:] ", token.Error())
		return token.Error()
	}
	fmt.Println("[PUSH TOPIC:]\n", mmq.connectInfoTopic)
	fmt.Println("[PUSH MSG:]\n", jsonData)
	return nil
}
