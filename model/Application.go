/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/12 10:34.
 */
package model

type Application struct {
	Name           string         `xml:"name"`
	InstanceInfo   Instance       `xml:"instance"`
	DataCenterInfo DataCenterInfo `xml:"dataCenterInfo"`
}
