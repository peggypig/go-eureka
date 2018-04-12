/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/12 15:17.
 */
package httpopt

import (
	"go-eureka/glog"
	"net/http"
	"go-eureka/config"
)

func SendHeartbeat()  {
	eurekaConfig := config.GetEurekaConfig()
	if len(eurekaConfig.Addresses)<=0 {
		glog.Println("no value of config :'eureka'->'addresses'")
	}else {
		client := &http.Client{}
		config := config.GetInstanceConfig()
		for _,addr := range eurekaConfig.Addresses {
			url := addr+"/apps/"+config.App+"/"+config.InstanceId
			req , err :=http.NewRequest(http.MethodPut,url,nil)
			if err != nil {
				glog.Println(err)
			}else {
				resp , err := client.Do(req)
				if err != nil {
					glog.Println(err)
				}else {
					if resp.StatusCode == 200 {
						glog.Println("send heartbeat success，code 200")
					}else if resp.StatusCode == 404 {
						glog.Println("send heartbeat success，but this instance is not exits,code 404")
					}else {
						glog.Println("the status of send heartbeat is unknown ，code ",resp.StatusCode)
					}
				}
			}
		}
	}
}
