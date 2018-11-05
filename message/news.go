package message

import (
	"app/core/weixin"
	"encoding/json"
)

//News 图文消息
type News struct {
	CommonToken

	ArticleCount int        `xml:"ArticleCount"`
	Articles     []*Article `xml:"Articles>item,omitempty"`
}

//NewNews 初始化图文消息
func NewNews(articles []*Article) *News {
	articleContent := make([]*Article, 0)
	for _, value := range articles {
		a := &Article{}
		a.URL = value.URL
		a.PicURL = weixin.StaticHost + value.PicURL + weixin.OssThumbNail
		a.Title = value.Title
		a.Description = value.Description
		articleContent = append(articleContent, a)
		break
	}

	news := new(News)
	news.ArticleCount = 1 //微信升级之后 只能是单图文
	news.Articles = articleContent
	return news
}

//Article 单篇文章
type Article struct {
	Title       string `xml:"Title,omitempty" json:"Title,omitempty"`
	Description string `xml:"Description,omitempty" json:"Description,omitempty"`
	PicURL      string `xml:"PicUrl,omitempty" json:"PicUrl,omitempty"`
	URL         string `xml:"Url,omitempty" json:"Url,omitempty"`
}

func NewNewsByJson(js string) (*News, error) {
	article := make([]*Article, 0)
	if err := json.Unmarshal([]byte(js), &article); err != nil {
		return nil, err
	}

	return NewNews(article), nil
}
