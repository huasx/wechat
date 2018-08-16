package receive

//user_consume_card   核销事件推送

/*
{"ToUserName":"xingbakezhongguo","FromUserName":"ouMrkq5_fZiVurrP-JD_-PU5JaTY","CreateTime":"1525769736","MsgType":"event","Event":"user_consume_card","CardId":"puMrkq5TXp1Ys5WqDu8MFh36Dj_c","UserCardCode":"7310183870031457178","ConsumeSource":"FROM_API","LocationName":"","StaffOpenId":"ouMrkq7iT8BTqPIa92fRAs_JWvL0","VerifyCode":"","RemarkAmount":"","OuterStr":"","LocationId":""}


<xml>
    <ToUserName><![CDATA[gh_fc0a06a20993]]></ToUserName>
    <FromUserName><![CDATA[oZI8Fj040-be6rlDohc6gkoPOQTQ]]></FromUserName>
    <CreateTime>1472549042</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[user_consume_card]]></Event>
    <CardId><![CDATA[pZI8Fj8y-E8hpvho2d1ZvpGwQBvA]]></CardId>
    <UserCardCode><![CDATA[452998530302]]></UserCardCode>
    <ConsumeSource><![CDATA[FROM_API]]></ConsumeSource>
    <LocationName><![CDATA[]]> </LocationName>
    <StaffOpenId><![CDATA[oZ********nJ3bPJu_Rtjkw4c]]></StaffOpenId>
    <VerifyCode><![CDATA[]]></VerifyCode>
    <RemarkAmount><![CDATA[]]></RemarkAmount>
    <OuterStr><![CDATA[xxxxx]]></OuterStr>
</xml>

*/

type EventUserConsumeCard struct {
	ToUserName    string `json:"ToUserName"`
	FromUserName  string `json:"FromUserName"`
	CreateTime    int64  `json:"CreateTime"`
	MsgType       string `json:"MsgType"`
	Event         string `json:"Event"`
	CardID        string `json:"CardId"`
	UserCardCode  string `json:"UserCardCode"`
	ConsumeSource string `json:"ConsumeSource"`
	LocationName  string `json:"LocationName"`
	StaffOpenID   string `json:"StaffOpenId"`
	VerifyCode    string `json:"VerifyCode"`
	RemarkAmount  string `json:"RemarkAmount"`
	OuterStr      string `json:"OuterStr"`
	LocationID    string `json:"LocationId"`
}
