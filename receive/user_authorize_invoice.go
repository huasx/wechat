package receive

//user_authorize_invoice  收取授权完成事件推送

/*
{"ToUserName":"xingbakezhongguo","FromUserName":"ouMrkq5puZ-vsxQmEi88QbT8Ceeo","CreateTime":"1525769761","MsgType":"event","Event":"user_authorize_invoice","SuccOrderId":"18050610074402878502300701525769754","FailOrderId":"","AuthorizeAppId":"wxf83dd777cffa4eab","Source":"web","SPAppId":"d3hmODNkZDc3N2NmZmE0ZWFiXyDcTFRposi2bKRMqFrqphW1boY7gmkXPCscC6VB_t-_"}


<xml>
   <ToUserName><![CDATA[gh_fc0a06a20993]]></ToUserName>
   <FromUserName><![CDATA[oZI8Fj040-be6rlDohc6gkoPOQTQ]]></FromUserName>
   <CreateTime>1475134700</CreateTime>
   <MsgType><![CDATA[event]]></MsgType>
   <Event><![CDATA[user_authorize_invoice]]></Event>
   <SuccOrderId><![CDATA[1202933957956]]></SuccOrderId>
   <FailOrderId><![CDATA[]]></FailOrderId> < AuthorizeAppId ><![CDATA[]]></ AuthorizeAppId > <Source><![CDATA[]]></Source>
</xml>

*/

type EventUserAuthorizeInvoice struct {
	ToUserName     string `json:"ToUserName"`
	FromUserName   string `json:"FromUserName"`
	CreateTime     int64  `json:"CreateTime"`
	MsgType        string `json:"MsgType"`
	Event          string `json:"Event"`
	SuccOrderID    string `json:"SuccOrderId"`
	FailOrderID    string `json:"FailOrderId"`
	AuthorizeAppID string `json:"AuthorizeAppId"`
	Source         string `json:"Source"`
	SPAppID        string `json:"SPAppId"`
}
