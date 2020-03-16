package douyin

// Douyin struct
type Douyin struct {
	*Config
}

// Config struct
type Config struct {
	// Appid 小程序 ID
	Appid string
	// Secret 小程序的 APP Secret，可以在开发者后台获取
	Secret string
}

// NewDouyin init
func NewDouyin(c *Config) *Douyin {
	return &Douyin{c}
}
