package message

import "encoding/xml"

// MsgType 基本消息类型
type MsgType string

// EventType 事件类型
type EventType string

const (
	MsgTypeText       = "text"                      //MsgTypeText 表示文本消息
	MsgTypeImage      = "image"                     //MsgTypeImage 表示图片消息
	MsgTypeVoice      = "voice"                     //MsgTypeVoice 表示语音消息
	MsgTypeVideo      = "video"                     //MsgTypeVideo 表示视频消息
	MsgTypeShortVideo = "shortvideo"                //MsgTypeShortVideo 表示短视频消息[限接收]
	MsgTypeLocation   = "location"                  //MsgTypeLocation 表示坐标消息[限接收]
	MsgTypeLink       = "link"                      //MsgTypeLink 表示链接消息[限接收]
	MsgTypeMusic      = "music"                     //MsgTypeMusic 表示音乐消息[限回复]
	MsgTypeNews       = "news"                      //MsgTypeNews 表示图文消息[限回复]
	MsgTypeTransfer   = "transfer_customer_service" //MsgTypeTransfer 表示消息消息转发到客服
	MsgTypeEvent      = "event"                     //MsgTypeEvent 表示事件推送消息
)

const (
	EventSubscribe                EventType = "subscribe"               //EventSubscribe 订阅
	EventUnsubscribe              EventType = "unsubscribe"             //EventUnsubscribe 取消订阅
	EventScan                     EventType = "SCAN"                    //EventScan 用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	EventLocation                 EventType = "LOCATION"                //EventLocation 上报地理位置事件
	EventClick                    EventType = "CLICK"                   //EventClick 点击菜单拉取消息时的事件推送
	EventView                     EventType = "VIEW"                    //EventView 点击菜单跳转链接时的事件推送
	EventScancodePush             EventType = "scancode_push"           //EventScancodePush 扫码推事件的事件推送
	EventScancodeWaitmsg          EventType = "scancode_waitmsg"        //EventScancodeWaitmsg 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventPicSysphoto              EventType = "pic_sysphoto"            //EventPicSysphoto 弹出系统拍照发图的事件推送
	EventPicPhotoOrAlbum          EventType = "pic_photo_or_album"      //EventPicPhotoOrAlbum 弹出拍照或者相册发图的事件推送
	EventPicWeixin                EventType = "pic_weixin"              //EventPicWeixin 弹出微信相册发图器的事件推送
	EventLocationSelect           EventType = "location_select"         //EventLocationSelect 弹出地理位置选择器的事件推送
	EventTemplateSendJobFinish    EventType = "TEMPLATESENDJOBFINISH"   //EventTemplateSendJobFinish 发送模板消息推送通知
	EventGiftcardPayDone          EventType = "giftcard_pay_done"       //用户购买礼品卡付款成功
	EventGiftcardSendToFriend     EventType = "giftcard_send_to_friend" //用户购买后赠送
	EventGiftcardSendUserAccept   EventType = "giftcard_user_accept"    //用户领取礼品卡成功
	EventUserGetCard              EventType = "user_get_card"           //领取
	EventUserGiftingCard          EventType = "user_gifting_card"       //转赠
	EventUserAuthorizeInvoice     EventType = "user_authorize_invoice"  //接收开票事件
	EventCardPassCheck            EventType = "card_pass_check"         //审核事件推送
	EventNotCardPassCheck         EventType = "card_not_pass_check"
	EventAddStoreAuditInfo        EventType = "add_store_audit_info"
	EventCreateMapPoiAuditInfo    EventType = "create_map_poi_audit_info"
	EventKfCreateSession          EventType = "kf_create_session"
	EventKfCloseSession           EventType = "kf_close_session"
	EventSubmitMembercardUserInfo EventType = "submit_membercard_user_info"
	EventUserConsumeCard          EventType = "user_consume_card"
	EventUserDelCard              EventType = "user_del_card"
	EventUserViewCard             EventType = "user_view_card"
	EventViewMiniprogram          EventType = "view_miniprogram"
)

type BaseMessage struct {
	ToUserName   string  `xml:"ToUserName" json:"ToUserName"`
	FromUserName string  `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64   `xml:"CreateTime" json:"CreateTime"`
	MsgType      MsgType `xml:"MsgType" json:"MsgType"`
	Event        string  `xml:"Event" json:"Event"`
}

func (m *BaseMessage) GetId() string {
	if m.Event == "" {
		return string(m.MsgType)
	}
	return string(m.MsgType) + ":" + m.Event
}

//EventPic 发图事件推送
type EventPic struct {
	PicMd5Sum string `xml:"PicMd5Sum"`
}

//EncryptedXMLMsg 安全模式下的消息体
type EncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"ToUserName"`
	EncryptedMsg string   `xml:"Encrypt"    json:"Encrypt"`
}

//ResponseEncryptedXMLMsg 需要返回的消息体
type ResponseEncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	EncryptedMsg string   `xml:"Encrypt"      json:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature" json:"MsgSignature"`
	Timestamp    int64    `xml:"TimeStamp"    json:"TimeStamp"`
	Nonce        string   `xml:"Nonce"        json:"Nonce"`
}

type CDATA struct {
	Value string `xml:",cdata"`
}

// CommonToken 消息中通用的结构
type CommonToken struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
}

//SetToUserName set ToUserName
func (msg *CommonToken) SetToUserName(toUserName string) {
	msg.ToUserName = CDATA{
		Value: toUserName,
	}
}

//SetFromUserName set FromUserName
func (msg *CommonToken) SetFromUserName(fromUserName string) {
	msg.FromUserName = CDATA{
		Value: fromUserName,
	}
}

//SetCreateTime set createTime
func (msg *CommonToken) SetCreateTime(createTime int64) {
	msg.CreateTime = createTime
}

//SetMsgType set MsgType
func (msg *CommonToken) SetMsgType(msgType string) {
	msg.MsgType = CDATA{
		Value: msgType,
	}
}
