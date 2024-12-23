# 配置
config.json  
```
{
    // 初始化时，请求该地址以获取基本配置
    "joinaddress": "http://124.192.140.240:7890/docking/iot/join",
    // 运行中的错误信息上报到该topic
    "errorinfotopic": "errorinfo",
    "keepalive": 120,
    "pingtimeout": 30
}
```
---
# 运行
```
./demo [配置文件路径] [tenantID] [deviceType] ...

参数以空格隔开，第一个参数为租户ID，第二个参数为设备类型，后面还可以跟多个其他参数
```
---

# 上报设备信息
```
{
    "base_info": {
        // 上报设备信息的消息类型为dev_info
        "msg_type": "dev_info",
        "mac": "00:16:2E:03:54:8C",
        "client_id": "0001",
        "tenant_id": "0001",
        "dev_type": "1",
        "exec_argc": "["xx", "xx", "xx"]",
        "time_stamp": 1592346708
    },
    "content": {}
}
```
---
# 接收服务端推送的脚本
```
{
    "base_info": {
        // 当消息类型为script_instruct时，设备执行下发的脚本
        msg_type: "script_instruct",
    },
    "content": {
        // 脚本执行结果会原样携带该id
        "script_id": "0001",
        // 拉取该脚本的url
        "remote_url": "http://xxx/xxx",
        // 拉取到的脚本存到设备该位置下
        "save_path": "./xx.sh",
        // 脚本执行结果会推送到该topic
        "result_topic": "xxx"
    }
}
```
---
# 上报脚本执行结果
```
{
    "base_info": {
        // 脚本执行结果的消息类型为instruct_ret
        "msg_type": "instruct_ret",
        "mac": "00:16:2E:03:54:8C",
        "client_id": "0001",
        "tenant_id": "0001",
        "dev_type": "1",
        "exec_argc": "["xx", "xx", "xx"]",
        "time_stamp": 1592346708
    },
    "content": {
        "script_args": "xxxxx",
        // 该id同接受执行脚本指令里的id
        "script_id": "0001",
        // 脚本是否成功执行
        "success": true,
        // 脚本执行结果(若脚本执行成功这里是执行结果，若脚本执行失败这里是错误信息)
        "result": "xxx",
    }
}
```
---
# 推送错误消息
当本程序在运行过程中出现异常错误时，会将错误信息推送到topic(errorinfo)
```
{
    "base_info": {
        // 错误信息的消息类型为error_info
        "msg_type": "error_info",
        "mac": "00:16:2E:03:54:8C",
        "client_id": "0001",
        "tenant_id": "0001",
        "dev_type": "1",
        "exec_argc": "["xx", "xx", "xx"]",
        "time_stamp": 1592346708
    },
    // 错误信息
    "content": "xxx"
}
```
---