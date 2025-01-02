package mymqtt

import (
	"encoding/json"
	"fmt"
	"mymqtt/tools"
	"net"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MyMQTT struct {
	keyMAC           string
	tenantID         string
	deviceType       string
	args             []string
	clientID         string
	cli              mqtt.Client
	errInfoTopic     string
	devInfoTopic     string
	connectInfoTopic string
	subscribeTopics  []string
	ota              tools.OTA
}

var myMQTTObj *MyMQTT

func NewMyMQTT(
	keyMAC,
	tenantID,
	deviceType string,
	args []string,
	serAddress, userID, password, clientID, devInfoTopic, errInfoTopic, mqttWillTopic, connectTopic string,
	keepAlive, pingTimeOut time.Duration,
	subscribeTopics []string,
	cdns *net.Dialer,
	otaConf tools.OTA,
) (myMQTT *MyMQTT, err error) {
	myMQTT = &MyMQTT{}
	myMQTT.clientID = clientID
	myMQTT.devInfoTopic = devInfoTopic
	myMQTT.errInfoTopic = errInfoTopic
	myMQTT.keyMAC = keyMAC
	myMQTT.tenantID = tenantID
	myMQTT.deviceType = deviceType
	myMQTT.args = args
	myMQTT.connectInfoTopic = connectTopic
	myMQTT.subscribeTopics = subscribeTopics
	myMQTT.ota = otaConf
	myMQTTObj = myMQTT

	// willMsg := fmt.Sprintf(`{"mac":"%s"}`, keyMAC)

	baseInfo := myMQTT.getBaseInfo()
	baseInfo.MsgType = msgTypeOfOffline
	willMsg, err := json.Marshal(baseInfo)
	if err != nil {
		return nil, err
	}

	// 初始化参数
	opts := mqtt.NewClientOptions().AddBroker(serAddress).SetClientID(clientID).SetUsername(userID).SetPassword(password)
	opts.SetKeepAlive(keepAlive).SetPingTimeout(pingTimeOut)
	opts.SetAutoReconnect(true).SetConnectRetry(true).SetMaxReconnectInterval(10 * time.Second).SetOnConnectHandler(_onConnect).SetConnectionLostHandler(_onConnectionLost).SetReconnectingHandler(_onReconnect)
	opts.SetConnectRetry(true)
	opts.SetWill(mqttWillTopic, string(willMsg), 0, false)
	if cdns != nil {
		opts.SetDialer(cdns)
	}
	// opts.SetDefaultPublishHandler(messagHandler)
	fmt.Println("***参数初始化完毕***")

	// 实例化客户端
	myMQTT.cli = mqtt.NewClient(opts)
	if token := myMQTT.cli.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
		return
	}
	fmt.Println("***MQTT客户端实例化完毕***")

	myMQTT._showConnectStatus()

	return
}

func (mm *MyMQTT) _showConnectStatus() {
	go func() {
		for {
			if !mm.cli.IsConnected() || !mm.cli.IsConnectionOpen() {
				fmt.Println("未连接MQTT")
			} else {
				fmt.Println("已连接MQTT")
			}
			time.Sleep(3 * time.Second)
		}
	}()
}

// 订阅主题
func (mm *MyMQTT) _subscribe() error {
	fmt.Println(mm.subscribeTopics)
	for _, topic := range mm.subscribeTopics {
		token := mm.cli.Subscribe(topic, 0, mm.allMessageHandler)
		if token.Wait() && token.Error() != nil {
			return token.Error()
		}
		fmt.Println("订阅主题: ", topic)
	}
	return nil
}

func _onConnect(c mqtt.Client) {
	fmt.Println("***建立连接***")

	err := myMQTTObj._subscribe()
	if err != nil {
		return
	}

	err = myMQTTObj.PublishConnectInfo()
	if err != nil {
		fmt.Println("发布连接信息失败:", err)
	}
}

func _onConnectionLost(c mqtt.Client, err error) {
	fmt.Println("***连接断开***")
	fmt.Println(err)
}

func _onReconnect(c mqtt.Client, opts *mqtt.ClientOptions) {
	fmt.Println("***尝试重新建立连接...***")
	if token := c.Connect(); token.Error() != nil {
		fmt.Println("重新连接失败:", token.Error())
		return
	}

	if !c.IsConnected() || !c.IsConnectionOpen() {
		fmt.Println("连接还未建立")
		return
	}
}
