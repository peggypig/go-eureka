/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/12 14:13.
 */
package common

import "regexp"

const ipReg  = "(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\." +
			"(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\." +
			"(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\." +
			"(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)"


func IsIpv4(ip string) bool  {
	reg :=regexp.MustCompile(ipReg)
	return reg.MatchString(ip)
}
