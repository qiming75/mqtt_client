package myhttp

import (
    "bytes"
    "context"
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io"
    "net"
    "net/http"
    "strconv"
    "strings"
    "time"
)

const retryDelay = 10 * time.Second

func JoinMQTT(url, deviceType, mac, tenantid, args string, cdns *net.Dialer) (mqttWillTopic, connectTopic, mqttAddress, mqttUserId, mqttPwd, mqttCliID, reportTopic string, reportInterval float64, subscribeTopics []string, err error) {
    client := &http.Client{
        Timeout: 10 * time.Second, // 设置超时时间
    }
    postData := []byte(fmt.Sprintf(`{"deviceType": "%s", "mac": "%s", "tenantId": "%s", "args": "%s"}`, deviceType, mac, tenantid, args))
    fmt.Println("请求参数为:", string(postData))

    if cdns != nil {
        fmt.Println("自定义HTTP DNS Resolver.")
        http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
            return cdns.DialContext(ctx, network, addr)
        }
        fmt.Println("自定义HTTP DNS Resolver完成.")
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
    if err != nil {
        fmt.Println("创建POST请求失败: ", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    var rsp *http.Response
    for {
        rsp, err = client.Do(req)
        if err == nil && rsp != nil && rsp.StatusCode == http.StatusOK {
            break
        }
        if err != nil {
            fmt.Printf("POST请求失败，将在%v后重试...\n", retryDelay)
            time.Sleep(retryDelay)
        } else {
            fmt.Printf("请求失败，状态码：%d，将在%v后重试...\n", rsp.StatusCode, retryDelay)
            rsp.Body.Close()
            time.Sleep(retryDelay)
        }
    }

    defer rsp.Body.Close()

    body, err := io.ReadAll(rsp.Body)
    if err != nil {
        fmt.Println("解析响应失败: ", err)
        return
    }

    fmt.Println("响应为:")
    fmt.Println(string(body))

    var data map[string]interface{}
    err = json.Unmarshal(body, &data)
    if err != nil {
        fmt.Println("解析 JSON 出错:", err)
        return
    }

    mqttWillTopic = data["mqttWillTopic"].(string)
    connectTopic = data["mqttConnectTopic"].(string)
    cryptographicMqttInfo := data["mqttHost"].(string)
    mqttPORT := data["mqttPort"].(string)
    mqttCliID = data["mqttClientId"].(string)
    mqttTopic := data["mqttTopic"].(string)
    mqttTopicGroup := data["mqttTopicGroup"].([]any)
    reportTopic = data["reportTopic"].(string)
    reportInterval = data["reportInterval"].(float64)
    mqttIp, mqttUserId, mqttPwd, err := getMqttAccount(cryptographicMqttInfo)
    mqttAddress = "tcp://" + mqttIp + ":" + mqttPORT
    subscribeTopics = append(subscribeTopics, mqttTopic)
    for _, t := range mqttTopicGroup {
        subscribeTopics = append(subscribeTopics, t.(string))
    }
    return
}

func getMqttAccount(encodedIp string) (string, string, string, error) {
    // 19进制转10进制
    dec, err := strconv.ParseInt(encodedIp, 19, 64)
    if err != nil {
        return "", "", "", err
    }

    str := strconv.FormatInt(dec, 10)
    fmt.Println("=========================")
    fmt.Println("encodedIp: ", encodedIp)
    fmt.Println("dec: ", dec)
    fmt.Println("str: ", str)

    //标志位
    flag := false
    //下标
    index := 0
    //段标
    fullindex := 0
    sb := ""
    for index < len(str) {
        size := 0
        if !flag {
            nw := string(str[index])
            size, err := strconv.Atoi(nw)
            if err != nil {
                return "", "", "", err
            }
            seg := str[fullindex+1 : fullindex+size+1]
            flag = true
            index = index + size
            sb += seg
        } else {
            flag = false
            index = index + size + 1
            sb += "."
        }
        fullindex = index
    }
    // 计算IP 账号 密码
    tmpStr := sb
    fmt.Println("+++++++++++++++++++++++++++++")
    fmt.Println("tmpStr: ", tmpStr)
    ip := tmpStr[0 : len(tmpStr)-1]

    // 计算 MD5 哈希值
    hash := md5.Sum([]byte(ip))
    // 将字节数组转换为十六进制字符串
    s := hex.EncodeToString(hash[:])
    fmt.Println("*****************************")
    fmt.Println("s: ", s)
    // size := strings.Index(ip, ".")
    size, err := strconv.Atoi(ip[:strings.Index(ip, ".")])
    if err != nil {
        return "", "", "", err
    }

    a := size % 9
    userId := s[a : a+8]
    pwd := s[a+8:]
    return ip, userId, pwd, nil
}
