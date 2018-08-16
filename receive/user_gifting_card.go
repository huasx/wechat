package receive

//user_gifting_card 转赠事件推送

/*
{"ToUserName":"xingbakezhongguo","FromUserName":"ouMrkq8D_GXgP1lZui6HFzBNnLZQ","CreateTime":"1525769762","MsgType":"event","Event":"user_gifting_card","CardId":"puMrkq4zepypLPFYojk29hlwroyU","UserCardCode":"064687852577","IsReturnBack":"1","FriendUserName":"ouMrkqyw1tL4xdy_MS9X4-Uh43qQ","IsChatRoom":""}

<xml>
  <ToUserName><![CDATA[gh_3fcea188bf78]]></ToUserName>
  <FromUserName><![CDATA[obLatjjwDolFjRRd3doGIdwNqRXw]]></FromUserName>
  <CreateTime>1474181868</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[user_gifting_card]]></Event>
  <CardId><![CDATA[pbLatjhU-3pik3d4PsbVzvBxZvJc]]></CardId>
  <UserCardCode><![CDATA[297466945104]]></UserCardCode>
  <IsReturnBack>0</IsReturnBack>
  <FriendUserName><![CDATA[obLatjlNerkb62HtSdQUx66C4NTU]]></FriendUserName>
  <IsChatRoom>0</IsChatRoom>
</xml>

*/

type EventUserGiftingCard struct {
	ToUserName     string `json:"ToUserName"`
	FromUserName   string `json:"FromUserName"`
	CreateTime     int64  `json:"CreateTime"`
	MsgType        string `json:"MsgType"`
	Event          string `json:"Event"`
	CardID         string `json:"CardId"`
	UserCardCode   string `json:"UserCardCode"`
	IsReturnBack   string `json:"IsReturnBack"`
	FriendUserName string `json:"FriendUserName"`
	IsChatRoom     string `json:"IsChatRoom"`
}
