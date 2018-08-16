package receive

//已关注的用户扫码
/*
<xml><ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
<FromUserName><![CDATA[oIawIjx961XLlZrGQStq3QR39-Ps]]></FromUserName>
<CreateTime>1527658705</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[SCAN]]></Event>
<EventKey><![CDATA[20222]]></EventKey>
<Ticket><![CDATA[gQHk7zoAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL05IUVA1dVRtbFktc3FTcFRpRmpsAAIEbMIQVwMEAAAAAA==]]></Ticket>
</xml>
*/

type EventScan struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	EventKey     string `json:"EventKey"`
	Ticket       string `json:"Ticket"`
}
