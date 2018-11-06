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
func NewText(content CDATA) *Text {
	text := new(Text)
	text.Content = content
	return text
}

//传入json格式 返回固定的微信格式
func NewTextByJson(jsContent string) (*Text, error) {
	text := &MpText{}
	if err := json.Unmarshal([]byte(jsContent), text); err != nil {
		return nil, err
	}
	return NewText(CDATA{Value: text.Content}), nil
}
