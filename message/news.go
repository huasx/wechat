package message

import (
	"app/core/weixin"
	"encoding/json"
	"strings"
)

//News 图文消息
type News struct {
	CommonToken

	ArticleCount int        `xml:"ArticleCount"`
	Articles     []*Article `xml:"Articles>item,omitempty"`
}

//NewNews 初始化图文消息
func NewNews(MpArticle []*MpArticle) *News {
	articleContent := make([]*Article, 0)
	for _, value := range MpArticle {
		a 			  := &Article{}
		a.URL          = CDATA{Value: value.URL}
		a.Title        = CDATA{Value: value.Title}
		a.Description  = CDATA{Value: value.Description}

		if !strings.Contains(value.PicURL, "https://")  && !strings.Contains(value.PicURL, "http://") {
			a.PicURL   = CDATA{Value: weixin.StaticHost + value.PicURL + weixin.OssThumbNail}
		} else {
			a.PicURL   = CDATA{Value: value.PicURL}
		}

		articleContent = append(articleContent, a)
		break
	}

	news := new(News)
	news.ArticleCount = 1 //微信升级之后, 只能是单图文
	news.Articles     = articleContent
	return news
}

//Article 单篇文章 都是必传参数
type Article struct {
	Title       CDATA `xml:"Title" json:"Title"`
	Description CDATA `xml:"Description" json:"Description"`
	PicURL      CDATA `xml:"PicUrl" json:"PicUrl"`
	URL         CDATA `xml:"Url" json:"Url"`
}

type MpArticle struct {
	Title       string `xml:"Title" json:"Title"`
	Description string `xml:"Description" json:"Description"`
	PicURL      string `xml:"PicUrl" json:"PicUrl"`
	URL         string `xml:"Url" json:"Url"`
}

func NewNewsByJson(js string) (*News, error) {
	MpArticle := make([]*MpArticle, 0)
	if err := json.Unmarshal([]byte(js), &MpArticle); err != nil {
		return nil, err
	}
	return NewNews(MpArticle), nil
}
