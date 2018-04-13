# go-eureka
go语言实现的简单的eureka client



# Start

###### 启动类
```go
package main

import (
    _ "go-eureka"
    "net/http"
    "io"
    "time"
)

func main() {
	http.HandleFunc("/" , func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer,("hello "+time.Now().String()))
	})
	http.ListenAndServe(":8090",nil)
}
```

###### 配置参数  ./resources/go-eureka.ini
```ini
[instance]
ipPrefer=10.0.0
port=8090

[eureka]
addresses=eureka-1.domain.com,eureka-2.domain.com

```

# 配置参数说明
 配置文件路径 ./resources/go-eureka.ini

 注册参数说明：（☆放在\[instance\]下面）

 |   配置参数           | 说明 |
 | ----------          | --- |
 | hostName            |  主机名 |
 | instanceId          |  实例ID |
 | app                 |  应用（服务）名称 |
 | ipAddr              |  ip地址 |
 | vipAddress          |  虚拟Ip地址 |
 | secureVipAddress    |  安全的虚拟Ip的端口 |
 | status              |  状态 UP DOWN STARTING OUT_OF_SERVICE UNKNOWN |
 | port                |  端口 |
 | securePort          |  安全端口 https |
 | homePageUrl         |  主页 |
 | statusPageUrl       |  状态主页 */info |
 | healthCheckUrl      |  健康监控页面 |
 | dataCenterInfo.name |  取值为  MyOwm Amazon |
 | ipPrefer            |  ip地址过滤，前缀，这对多网卡情况 |

 注册中心地址参数 （☆放在\[eureka\]下面）

|    配置参数 | 说明 |
| ---------- | --- |
| addresses  |  注册中心地址，多个以逗号分隔 |
