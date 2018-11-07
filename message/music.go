package message

import (
	"app/core/weixin"
	"encoding/json"
)

//Music 音乐消息
type Music struct {
	CommonToken

	Music struct {
		Title        CDATA `xml:"Title"        json:"Title"`
		Description  CDATA `xml:"Description"  json:"Description"`
		MusicURL     CDATA `xml:"MusicUrl"     json:"MusicUrl"`
		HQMusicURL   CDATA `xml:"HQMusicUrl"   json:"HQMusicUrl"`
		ThumbMediaID CDATA `xml:"ThumbMediaId" json:"ThumbMediaId"`
	} `xml:"Music"`
}

type MpMusic struct {
	Title        string `xml:"Title"        json:"Title"`
	Description  string `xml:"Description"  json:"Description"`
	MusicURL     string `xml:"MusicUrl"     json:"MusicUrl"`
	HQMusicURL   string `xml:"HQMusicUrl"   json:"HQMusicUrl"`
	ThumbMediaID string `xml:"ThumbMediaId" json:"ThumbMediaId"`
}

//NewMusic  回复音乐消息
func NewMusic(MpMusic *MpMusic) *Music {
	music := new(Music)
	music.Music.Title 		 = CDATA{Value: MpMusic.Title}
	music.Music.Description  = CDATA{Value: MpMusic.Description}
	music.Music.MusicURL     = CDATA{Value: weixin.StaticHost + MpMusic.MusicURL}
	music.Music.HQMusicURL   = CDATA{Value: weixin.StaticHost + MpMusic.HQMusicURL}
	music.Music.ThumbMediaID = CDATA{Value: MpMusic.ThumbMediaID}
	return music
}

func NewMusicByJson(js string) (*Music, error) {
	music  := &MpMusic{}
	if err := json.Unmarshal([]byte(js), music); err != nil {
		return nil, err
	}

	return NewMusic(music), nil
}
