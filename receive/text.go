package receive

/*
<xml>
  <ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
  <FromUserName><![CDATA[oIawIj6HjtMakoO42ghbHJeAIBEY]]></FromUserName>
  <CreateTime>1520908477</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[123]]></Content>
  <MsgId>6532252169355276694</MsgId>
</xml>
*/

type Text struct {
	ToUserName   string `xml:"ToUserName" json:"ToUserName"`
	FromUserName string `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime" json:"CreateTime"`
	MsgType      string `xml:"MsgType" json:"MsgType"`
	Content      string `xml:"Content" json:"Content"`
	MsgID        int64  `xml:"MsgId" json:"MsgId"`
}
