/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/12 11:40.
 */
package glog

import "log"

const TAG  = "go-eureka :"

func Println(v ...interface{}){
	log.Println(TAG, v)
}

