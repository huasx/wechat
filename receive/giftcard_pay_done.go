package receive

/*
giftcard_pay_done 用户购买礼品卡付款成功
<xml>
	<ToUserName><![CDATA[gh_3fcea188bf78]]></ToUserName>
	<FromUserName><![CDATA[obLatjgoYejavUtHsWwrX-2GtFJE]]></FromUserName>
	<CreateTime>1472631550</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[giftcard_pay_done]]></Event>
	<PageId><![CDATA[OQK0R3MaFnCm74Phw5hwFJlz5sn+jy1zzM2amDidDbU=]]></PageId>
	<OrderId><![CDATA[Z2y2rY74ksZX1ceuGA]]></OrderId>
</xml>
*/

type EventGiftcardPayDone struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	PageID       string `json:"PageId"`
	OrderID      string `json:"OrderId"`
}
