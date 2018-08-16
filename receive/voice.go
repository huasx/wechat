package receive

//语音消息
/**
<xml><ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
<FromUserName><![CDATA[oIawIjx961XLlZrGQStq3QR39-Ps]]></FromUserName>
<CreateTime>1533798738</CreateTime>
<MsgType><![CDATA[voice]]></MsgType>
<MediaId><![CDATA[vuUTV-gngn2pLsgjoxzllVpRbUasJQUzkFjKE_unQ7sJ1MuX_J1pRTNhB7UGTkJu]]></MediaId>
<Format><![CDATA[amr]]></Format>
<MsgId>6587615418356072448</MsgId>
<Recognition><![CDATA[12345。]]></Recognition>
</xml>
*/

type Voice struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	MediaID      string `json:"MediaId"`
	Format       string `json:"Format"`
	MsgID        int64  `json:"MsgId"`
	Recognition  string `json:"Recognition"`
}
