package message

import "encoding/json"

//Video 视频消息
type Video struct {
	CommonToken

	Video struct {
		MediaID     string `xml:"MediaId"`
		Title       string `xml:"Title,omitempty"`
		Description string `xml:"Description,omitempty"`
	} `xml:"Video"`
}

type MpVideo struct {
	MediaID     string `xml:"MediaId" json:"MediaId"`
	Title       string `xml:"Title,omitempty" json:"Title,omitempty"`
	Description string `xml:"Description,omitempty" json:"Description,omitempty"`
}

//NewVideo 回复图片消息
func NewVideo(mediaID, title, description string) *Video {
	video := new(Video)
	video.Video.MediaID = mediaID
	video.Video.Title = title
	video.Video.Description = description
	return video
}

//数据库目前没有保存Title和 Description
func NewVideoByJson(js string) (*Video, error) {
	video := MpVideo{}
	if err := json.Unmarshal([]byte(js), video); err != nil {
		return nil, err
	}
	return NewVideo(video.MediaID, "", ""), nil
}
