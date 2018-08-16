package receive

/*
{"ToUserName":"gh_ba43c0866a80","FromUserName":"oIawIjx961XLlZrGQStq3QR39-Ps","CreateTime":"1529660135","MsgType":"link","Title":"\u6545\u969c\u5904\u7406\u6700\u4f73\u5b9e\u8df5\uff1a\u6545\u969c\u6539\u8fdb|\u9080\u8bf7\u4f60\u8bfb","Description":{},"Url":"https:\/\/time.geekbang.org\/column\/article\/aa70d7c538b621c08eb6f8706f5d182a\/share?from=timeline","MsgId":"6569840254252221623"}

<xml><ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
<FromUserName><![CDATA[oIawIjx961XLlZrGQStq3QR39-Ps]]></FromUserName>
<CreateTime>1529660135</CreateTime>
<MsgType><![CDATA[link]]></MsgType>
<Title><![CDATA[故障处理最佳实践：故障改进 | 邀请你读]]></Title>
<Description><![CDATA[]]></Description>
<Url><![CDATA[https://time.geekbang.org/column/article/aa70d7c538b621c08eb6f8706f5d182a/share?from=timeline]]></Url>
<MsgId>6569840254252221623</MsgId>
</xml>
*/

type Link struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Title        string `json:"Title"`
	Description  string `json:"Description"`
	URL          string `json:"Url"`
	MsgID        int64  `json:"MsgId"`
}
