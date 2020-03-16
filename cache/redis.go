package cache

import (
	"encoding/json"
	"time"

	"github.com/gomodule/redigo/redis"
)

//Redis redis cache
type Redis struct {
	conn *redis.Pool
}

//RedisOpts redis 连接属性
type RedisOpts struct {
	Host        string `json:"host"`
	Password    string `json:"password"`
	Database    int    `json:"database"`
	MaxIdle     int    `json:"max_idle"`
	MaxActive   int    `json:"max_active"`
	IdleTimeout int32  `json:"idle_timeout"` //second
}

//NewRedis 实例化
func NewRedis(opts *RedisOpts) *Redis {
	pool := &redis.Pool{
		MaxActive:   opts.MaxActive,
		MaxIdle:     opts.MaxIdle,
		IdleTimeout: time.Second * time.Duration(opts.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", opts.Host,
				redis.DialDatabase(opts.Database),
				redis.DialPassword(opts.Password),
			)
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
	return &Redis{pool}
}

//SetConn 设置conn
func (r *Redis) SetConn(conn *redis.Pool) {
	r.conn = conn
}

//GetString 获取 string 数据
func (r *Redis) GetString(key string) (string, error) {
	conn := r.conn.Get()
	defer conn.Close()

	return redis.String(conn.Do("GET", key))
}

//GetJSON 获取 json 数据
func (r *Redis) GetJSON(key string, reply interface{}) (err error) {
	conn := r.conn.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &reply)

	return
}

//SetString 保存 string 数据
func (r *Redis) SetString(key, val string, timeout int64) (err error) {
	conn := r.conn.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", key, timeout, val)

	return
}

//SetJSON 保存 json 数据
func (r *Redis) SetJSON(key string, val interface{}, timeout int64) (err error) {
	conn := r.conn.Get()
	defer conn.Close()

	data, err := json.Marshal(val)
	if err != nil {
		return
	}

	_, err = conn.Do("SETEX", key, timeout, data)

	return
}

//IsExist 判断key是否存在
func (r *Redis) IsExist(key string) bool {
	conn := r.conn.Get()
	defer conn.Close()

	a, _ := conn.Do("EXISTS", key)
	i := a.(int64)
	if i > 0 {
		return true
	}
	return false
}

//Delete 删除
func (r *Redis) Delete(key string) error {
	conn := r.conn.Get()
	defer conn.Close()

	if _, err := conn.Do("DEL", key); err != nil {
		return err
	}

	return nil
}
