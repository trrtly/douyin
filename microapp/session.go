package microapp

import (
	"encoding/json"

	"github.com/trrtly/douyin/context"
	"github.com/trrtly/douyin/request"
)

const (
	code2SessionURL = "https://developer.toutiao.com/api/apps/jscode2session"
)

// RespCode2Session struct
type RespCode2Session struct {
	context.Error

	Openid     string `json:"openid"`      // 用户在当前小程序的 ID，如果请求时有 code 参数才会返回
	SessionKey string `json:"session_key"` // 会话密钥，如果请求时有 code 参数才会返回
	AnonymousOpenid string `json:"anonymous_openid"` // 匿名用户在当前小程序的 ID，如果请求时有 anonymous_code 参数才会返回
}

// Code2Session 获取session_key和openId
func (a *Microapp) Code2Session(code, anonymousCode string) (resp *RespCode2Session, err error) {
	queryData := map[string]string{
		"appid": a.Appid,
		"secret": a.Secret,
		"code": code,
		"anonymous_code": anonymousCode,
	}
	respData, err := request.Get(code2SessionURL, queryData)
	if err != nil {
		return
	}

	resp = &RespCode2Session{}
	err = json.Unmarshal(respData, resp)
	if err != nil {
		return
	}
	
	if resp.Code != 0 {
		err = fmt.Errorf("Code2Session error : errcode=%v , errmsg=%v", resp.Code, resp.Msg)
		return
	}
	return
}
