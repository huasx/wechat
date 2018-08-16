package receive

/*
templatesendjobfinish
<xml>
	<ToUserName><![CDATA[gh_7bda750f4303]]></ToUserName>
	<FromUserName><![CDATA[ocv_Uw4wKEHQianelWoc1aRkgmgg]]></FromUserName>
	<CreateTime>1524676046</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[TEMPLATESENDJOBFINISH]]></Event>
	<MsgID>253063755649368064</MsgID>
	<Status><![CDATA[success]]></Status>
</xml>
*/

type EventTemplateSendJobFinish struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"` //PHP版本是小写
	MsgID        int64  `json:"MsgID"`
	Status       string `json:"Status"`
}
