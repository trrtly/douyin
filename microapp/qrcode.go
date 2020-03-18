package microapp

const (
	createQRCodeURL = "https://developer.toutiao.com/api/apps/qrcode"
)

// QRCodeParams 获取小程序/小游戏的二维码的请求参数
type QRCodeParams struct {
	// 是打开二维码的字节系 app 名称，默认为今日头条，取值如下表所示
	// | appname | 对应字节系 app |
	// | ------- | ------------  |
	// | toutiao | 今日头条       |
	// | douyin  | 抖音          |
	// | pipixia | 皮皮虾         |
	// | huoshan | 火山小视频      |
	Appname string `json:"appname"`

	// 小程序/小游戏启动参数，小程序则格式为 encode({path}?{query})，
	// 小游戏则格式为 JSON 字符串，默认为空
	Path string `json:"path"`

	// 二维码宽度，单位 px，最小 280px，最大 1280px，默认为 430px
	Width uint16 `json:"width"`

	// 二维码线条颜色，默认为黑色
	LineColor Color `json:"line_color"`

	// 二维码背景颜色，默认为透明
	Background Color `json:"background"`

	// 是否展示小程序/小游戏 icon，默认不展示
	SetIcon bool `json:"set_icon"`
}

// Color struct
type Color struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}
