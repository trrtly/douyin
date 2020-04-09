package douyin

import (
	"sync"

	"github.com/trrtly/douyin/cache"
	"github.com/trrtly/douyin/context"
	"github.com/trrtly/douyin/microapp"
)

// Douyin struct
type Douyin struct {
	Context *context.Context
}

// Config struct
type Config struct {
	// Appid 小程序 ID
	Appid string
	// Secret 小程序的 APP Secret，可以在开发者后台获取
	Secret string

	Cache cache.Cache
}

// NewDouyin init
func NewDouyin(c *Config) *Douyin {
	contextObj := &context.Context{
		Appid:  c.Appid,
		Secret: c.Secret,
		Cache:  c.Cache,
	}
	contextObj.SetAccessTokenLock(new(sync.RWMutex))
	return &Douyin{Context: contextObj}
}

// GetMicroApp 获取小程序的实例
func (d *Douyin) GetMicroApp() *microapp.Microapp {
	return microapp.New(d.Context)
}
