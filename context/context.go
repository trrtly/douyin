package context

import (
	"net/http"
	"sync"

	"github.com/trrtly/douyin/cache"
)

// Context struct
type Context struct {
	// Appid 小程序 ID
	Appid string
	// Secret 小程序的 APP Secret，可以在开发者后台获取
	Secret string

	// Cache 缓存
	Cache cache.Cache

	Writer  http.ResponseWriter
	Request *http.Request

	//accessTokenLock 读写锁 同一个Appid一个
	accessTokenLock *sync.RWMutex
}
