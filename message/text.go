package message

import "encoding/json"

//Text 文本消息
type Text struct {
	CommonToken
	Content CDATA `xml:"Content" json:"Content"`
}

type MpText struct {
	Content string `json:"Content"`
}

//NewText 初始化文本消息
func NewText(MpText *MpText) *Text {
	text := new(Text)
	text.Content = CDATA{Value:MpText.Content}
	return text
}

//传入json格式 返回固定的微信格式
func NewTextByJson(jsContent string) (*Text, error) {
	MpText := &MpText{}
	if err := json.Unmarshal([]byte(jsContent), MpText); err != nil {
		return nil, err
	}
	return NewText(MpText), nil
}
