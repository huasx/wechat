package receive

/*
//user_view_card 进入会员卡事件推送

{"ToUserName":"gh_cc07165ce6a9","FromUserName":"o4zAjtzMRFQix9nqOA5YhXakAkJg","CreateTime":"1525769628","MsgType":"event","Event":"user_view_card","CardId":"p4zAjt5EMiNeQI9qDHjGEA7sldwY","UserCardCode":"056309926961","OuterStr":"","Encrypt":"WuCIxx/C0El78odY7vFXPCVxRL71hPiKw+SPm2wfN8iAV3FCVIteozNS0B0PWqE/4jmlSHZVoTsRPN4b1oTrMeCn9wj2bR6/jE9jyNeipeeKNCtFn9URKV+itZ+Hy5RrMV9mZUMFYLbnnMGR4lJXE9fyg72e3AuJbEX/n/OtRRXATnGTxX+EBOTkSMesPMCqzylXpjYifZ9g1DyMnIYA8pFTwvFceta4h16lZoCf39q2aXOOBvL/8hggyOwLuafCBuWWcZhsewskFMNy1ASdJ9Tz+7uE1GErYt87XTSdIkU/PR/z8bmGrMiU2yCd02tc1f5OZxeWjzAV1LombI/DJqgaBk3KtGnGIp1S4N3Y4dcZsC1vgRUm9xbGWek9kdXs+Zp3zlK+LbJSbRsTzdyYYrTryKYAJ1DMpKKwsFkfbe0sx4pDTSyHTCjXGuNqdpSxpPOFwmLDhKLjT49ho8omr0cAfWAezH+kG3PP7H1z4Q09YZMACD7nCeY+fXqjCDMP2gQiPm4LhShwM0kgyJbC5Fuug99OH6Q7q8O3ZZlflgvloBrwWiaCuYf2Fx7RqwoqnFSzElcc4ll5RnhbeEGiOA=="}

{"ToUserName":"xingbakezhongguo","FromUserName":"ouMrkq_iR5J8l7t-vMISpvEQcJ9g","CreateTime":"1525769628","MsgType":"event","Event":"user_view_card","CardId":"puMrkq-uCBrVuEimgcg-Ne4nrbAc","UserCardCode":"7310292440006153094","OuterStr":""}

<xml>
    <ToUserName><![CDATA[gh_fcxxxx6a20993]]></ToUserName>
    <FromUserName><![CDATA[oZI8Fj040-xxxxx6gkoPOQTQ]]></FromUserName>
    <CreateTime>1467811138</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[user_view_card]]></Event>
    <CardId><![CDATA[pZI8Fj2ezBbxxxxxT2UbiiWLb7Bg]]></CardId>
    <UserCardCode><![CDATA[4xxxxxxxx8558]]></UserCardCode>
    <OuterStr><![CDATA[12b]]></OuterStr>
</xml>

*/

type EventUserViewCard struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	CardID       string `json:"CardId"`
	UserCardCode string `json:"UserCardCode"`
	OuterStr     string `json:"OuterStr"`
}
