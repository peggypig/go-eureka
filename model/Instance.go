/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/12 10:29.
 */
package model

type Instance struct {
	InstanceId                    string `xml:"instanceId"`
	HostName                      string `xml:"hostName"`
	App                           string `xml:"app"`
	IpAddr                        string `xml:"ipAddr"`
	Status                        string `xml:"status"`
	Overriddenstatus              string `xml:"overriddenstatus,omitempty"`
	Port                          string `xml:"port"`
	SecurePort                    string `xml:"securePort"`
	CountryId                     string `xml:"countryId,omitempty"`
	HomePageUrl                   string `xml:"homePageUrl"`
	StatusPageUrl                 string `xml:"statusPageUrl"`
	HealthCheckUrl                string `xml:"healthCheckUrl"`
	VipAddress                    string `xml:"vipAddress"`
	SecureVipAddress              string `xml:"secureVipAddress"`
	IsCoordinatingDiscoveryServer string `xml:"isCoordinatingDiscoveryServer,omitempty"`
	LastUpdatedTimestamp          string `xml:"lastUpdatedTimestamp,omitempty"`
	LastDirtyTimestamp            string `xml:"lastDirtyTimestamp,omitempty"`
	ActionType                    string `xml:"actionType,omitempty"`
}
