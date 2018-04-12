/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/12 14:38.
 */
package config

import (
	"github.com/Unknwon/goconfig"
	"go-eureka/glog"
	"strings"
	"go-eureka/model"
)

var eurekaConfig model.EurekaConfig

func loadEureka()  {
	cfg ,err :=goconfig.LoadConfigFile(configPath)
	if err != nil {
		glog.Println(err)
	}else {
		eurekas,err := cfg.GetSection("eureka")
		if err != nil {
			glog.Println(err)
		}else {
			if values ,ok :=eurekas["addresses"];ok {
				for _, value := range  strings.Split(values,",")  {
					if !strings.HasPrefix(value,"http://") && !strings.HasPrefix(value,"https://") {
						value = "http://"+value
					}
					if strings.HasSuffix(value,"/eureka") {
						eurekaConfig.Addresses = append(eurekaConfig.Addresses,value )
					}else {
						eurekaConfig.Addresses = append(eurekaConfig.Addresses,value+"/eureka" )
					}
				}
			}
		}
	}

}

func GetEurekaConfig() model.EurekaConfig {
	return eurekaConfig
}