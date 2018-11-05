package message

import "encoding/json"

//Image 图片消息
type Image struct {
	CommonToken

	Image struct {
		MediaID string `xml:"MediaId" json:"MediaId"`
	} `xml:"Image"`
}

type MpImage struct {
	MediaID string `xml:"MediaId" json:"MediaId"`
}

//NewImage 回复图片消息
func NewImage(mediaID string) *Image {
	image := new(Image)
	image.Image.MediaID = mediaID
	return image
}

func NewImageByJson(js string) (*Image, error) {
	MpImage := &MpImage{}
	if err := json.Unmarshal([]byte(js), MpImage); err != nil {
		return nil, err
	}

	return NewImage(MpImage.MediaID), nil
}
