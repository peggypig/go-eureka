/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/12 12:09.
 */
package model

import "encoding/xml"

type InstanceConfig struct {
	Instance
	XMLName 		xml.Name `xml:"instance"` // 指定最外层的标签为instance
	DataCenterInfo DataCenterInfo `xml:"dataCenterInfo"`
}
