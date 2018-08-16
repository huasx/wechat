package receive

//视频消息
/*
<xml><ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
<FromUserName><![CDATA[oIawIjx961XLlZrGQStq3QR39-Ps]]></FromUserName>
<CreateTime>1533799071</CreateTime>
<MsgType><![CDATA[video]]></MsgType>
<MediaId><![CDATA[HBDT64cX2n4FptFClCXK6E8JkY3kYyStKBn7Duxuim4AvmCXahObd9tItEXC2i2b]]></MediaId>
<ThumbMediaId><![CDATA[hMGoVdWIRCDVY8uIeN696zRYdIfAWVtsXhJTY2pBSw4F7yxglnzLd-rWZIGFflu_]]></ThumbMediaId>
<MsgId>6587616849013043953</MsgId>
</xml>
*/

type Video struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	MediaID      string `json:"MediaId"`
	ThumbMediaID string `json:"ThumbMediaId"`
	MsgID        int64  `json:"MsgId"`
}
