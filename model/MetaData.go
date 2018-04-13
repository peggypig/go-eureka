/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/13 16:44.
 */
package model

import "encoding/xml"

type MetaData struct {
	XMLName xml.Name `xml:"metadata"`
	Port    string   `xml:"management.port"`
}
