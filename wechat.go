package wechat

import (
	"net/http"
	"sync"

	"github.com/huasx/wechat/cache"
	"github.com/huasx/wechat/context"
	"github.com/huasx/wechat/js"
	"github.com/huasx/wechat/material"
	"github.com/huasx/wechat/menu"
	"github.com/huasx/wechat/oauth"
	"github.com/huasx/wechat/pay"
	"github.com/huasx/wechat/server"
	"github.com/huasx/wechat/template"
	"github.com/huasx/wechat/user"
)

// Wechat struct
type Wechat struct {
	Context *context.Context
}

// Config for user
type Config struct {
	AppID string
	//AppSecret      string//弃用 安全级别高
	ClientCode     string
	ClientSecret   string
	Token          string
	EncodingAESKey string
	PayMchID       string //支付 - 商户 ID
	PayNotifyURL   string //支付 - 接受微信支付结果通知的接口地址
	PayKey         string //支付 - 商户后台设置的支付 key
	Cache          cache.Cache
}

// NewWechat init
func NewWechat(cfg *Config) *Wechat {
	ctx := new(context.Context)
	copyConfigToContext(cfg, ctx)
	return &Wechat{ctx}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppID = cfg.AppID
	//context.AppSecret = cfg.AppSecret
	context.ClientCode = cfg.ClientCode
	context.ClientSecret = cfg.ClientSecret
	context.Token = cfg.Token
	context.EncodingAESKey = cfg.EncodingAESKey
	//context.PayMchID = cfg.PayMchID
	//context.PayKey = cfg.PayKey
	//context.PayNotifyURL = cfg.PayNotifyURL
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
	context.SetJsAPITicketLock(new(sync.RWMutex))
}

// GetServer 消息管理
func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return server.NewServer(wc.Context)
}

//GetAccessToken 获取access_token
func (wc *Wechat) GetAccessToken() (string, error) {
	return wc.Context.GetAccessToken()
}

// GetOauth oauth2网页授权
func (wc *Wechat) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(wc.Context)
}

// GetMaterial 素材管理
func (wc *Wechat) GetMaterial() *material.Material {
	return material.NewMaterial(wc.Context)
}

// GetJs js-sdk配置
func (wc *Wechat) GetJs() *js.Js {
	return js.NewJs(wc.Context)
}

// GetMenu 菜单管理接口
func (wc *Wechat) GetMenu() *menu.Menu {
	return menu.NewMenu(wc.Context)
}

// GetUser 用户管理接口
func (wc *Wechat) GetUser() *user.User {
	return user.NewUser(wc.Context)
}

// GetTemplate 模板消息接口
func (wc *Wechat) GetTemplate() *template.Template {
	return template.NewTemplate(wc.Context)
}

// GetPay 返回支付消息的实例
func (wc *Wechat) GetPay() *pay.Pay {
	return pay.NewPay(wc.Context)
}
