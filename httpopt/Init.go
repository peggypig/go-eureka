/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/12 16:52.
 */
package httpopt

import (
	"time"
)

func init()  {
	//自动注册
	Register()
	go autoSendHeartbeat()
}

func autoSendHeartbeat() {
	defer UnRegister()
	for {
		time.Sleep(25*time.Second)
		if  len(RegisterList) >0 {
			SendHeartbeat()
		}else {
			break
		}
	}
}
