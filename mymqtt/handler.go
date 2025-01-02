package mymqtt

import (
	"encoding/json"
	"fmt"
	"mymqtt/tools"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// 所有消息的汇总处理函数
func (mmq *MyMQTT) allMessageHandler(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payLoad := msg.Payload()
	fmt.Printf("GET TOPIC: %s\n", topic)
	fmt.Printf("GET MSG: %s\n", payLoad)

	var instruct Instruct
	err := json.Unmarshal(payLoad, &instruct)
	if err != nil {
		mmq.PublishErrInfo(err)
		return
	}

	fmt.Println("+++++++++++++++++++++++++++++++++")
	fmt.Println(string(payLoad))
	switch instruct.BaseInfo.MsgType {
	case msgTypeOfScriptInstruct:
		mmq.scriptInstructHandler(instruct, string(payLoad))
	default:
		fmt.Println("未知的消息类型: ", instruct.BaseInfo.MsgType)
	}
}

// 脚本执行消息的处理函数
func (mmq *MyMQTT) scriptInstructHandler(instruct Instruct, scriptArgs string) {
	instruct.Content.Success = false
	instruct.Content.ScriptArgs = scriptArgs

	scriptPath := instruct.Content.SavePath
	if mmq.ota.EnableLocal && mmq.ota.LocalPath != "" {
		if len(scriptPath) <= 0 {
			fmt.Println("非法的脚本路径：" + scriptPath)
			return
		}
		items := strings.Split(scriptPath, "/")
		l := len(items)
		if l <= 0 {
			fmt.Println("非法的脚本路径：" + scriptPath)
			return
		}
		if strings.LastIndex(items[l-1], ".sh") < 0 {
			fmt.Println("非法的脚本文件：" + items[l-1] + ", 脚本全路径：" + scriptPath)
			return
		}
		scriptPath = fmt.Sprintf("%s/%s", mmq.ota.LocalPath, items[l-1])
	}

	ret, err := tools.ExecRemoteScript(instruct.Content.RemoteURL, scriptPath)
	if err != nil {
		instruct.Content.Result = err.Error()
	} else {
		instruct.Content.Success = true
		instruct.Content.Result = ret
	}
	mmq.PublishScriptExecRet(instruct.Content) // 发布脚本执行结果
}
