package microapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/trrtly/douyin/context"
	"github.com/trrtly/douyin/request"
)

const (
	createQRCodeURL = "https://developer.toutiao.com/api/apps/qrcode"

	// 打开二维码的字节系 app 名称，取值如下
	_ AppnameType = iota

	// AppnameToutiao 今日头条
	AppnameToutiao

	// AppnameDouyin 抖音
	AppnameDouyin

	// AppnamePipixia 皮皮虾
	AppnamePipixia

	// AppnameHuoshan 火山小视频
	AppnameHuoshan
)

// AppnameType 打开二维码的字节系 app 名称
type AppnameType int

// Color struct
type Color struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}

// QRCodeParams 获取小程序/小游戏的二维码的请求参数
type QRCodeParams struct {
	AccessToken string `json:"access_token"`

	// 是打开二维码的字节系 app 名称，默认为今日头条，取值如下表所示
	// | appname | 对应字节系 app |
	// | ------- | ------------  |
	// | toutiao | 今日头条       |
	// | douyin  | 抖音          |
	// | pipixia | 皮皮虾         |
	// | huoshan | 火山小视频      |
	Appname string `json:"appname,omitempty"`

	// 小程序/小游戏启动参数，小程序则格式为 encode({path}?{query})，
	// 小游戏则格式为 JSON 字符串，默认为空
	Path string `json:"path,omitempty"`

	// 二维码宽度，单位 px，最小 280px，最大 1280px，默认为 430px
	Width uint16 `json:"width,omitempty"`

	// 二维码线条颜色，默认为黑色
	LineColor Color `json:"line_color,omitempty"`

	// 二维码背景颜色，默认为透明
	Background Color `json:"background,omitempty"`

	// 是否展示小程序/小游戏 icon，默认不展示
	SetIcon bool `json:"set_icon,omitempty"`
}

func (a AppnameType) String() string {
	switch a {
	case AppnameToutiao:
		return "toutiao"
	case AppnameDouyin:
		return "douyin"
	case AppnamePipixia:
		return "pipixia"
	case AppnameHuoshan:
		return "huoshan"
	}
	return "toutiao"
}

// FectQrCode 获取小程序/小游戏的二维码
// 该二维码可通过任意 app 扫码打开，能跳转到开发者指定的对应字节系 app 内拉起小程序/小游戏，
// 并传入开发者指定的参数。通过该接口生成的二维码，永久有效，暂无数量限制。
func (a *Microapp) FectQrCode(params *QRCodeParams) (response []byte, err error) {
	accessToken, err := a.GetAccessToken()
	if err != nil {
		return
	}
	params.AccessToken = accessToken

	resp, err := request.PostJSON(createQRCodeURL, params)
	if err != nil {
		return
	}

	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		result := &context.Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return
		}
		if result.Code != 0 {
			err = fmt.Errorf("fetchQrCode error : errcode=%v , errmsg=%v", result.Code, result.Msg)
			return
		}
	} else if contentType == "image/jpeg" {
		response, err = ioutil.ReadAll(resp.Body)
		return
	} else {
		err = fmt.Errorf("fetchQrCode error : unknown response content type - %v", contentType)
		return
	}

	return
}
