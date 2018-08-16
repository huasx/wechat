package receive

//自动上报地理位置
/*
<xml>
  <ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
  <FromUserName><![CDATA[oIawIj6HjtMakoO42ghbHJeAIBEY]]></FromUserName>
  <CreateTime>1520908043</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[LOCATION]]></Event>
  <Latitude>31.197428</Latitude>
  <Longitude>121.433556</Longitude>
  <Precision>30.000000</Precision>
</xml>

<xml>
	<ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
	<FromUserName><![CDATA[oIawIjx961XLlZrGQStq3QR39-Ps]]></FromUserName>
	<CreateTime>1524737876</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[LOCATION]]></Event>
	<Latitude>31.196934</Latitude>
	<Longitude>121.433250</Longitude>
	<Precision>90.000000</Precision>
</xml>
*/

type EventLocation struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	Latitude     string `json:"Latitude"`
	Longitude    string `json:"Longitude"`
	Precision    string `json:"Precision"`
}
