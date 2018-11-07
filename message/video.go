package message

import "encoding/json"

//Video 视频消息
type Video struct {
	CommonToken

	Video struct {
		MediaID     CDATA `xml:"MediaId"`
		Title       CDATA `xml:"Title,omitempty"`
		Description CDATA `xml:"Description,omitempty"`
	} `xml:"Video"`
}

type MpVideo struct {
	MediaID     string `xml:"MediaId" json:"MediaId"`
	Title       string `xml:"Title,omitempty" json:"Title,omitempty"`
	Description string `xml:"Description,omitempty" json:"Description,omitempty"`
}

//NewVideo 回复图片消息
func NewVideo(MpVideo *MpVideo) *Video {
	video := new(Video)
	video.Video.MediaID     = CDATA{Value:MpVideo.MediaID}
	video.Video.Title       = CDATA{Value:MpVideo.Title}
	video.Video.Description = CDATA{Value:MpVideo.Description}
	return video
}

//数据库目前没有保存Title和 Description
func NewVideoByJson(js string) (*Video, error) {
	video := &MpVideo{}
	if err := json.Unmarshal([]byte(js), video); err != nil {
		return nil, err
	}
	return NewVideo(video), nil
}
