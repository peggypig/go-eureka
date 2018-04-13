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
	"encoding/xml"
	"strings"
)

var RegisterList []string


func Register() {
	eurekaConfig := config.GetEurekaConfig()
	if len(eurekaConfig.Addresses)<=0 {
		glog.Println("no value of config :'eureka'->'addresses'")
	}else {
		config := config.GetInstanceConfig()
		bytes , err :=xml.MarshalIndent(config,"","")
		if err != nil {
			glog.Println(err)
		}else {
			for _,addr := range eurekaConfig.Addresses {
				resp , err :=http.Post(addr+"/apps/"+config.App," application/xml",strings.NewReader(string(bytes)))
				if err != nil{
					glog.Println(err)
				}else {
					if resp.StatusCode == 204 {
						glog.Println("register success，code 204",addr)
						RegisterList = append(RegisterList, addr)
					}else {
						glog.Println("Registration Center address may be incorrect：",addr)
					}
				}
			}
		}
	}
}
