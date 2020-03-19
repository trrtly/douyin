package context

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/trrtly/douyin/request"
)

const (
	accessTokenURL      = "https://developer.toutiao.com/api/apps/token"
	accessTokenCacheKey = "access_token_%s"
)

//ResAccessToken struct
type ResAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Error
}

//GetAccessTokenFunc 获取 access token 的函数签名
type GetAccessTokenFunc func(ctx *Context) (accessToken string, err error)

//SetAccessTokenLock 设置读写锁（一个appid一个读写锁）
func (ctx *Context) SetAccessTokenLock(l *sync.RWMutex) {
	ctx.accessTokenLock = l
}

//SetGetAccessTokenFunc 设置自定义获取accessToken的方式, 需要自己实现缓存
func (ctx *Context) SetGetAccessTokenFunc(f GetAccessTokenFunc) {
	ctx.accessTokenFunc = f
}

func (ctx *Context) getCacheKey() string {
	return fmt.Sprintf(accessTokenCacheKey, ctx.Appid)
}

//GetAccessToken 获取access_token
func (ctx *Context) GetAccessToken() (accessToken string, err error) {
	ctx.accessTokenLock.Lock()
	defer ctx.accessTokenLock.Unlock()

	if ctx.accessTokenFunc != nil {
		return ctx.accessTokenFunc(ctx)
	}
	accessToken, err = ctx.Cache.GetString(ctx.getCacheKey())
	if accessToken != "" {
		return
	}

	resAccessToken, err := ctx.GetAccessTokenFromServer()
	if err != nil {
		return
	}

	accessToken = resAccessToken.AccessToken
	return
}

//GetAccessTokenFromServer 强制获取 access token
func (ctx *Context) GetAccessTokenFromServer() (resAccessToken ResAccessToken, err error) {
	queryData := map[string]string{
		"appid":      ctx.Appid,
		"secret":     ctx.Secret,
		"grant_type": "client_credential",
	}
	body, err := request.Get(accessTokenURL, queryData)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resAccessToken)
	if err != nil {
		return
	}
	if resAccessToken.Code != 0 {
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resAccessToken.Code, resAccessToken.Msg)
		return
	}

	expires := resAccessToken.ExpiresIn - 1500
	err = ctx.Cache.SetString(ctx.getCacheKey(), resAccessToken.AccessToken, expires)
	return
}
