package receive

/*
<xml>
	<ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
	<FromUserName><![CDATA[oIawIjx961XLlZrGQStq3QR39-Ps]]></FromUserName>
	<CreateTime>1524729330</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[VIEW]]></Event>
	<EventKey><![CDATA[http://www.verystar.cn]]></EventKey>
	<MenuId>431568521</MenuId>
</xml>
*/

//菜单跳转url
type EventView struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	EventKey     string `json:"EventKey,omitempty"`
	MenuID       string `json:"MenuId,omitempty"`
}
