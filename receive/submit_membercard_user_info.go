package receive

//submit_membercard_user_info   会员卡激活事件推送

/*
{"ToUserName":"gh_6e9a448fb4cc","FromUserName":"o7_2gwAm3K1kDonw3zgYIHLZk3X8","CreateTime":"1525769761","MsgType":"event","Event":"submit_membercard_user_info","CardId":"p7_2gwEdnMx3Ma3i6Xi3hH-s7oCE","UserCardCode":"869643314887"}

<xml>
   <ToUserName> <![CDATA[gh_3fcea188bf78]]></ToUserName>
    <FromUserName><![CDATA[obLatjlaNQKb8FqOvt1M1x1lIBFE]]></FromUserName>
    <CreateTime>1432668700</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[submit_membercard_user_info]]></Event>
    <CardId><![CDATA[pbLatjtZ7v1BG_ZnTjbW85GYc_E8]]></CardId>
    <UserCardCode><![CDATA[018255396048]]></UserCardCode>
</xml>
*/
type EventSubmitMembercardUserInfo struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	CardID       string `json:"CardId"`
	UserCardCode string `json:"UserCardCode"`
}
