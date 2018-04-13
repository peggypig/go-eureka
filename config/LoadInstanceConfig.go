/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/12 10:53.
 */
package config

import (
	"os"
	"go-eureka/glog"
	"net"
	"strings"
	"github.com/Unknwon/goconfig"
	"strconv"
	"go-eureka/common"
	"go-eureka/model"
)

const  configPath = "./resources/go-eureka.ini"

var config model.InstanceConfig

/**
 * 加载instance的配置参数
 */
func init() {
	isExits , err :=pathExists(configPath)
	if err != nil{
		glog.Println(err)
	}else {
		loadEureka()
		setDefaultConfig()
		if isExits {
			//文件存在
			loadConfig()
		}else {
			//文件不存在
			glog.Println("the file ./resources/go-eureka.ini  is not exits ,go-eureka will register with default  value")
		}
	}
	glog.Println("eurekaConfig:",eurekaConfig)
	glog.Println("instanceConfig:",config)
}

/**
 * 判断文件是否存在
 */
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
 * 获取instance的参数
 */
func GetInstanceConfig() (model.InstanceConfig) {
	return config
}

/**
 * 加载默认配置
 */
func setDefaultConfig()  {
	//默认主机名
	config.HostName , _ = os.Hostname()

	//默认的APP名称
	app:=config.HostName+":GO-EUREKA"
	config.App = strings.ToUpper(app)


	//默认InstanceId
	config.InstanceId = strings.ToUpper(app)

	//随机寻找一个非循环地址的IP
	ipAddrs ,_ := net.InterfaceAddrs()
	for _, address := range ipAddrs{
		if ipnet , ok :=address.(*net.IPNet);ok && !ipnet.IP.IsLoopback() && common.IsIpv4(ipnet.IP.String()){
			config.IpAddr = ipnet.IP.String()
			break
		}
	}
	//状态  必须为 UP DOWN STARTING OUT_OF_SERVICE UNKNOWN 之一
	config.Status = "UP"
	//必须为 Amazon 或 MyOwn
	config.DataCenterInfo.Name = "MyOwn"
	//默认端口
	config.Port  = "80"

	//健康检测的地址
	config.HealthCheckUrl = "http://"+config.IpAddr+":"+config.Port+"/health"
	//主页地址
	config.HomePageUrl ="http://"+config.IpAddr+":"+config.Port+"/"
	//状态信息地址
	config.StatusPageUrl ="http://"+config.IpAddr+":"+config.Port+"/info"

	config.VipAddress  = strings.ToLower(app)

	config.SecureVipAddress = strings.ToLower(app)

	config.SecurePort = "443"
}


/**
 * 加载配置文件信息
 */
func loadConfig()  {
	cfg , _ :=goconfig.LoadConfigFile(configPath)
	configs,err := cfg.GetSection("instance")
	if err !=  nil {
		glog.Println(err)
	}else {


		if value, ok := configs["ipAddr"]; ok {
			//手动配置IP
			config.IpAddr = value
		}else {

			if value, ok := configs["port"]; ok {
				realValue ,err := strconv.Atoi(config.Port)
				if err != nil {
					glog.Println("value of 'port' must be a number and the value must between 1 and 65535 ")
				}else {
					if realValue >1 && realValue <65535 {
						config.Port = value
					}
				}
			}
			if value, ok := configs["securePort"]; ok {
				realValue ,err := strconv.Atoi(config.Port)
				if err != nil {
					glog.Println("value of 'securePort' must be a number and the value must between 1 and 65535 ")
				}else {
					if realValue >1 && realValue <65535 {
						config.Port = value
					}
				}
			}

			//如果没有手动配置IP ，但是配置了IP过滤器
			if value , ok := configs["ipPrefer"];ok {
				ipAddr ,_ := net.InterfaceAddrs()
				for _, address := range ipAddr{
					if ipnet , ok :=address.(*net.IPNet);ok && !ipnet.IP.IsLoopback() && common.IsIpv4(ipnet.IP.String()){
						if strings.HasPrefix(ipnet.IP.String(),value) {
							config.IpAddr = ipnet.IP.String()
							//替换下面的地址
							//健康检测的地址
							config.HealthCheckUrl = "http://"+config.IpAddr+":"+config.Port+"/health"
							//主页地址
							config.HomePageUrl ="http://"+config.IpAddr+":"+config.Port+"/"
							//状态信息地址
							config.StatusPageUrl ="http://"+config.IpAddr+":"+config.Port+"/info"
							break
						}
					}
				}
			}
		}

		config.HostName= config.IpAddr
		config.MetaDataInfo.Port = config.Port


		if value, ok := configs["app"]; ok {
			config.InstanceId =config.IpAddr+":"+ value +":"+config.Port
			config.App =  value // strings.ToUpper(value)
		}
		if value, ok := configs["hostName"]; ok {
			config.HostName = value
		}

		if value, ok := configs["instanceId"]; ok {
			config.InstanceId = strings.ToUpper(value)
		}
		if value, ok := configs["status"]; ok {
			realValue := strings.ToUpper(value)
			if realValue == "UP" || realValue == "DOWN"||realValue == "STARTING" ||
				realValue == "OUT_OF_SERVICE" || realValue == "UNKNOWN" {
				config.Status = realValue
			}else {
				glog.Println(" value of 'status' must be one in : UP 、DOWN 、STARTING 、OUT_OF_SERVICE 、 UNKNOWM ")
			}
		}
		if value, ok := configs["homePageUrl"]; ok {
			if !strings.HasPrefix(value , "http://") &&!strings.HasPrefix(value , "https://")  {
				value ="http://"+value
			}
			config.HomePageUrl = value
		}
		if value, ok := configs["statusPageUrl"]; ok {
			if !strings.HasPrefix(value , "http://") &&!strings.HasPrefix(value , "https://")  {
				value ="http://"+value
			}
			if !strings.HasSuffix(value,"/info") {
				value += "/value"
			}
			config.HomePageUrl = value
		}
		if value, ok := configs["healthCheckUrl"]; ok {
			if !strings.HasPrefix(value , "http://") &&!strings.HasPrefix(value , "https://")  {
				value ="http://"+value
			}
			if !strings.HasSuffix(value,"/health") {
				value += "/health"
			}
			config.HomePageUrl = value
		}
		if value, ok := configs["vipAddress"]; ok {
			config.VipAddress = value
		}else {
			config.VipAddress = config.App
		}
		if value, ok := configs["secureVipAddress"]; ok {
			config.SecureVipAddress = value
		}else {
			config.SecureVipAddress = config.App
		}
		if value, ok := configs["dataCenterInfoName"]; ok {
			config.DataCenterInfo.Name = value
		}
	}
}

