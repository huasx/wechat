package message

import "encoding/json"

//Music 音乐消息
type Music struct {
	CommonToken

	Music struct {
		Title        string `xml:"Title"        json:"Title"`
		Description  string `xml:"Description"  json:"Description"`
		MusicURL     string `xml:"MusicUrl"     json:"MusicUrl"`
		HQMusicURL   string `xml:"HQMusicUrl"   json:"HQMusicUrl"`
		ThumbMediaID string `xml:"ThumbMediaId" json:"ThumbMediaId"`
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
func NewMusic(title, description, musicURL, hQMusicURL, thumbMediaID string) *Music {
	music := new(Music)
	music.Music.Title = title
	music.Music.Description = description
	music.Music.MusicURL = musicURL
	music.Music.HQMusicURL = hQMusicURL
	music.Music.ThumbMediaID = thumbMediaID
	return music
}

func NewMusicByJson(js string) (*Music, error) {
	music := &MpMusic{}
	if err := json.Unmarshal([]byte(js), music); err != nil {
		return nil, err
	}

	return NewMusic(music.Title, music.Description, music.MusicURL, music.HQMusicURL, music.ThumbMediaID), nil
}
