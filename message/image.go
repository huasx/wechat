package message

import "encoding/json"

//Image 图片消息
type Image struct {
	CommonToken

	Image struct {
		MediaID CDATA `xml:"MediaId" json:"MediaId"`
	} `xml:"Image"`
}

type MpImage struct {
	MediaID string `xml:"MediaId" json:"MediaId"`
}

//NewImage 回复图片消息
func NewImage(MpImage *MpImage) *Image {
	image := new(Image)
	image.Image.MediaID = CDATA{Value: MpImage.MediaID}
	return image
}

func NewImageByJson(js string) (*Image, error) {
	MpImage := &MpImage{}
	if err := json.Unmarshal([]byte(js), MpImage); err != nil {
		return nil, err
	}

	return NewImage(MpImage), nil
}
