package receive

/*
{"ToUserName":"gh_ba43c0866a80","FromUserName":"oIawIjx961XLlZrGQStq3QR39-Ps","CreateTime":"1529653497","MsgType":"image","PicUrl":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/eiccnC2JGxnj7cpjGlkFgOKndmO84JFSfk2TBln9EYmxadWJLszIFuYNh1uDrZRlCMrECib6X5ic3BMFxpD6v7r4A\/0","MsgId":"6569811744259309730","MediaId":"xzyTZY_nF4pX-w1EFMr3zHbdRS0YdWIDJYZcicAwbIsS5WaRnpOKcLlqy52CjHmY"}

<xml><ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
<FromUserName><![CDATA[oIawIjx961XLlZrGQStq3QR39-Ps]]></FromUserName>
<CreateTime>1529653497</CreateTime>
<MsgType><![CDATA[image]]></MsgType>
<PicUrl><![CDATA[http://mmbiz.qpic.cn/mmbiz_jpg/eiccnC2JGxnj7cpjGlkFgOKndmO84JFSfk2TBln9EYmxadWJLszIFuYNh1uDrZRlCMrECib6X5ic3BMFxpD6v7r4A/0]]></PicUrl>
<MsgId>6569811744259309730</MsgId>
<MediaId><![CDATA[xzyTZY_nF4pX-w1EFMr3zHbdRS0YdWIDJYZcicAwbIsS5WaRnpOKcLlqy52CjHmY]]></MediaId>
</xml>
*/

type Image struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	PicURL       string `json:"PicUrl"`
	MsgID        int64  `json:"MsgId"`
	MediaID      string `json:"MediaId"`
}
