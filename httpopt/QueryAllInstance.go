/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/12 15:24.
 */
package httpopt

import (
	"go-eureka/model"
	"go-eureka/config"
	"go-eureka/glog"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

func  QueryAllInstance() model.Applications{
	var apps model.Applications
	eurekaConfig := config.GetEurekaConfig()
	if len(eurekaConfig.Addresses)<=0 {
		glog.Println("no value of config :'eureka'->'addresses'")
	}else {
		for _,addr := range eurekaConfig.Addresses{
			resp , err := http.Get((addr+"/apps"))
			if  err != nil{
				glog.Println(resp)
			}else {
				defer resp.Body.Close()
				body ,err :=ioutil.ReadAll(resp.Body)
				if err != nil {
					glog.Println(err)
				}else {
					err :=xml.Unmarshal(body,&apps)
					if err != nil {
						glog.Println(err)
					}else {
						//如果没有出错 那么只需要拉去一个eureka的地址即可
						break
					}
				}
			}
		}
	}
	return apps
}
