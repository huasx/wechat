package receive

//user_del_card  删除事件推送

/*
{"ToUserName":"xingbakezhongguo","FromUserName":"ouMrkq_bqm9zbJLbxSBjv1Lguujo","CreateTime":"1525769734","MsgType":"event","Event":"user_del_card","CardId":"puMrkq7z3ckqNjraG1I7DIPbKBvo","UserCardCode":"7310183660024958926"}

<xml>
 <ToUserName><![CDATA[toUser]]></ToUserName>
 <FromUserName><![CDATA[FromUser]]></FromUserName>
 <CreateTime>123456789</CreateTime>
 <MsgType><![CDATA[event]]></MsgType>
 <Event><![CDATA[user_del_card]]></Event>
 <CardId><![CDATA[cardid]]></CardId>
 <UserCardCode><![CDATA[12312312]]></UserCardCode>
</xml>
*/
type EventUserDelCard struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	CardID       string `json:"CardId"`
	UserCardCode string `json:"UserCardCode"`
}
