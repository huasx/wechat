package receive

//user_get_card  领取事件推送
/*
{"ToUserName":"xingbakezhongguo","FromUserName":"ouMrkq05TrpGHIyIMwyPhXfeq3Sg","CreateTime":"1525769734","MsgType":"event","Event":"user_get_card","CardId":"puMrkq-PQa6BDeqc7WhTTi6A4qPI","IsGiveByFriend":"","UserCardCode":"7310183920008006849","FriendUserName":"","OuterId":"","OldUserCardCode":"","IsRestoreMemberCard":"","IsRecommendByFriend":"","SourceScene":"SOURCE_SCENE_USER_ACCEPT_GIFTCARD","UnionId":"oGMNcv3WuDNBtcH1-6qOL08gPF-U"}

<xml>
   <ToUserName><![CDATA[gh_fc0a06a20993]]></ToUserName>
    <FromUserName><![CDATA[oZI8Fj040-be6rlDohc6gkoPOQTQ]]></FromUserName>
    <CreateTime>1472551036</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[user_get_card]]></Event>
    <CardId><![CDATA[pZI8Fjwsy5fVPRBeD78J4RmqVvBc]]></CardId>
    <IsGiveByFriend>0</IsGiveByFriend>
    <UserCardCode><![CDATA[226009850808]]></UserCardCode>
    <FriendUserName><![CDATA[]]></FriendUserName>
    <OuterId>0</OuterId>
    <OldUserCardCode><![CDATA[]]></OldUserCardCode>
    <OuterStr><![CDATA[12b]]></OuterStr>
    <IsRestoreMemberCard>0</IsRestoreMemberCard>
    <IsRecommendByFriend>0</IsRecommendByFriend>
</xml>
*/

type EventUserGetCard struct {
	ToUserName          string `json:"ToUserName"`
	FromUserName        string `json:"FromUserName"`
	CreateTime          int64  `json:"CreateTime"`
	MsgType             string `json:"MsgType"`
	Event               string `json:"Event"`
	CardID              string `json:"CardId"`
	IsGiveByFriend      string `json:"IsGiveByFriend"`
	UserCardCode        string `json:"UserCardCode"`
	FriendUserName      string `json:"FriendUserName"`
	OuterID             string `json:"OuterId"`
	OldUserCardCode     string `json:"OldUserCardCode"`
	IsRestoreMemberCard string `json:"IsRestoreMemberCard"`
	IsRecommendByFriend string `json:"IsRecommendByFriend"`
	SourceScene         string `json:"SourceScene"`
	UnionID             string `json:"UnionId"`
}
