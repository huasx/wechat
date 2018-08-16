package receive

//add_store_audit_info 创建门店的审核结果
/*
{"ToUserName":"youyiku811016","FromUserName":"oNJj3nwHtGXLkR2m1RSHam9rcFhE","CreateTime":"1527564786","MsgType":"event","Event":"add_store_audit_info","audit_id":"418233436","status":"1","reason":"","is_upgrade":"","poiid":"487956548"}

<xml><ToUserName>< ![CDATA[gh_4346ac1514d8] ]></ToUserName>
<FromUserName>< ![CDATA[od1P50M-fNQI5Gcq-trm4a7apsU8] ]></FromUserName>
<CreateTime>1488856741</CreateTime>
<MsgType>< ![CDATA[event] ]></MsgType>
<Event>< ![CDATA[add_store_audit_info] ]></Event>
<audit_id>11111</audit_id>
<status>3</status>
<reason>< ![CDATA[xxx] ]></reason>
<is_upgrade>3</is_upgrade>
<poiid>12344</poiid>
</xml>

*/
//创建门店的审核结果
type EventAddStoreAuditInfo struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	AuditID      string `json:"audit_id"`
	Status       string `json:"status"`
	Reason       string `json:"reason"`
	IsUpgrade    string `json:"is_upgrade"`
	Poiid        string `json:"poiid"`
}
