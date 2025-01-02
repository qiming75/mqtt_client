package mymqtt

import "mymqtt/tools"

type msgType string

type BaseInfo struct {
	MAC        string   `json:"mac"`
	MsgType    msgType  `json:"msg_type"`
	ClientID   string   `json:"client_id"`
	TenantID   string   `json:"tenant_id"`
	DeviceType string   `json:"device_type"`
	ExecArgs   []string `json:"exec_args"`
	ErrorInfo  string   `json:"error_info"`
	TimeStamp  int64    `json:"timestamp"`
}

type InstructInfo struct {
	ScriptArgs	string `json:"script_args"`
	ScriptID    string `json:"script_id"`
	RemoteURL   string `json:"remote_url"`
	SavePath    string `json:"save_path"`
	ResultTopic string `json:"result_topic"`
	Success     bool   `json:"success"`
	Result      string `json:"result"`
}

type DevInfo struct {
	BaseInfo BaseInfo      `json:"base_info"`
	Content  tools.DEVInfo `json:"content"`
}

type Instruct struct {
	BaseInfo BaseInfo     `json:"base_info"`
	Content  InstructInfo `json:"content"`
}

type ErrorInfo struct {
	BaseInfo BaseInfo `json:"base_info"`
	Content  string   `json:"content"`
}

// 设备上报的上线消息
var msgTypeOfOnline msgType = "online"

// 设备上报的离线消息
var msgTypeOfOffline msgType = "offline"

// 设备上报的错误类信息
var msgTypeOfErrInfo msgType = "error_info"

// 设备上报的设备信息类消息
var msgTypeOfDevInfo msgType = "dev_info"

// 服务端下发的脚本类消息
var msgTypeOfScriptInstruct msgType = "script_instruct"

// 设备上报的脚本执行结果类消息
var msgTypeOfInstructRet msgType = "instruct_ret"
