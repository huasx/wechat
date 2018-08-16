package receive

/*
普通关注事件
<xml>
	<ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
	<FromUserName><![CDATA[oIawIjx961XLlZrGQStq3QR39-Ps]]></FromUserName>
	<CreateTime>1524729094</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[subscribe]]></Event>
	<EventKey><![CDATA[]]></EventKey>
</xml>

扫码关注事件
<xml>
  <ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
  <FromUserName><![CDATA[oIawIj6HjtMakoO42ghbHJeAIBEY]]></FromUserName>
  <CreateTime>1520908033</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[subscribe]]></Event>
  <EventKey><![CDATA[qrscene_yqym_11c4331372e5080690d12c4e4749630e]]></EventKey>
  <Ticket><![CDATA[gQE18jwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAyUm0yMncteEk4OGYxMDAwMDAwN2wAAgTT0qNaAwQAAAAA]]></Ticket>
</xml>
*/
type EventSubscribe struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	EventKey     string `json:"EventKey"`
	Ticket       string `json:"Ticket"`
}
