package message

import "encoding/xml"

// MsgType 基本消息类型
type MsgType string

// EventType 事件类型
type EventType string

const (
	//MsgTypeText 表示文本消息
	MsgTypeText MsgType = "text"
	//MsgTypeImage 表示图片消息
	MsgTypeImage = "image"
	//MsgTypeVoice 表示语音消息
	MsgTypeVoice = "voice"
	//MsgTypeVideo 表示视频消息
	MsgTypeVideo = "video"
	//MsgTypeShortVideo 表示短视频消息[限接收]
	MsgTypeShortVideo = "shortvideo"
	//MsgTypeLocation 表示坐标消息[限接收]
	MsgTypeLocation = "location"
	//MsgTypeLink 表示链接消息[限接收]
	MsgTypeLink = "link"
	//MsgTypeMusic 表示音乐消息[限回复]
	MsgTypeMusic = "music"
	//MsgTypeNews 表示图文消息[限回复]
	MsgTypeNews = "news"
	//MsgTypeTransfer 表示消息消息转发到客服
	MsgTypeTransfer = "transfer_customer_service"
	//MsgTypeEvent 表示事件推送消息
	MsgTypeEvent = "event"
)

const (
	//EventSubscribe 订阅
	EventSubscribe EventType = "subscribe"
	//EventUnsubscribe 取消订阅
	EventUnsubscribe = "unsubscribe"
	//EventScan 用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	EventScan = "SCAN"
	//EventLocation 上报地理位置事件
	EventLocation = "LOCATION"
	//EventClick 点击菜单拉取消息时的事件推送
	EventClick = "CLICK"
	//EventView 点击菜单跳转链接时的事件推送
	EventView = "VIEW"
	//EventScancodePush 扫码推事件的事件推送
	EventScancodePush = "scancode_push"
	//EventScancodeWaitmsg 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventScancodeWaitmsg = "scancode_waitmsg"
	//EventPicSysphoto 弹出系统拍照发图的事件推送
	EventPicSysphoto = "pic_sysphoto"
	//EventPicPhotoOrAlbum 弹出拍照或者相册发图的事件推送
	EventPicPhotoOrAlbum = "pic_photo_or_album"
	//EventPicWeixin 弹出微信相册发图器的事件推送
	EventPicWeixin = "pic_weixin"
	//EventLocationSelect 弹出地理位置选择器的事件推送
	EventLocationSelect = "location_select"
	//EventTemplateSendJobFinish 发送模板消息推送通知
	EventTemplateSendJobFinish = "TEMPLATESENDJOBFINISH"
	//用户购买礼品卡付款成功
	EventGiftcardPayDone = "giftcard_pay_done"
	//用户购买后赠送
	EventGiftcardSendToFriend = "giftcard_send_to_friend"
	//用户领取礼品卡成功
	EventGiftcardSendUserAccept = "giftcard_user_accept"
	//领取
	EventUserGetCard = "user_get_card"
	//转赠
	EventUserGiftingCard = "user_gifting_card"
	//接收开票事件
	EventUserAuthorizeInvoice = "user_authorize_invoice"
	//审核事件推送
	EventCardPassCheck            = "card_pass_check"
	EventNotCardPassCheck         = "card_not_pass_check"
	EventAddStoreAuditInfo        = "add_store_audit_info"
	EventCreateMapPoiAuditInfo    = "create_map_poi_audit_info"
	EventKfCreateSession          = "kf_create_session"
	EventKfCloseSession           = "kf_close_session"
	EventSubmitMembercardUserInfo = "submit_membercard_user_info"
	EventUserConsumeCard          = "user_consume_card"
	EventUserDelCard              = "user_del_card"
	EventUserViewCard             = "user_view_card"
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

// CommonToken 消息中通用的结构
type CommonToken struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      MsgType  `xml:"MsgType"`
}

//SetToUserName set ToUserName
func (msg *CommonToken) SetToUserName(toUserName string) {
	msg.ToUserName = toUserName
}

//SetFromUserName set FromUserName
func (msg *CommonToken) SetFromUserName(fromUserName string) {
	msg.FromUserName = fromUserName
}

//SetCreateTime set createTime
func (msg *CommonToken) SetCreateTime(createTime int64) {
	msg.CreateTime = createTime
}

//SetMsgType set MsgType
func (msg *CommonToken) SetMsgType(msgType MsgType) {
	msg.MsgType = msgType
}
