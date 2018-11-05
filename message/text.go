package message

import "encoding/json"

//Text 文本消息
type Text struct {
	CommonToken
	Content string `xml:"Content" json:"Content"`
}

//NewText 初始化文本消息
func NewText(content string) *Text {
	text := new(Text)
	text.Content = content
	return text
}

//传入json格式 返回固定的微信格式
func NewTextByJson(jsContent string) (*Text, error) {
	text := &Text{}
	if err := json.Unmarshal([]byte(jsContent), text); err != nil {
		return nil, err
	}
	return NewText(text.Content), nil
}
