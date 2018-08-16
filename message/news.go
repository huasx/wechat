package message

//News 图文消息
type News struct {
	CommonToken

	ArticleCount int        `xml:"ArticleCount"`
	Articles     []*Article `xml:"Articles>item,omitempty"`
}

//NewNews 初始化图文消息
func NewNews(articles []*Article) *News {
	news := new(News)
	news.ArticleCount = len(articles)
	news.Articles = articles
	return news
}

//Article 单篇文章
type Article struct {
	Title        string      `xml:"Title,omitempty" json:"Title,omitempty"`
	Description  string      `xml:"Description,omitempty" json:"Title,omitempty"`
	PicURL       string      `xml:"PicUrl,omitempty" json:"Title,omitempty"`
	URL          string      `xml:"Url,omitempty" json:"Title,omitempty"`
	Author       string      `xml:"-" json:"Author,omitempty"`
	Content      string      `xml:"-" json:"Content,omitempty"`
	ShowCoverPic interface{} `xml:"-" json:"ShowCoverPic,omitempty"`
}

//NewArticle 初始化文章
func NewArticle(title, description, picURL, url string) *Article {
	article := new(Article)
	article.Title = title
	article.Description = description
	article.PicURL = picURL
	article.URL = url
	return article
}
