package receive

/*
create_map_poi_audit_info 腾讯地图中创建门店的审核结果
{"ToUserName":"youyiku811016","FromUserName":"oNJj3nwHtGXLkR2m1RSHam9rcFhE","CreateTime":"1527562153","MsgType":"event","Event":"create_map_poi_audit_info","audit_id":"3153770","map_poi_id":"8674870308828146722","name":"优衣库(文峰城市广场店)","address":"南通市崇川区虹桥路1号7幢文峰城市广场1F","latitude":"31.997041675357","longitude":"120.884002815183","status":"","sh_remark":""}


<xml><ToUserName>< ![CDATA[gh_4346ac1514d8] ]></ToUserName>
<FromUserName>< ![CDATA[od1P50M-fNQI5Gcq-trm4a7apsU8] ]></FromUserName>
<CreateTime>1488856741</CreateTime>
<MsgType>< ![CDATA[event] ]></MsgType>
<Event>< ![CDATA[create_map_poi_audit_info] ]></Event>
<audit_id>11111</audit_id>
<status>1</status>
<map_poi_id>< ![CDATA[xxx] ]></map_poi_id>
<name>< ![CDATA[xxx] ]></name>
<address>< ![CDATA[xxx] ]></address>
<latitude>< ![CDATA[xxx] ]></latitude>
<longitude>< ![CDATA[xxx] ]></longitude>
<sh_remark>< ![CDATA[xxx] ]></sh_remark>
</xml>
*/

type EventCreateMapPoiAuditInfo struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	AuditID      string `json:"audit_id"`
	MapPoiID     string `json:"map_poi_id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Status       string `json:"status"`
	ShRemark     string `json:"sh_remark"`
}
