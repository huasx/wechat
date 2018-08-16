package receive

/*
card_pass_check 审核事件推送
{"ToUserName":"gh_6e9a448fb4cc","FromUserName":"o7_2gwJ20snQQ64HSgiYTWmm0BJU","CreateTime":"1525914463","MsgType":"event","Event":"card_pass_check","CardId":"p7_2gwC9bC2UnQWf-vmbr5y8Cb9o","RefuseReason":""}

<xml>
   <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[card_pass_check]]></Event> //不通过为card_not_pass_check
   <CardId><![CDATA[cardid]]></CardId>
    <RefuseReason><![CDATA[非法代制]]></RefuseReason>
</xml>
*/

type EventCardPassCheck struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	CardID       string `json:"CardId"`
	RefuseReason string `json:"RefuseReason"`
}
