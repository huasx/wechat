package receive

/*
giftcard_user_accept 用户领取礼品卡成功
<xml>
	<ToUserName><![CDATA[gh_3fcea188bf78]]></ToUserName>
	<FromUserName><![CDATA[obLatjgoYejavUtHsWwrX-2GtFJE]]></FromUserName>
	<CreateTime>1472631800</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[giftcard_user_accept]]></Event>
	<PageId><![CDATA[OQK0R3MaFnCm74Phw5hwFJlz5sn+jy1zzM2amDidDbU=]]></PageId>
	<OrderId><![CDATA[Z2y2rY74ksZX1ceuGA]]></OrderId>
	<IsChatRoom>true</IsChatRoom>
</xml>
*/

type EventGiftcardUserAccept struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	PageID       string `json:"PageId"`
	OrderID      string `json:"OrderId"`
	IsChatRoom   string `json:"IsChatRoom"`
}
