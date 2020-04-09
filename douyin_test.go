package douyin

import (
	"testing"

	"github.com/trrtly/douyin/cache"
	"github.com/trrtly/douyin/microapp"
)

func TestNewDouyin(t *testing.T) {

	dy := NewDouyin(&Config{
		Appid:  "xxx",
		Secret: "xxx",
		Cache: cache.NewRedis(&cache.RedisOpts{
			Host:        "127.0.0.1:6379",
			Password:    "xxx",
			Database:    3,
			MaxIdle:     30,
			MaxActive:   30,
			IdleTimeout: 200,
		}),
	})
	testMicroApp(t, dy.GetMicroApp())
}

func testMicroApp(t *testing.T, mapp *microapp.Microapp) {
	t.Run("Code2Session", func(t *testing.T) {
		resp, err := mapp.Code2Session("xxxxx", "")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
	})

	t.Run("FectQrCode", func(t *testing.T) {
		resp, err := mapp.FectQrCode(&microapp.QRCodeParams{
			Appname: microapp.AppnameDouyin.String(),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
	})
}
