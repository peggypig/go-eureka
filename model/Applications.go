/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/12 10:36.
 */
package model

type Applications struct {
	VersionsDelta string        `xml:"versions_delta"`
	AppsHashcode  string        `xml:"apps_hashcode"`
	Applications   []Application `xml:"application"`
}
