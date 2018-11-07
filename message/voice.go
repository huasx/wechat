package message

import "encoding/json"

//Voice 语音消息
type Voice struct {
	CommonToken

	Voice struct {
		MediaID CDATA `xml:"MediaId" json:"MediaId"`
	} `xml:"Voice"`
}

type MpVoice struct {
	MediaID string `xml:"MediaId" json:"MediaId"`
}

//NewVoice 回复语音消息
func NewVoice(MpVoice *MpVoice) *Voice {
	voice := new(Voice)
	voice.Voice.MediaID = CDATA{Value:MpVoice.MediaID}
	return voice
}

func NewVoiceByJson(js string) (*Voice, error) {
	voice := &MpVoice{}
	if err := json.Unmarshal([]byte(js), voice); err != nil {
		return nil, err
	}

	return NewVoice(voice), nil
}
