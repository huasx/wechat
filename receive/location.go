package receive

//主动发送地理位置
/*
<xml>
	<ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
	<FromUserName><![CDATA[oIawIjx961XLlZrGQStq3QR39-Ps]]></FromUserName>
	<CreateTime>1524737684</CreateTime>
	<MsgType><![CDATA[location]]></MsgType>
	<Location_X>31.195066</Location_X>
	<Location_Y>121.437881</Location_Y>
	<Scale>16</Scale>
	<Label><![CDATA[上海市徐汇区虹桥路1号(华山路口)]]></Label>
	<MsgId>6548698488190351969</MsgId>
</xml>
*/
type Location struct {
	ToUserName   string  `json:"ToUserName"`
	FromUserName string  `json:"FromUserName"`
	CreateTime   int64   `json:"CreateTime"`
	MsgType      string  `json:"MsgType"`
	LocationX    float64 `json:"Location_X"` //！！！自动转换结构体 这里转换成了string类型 ！！！！
	LocationY    float64 `json:"Location_Y"` //！！！自动转换结构体 这里转换成了string类型 ！！！！
	Scale        float64 `json:"Scale"`      //！！！自动转换结构体 这里转换成了string类型 ！！！！
	Label        string  `json:"Label"`
	MsgID        int64   `json:"MsgId"` //！！！自动转换结构体 这里转换成了string类型 ！！！！
}
