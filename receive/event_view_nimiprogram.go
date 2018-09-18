package receive

/*
<xml><ToUserName><![CDATA[gh_74617c2e85e1]]></ToUserName>
<FromUserName><![CDATA[oi4ihw8OOJhvCUT_57u72TGroG5o]]></FromUserName>
<CreateTime>1537252084</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[view_miniprogram]]></Event>
<EventKey><![CDATA[pages/meeting/meeting]]></EventKey>
<MenuId>433418371</MenuId>
</xml>
*/

//菜单跳转小程序
type EventViewMiniprogram struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	EventKey     string `json:"EventKey,omitempty"`
	MenuID       string `json:"MenuId,omitempty"`
}
