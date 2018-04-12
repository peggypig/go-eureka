/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/12 14:36.
 */
package httpopt

import (
	"net/http"
	"go-eureka/config"
	"go-eureka/glog"
	"go-eureka/common"
)


func UnRegister()  {
	eurekaConfig := config.GetEurekaConfig()
	if len(eurekaConfig.Addresses)<=0 {
		glog.Println("no value of config :'eureka'->'addresses'")
	}else {
		config := config.GetInstanceConfig()
		client := &http.Client{}
		for _,addr := range eurekaConfig.Addresses {
			req , err :=http.NewRequest(http.MethodDelete,addr+"/apps/"+config.App+"/"+config.InstanceId,nil)
			if err != nil {
				glog.Println(err)
			}else {
				resp , err :=client.Do(req)
				if err != nil {
					glog.Println(err)
				}else {
					if resp.StatusCode == 200 {
						glog.Println("unRegister success，code 200")
						common.Remove(RegisterList,addr)
					}
				}
			}
		}
	}
}


