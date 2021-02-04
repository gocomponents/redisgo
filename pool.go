/**
 * @Author: JYD
 * @Description:
 * @File:  pool
 * @Version: 1.0.0
 * @Date: 2021/2/4 13:34
 */

package redisgo

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"time"
)

var pool *redis.Pool

func initPool(config *Config) {
	pool = &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: 30 * time.Second,
		Wait:        config.Wait,
		Dial: func() (redis.Conn, error) {
			return setDialog(config)
		},
	}
}

func setDialog(config *Config) (redis.Conn, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Error(fmt.Sprintf("init redis failed! %v", config))
	}
	if len(config.Password) != 0 {
		if _, err := conn.Do("AUTH", config.Password); err != nil {
			conn.Close()
			log.Error(err)
		}
	}
	if _, err := conn.Do("SELECT", config.Db); err != nil {
		conn.Close()
		log.Error(err)
	}
	r, err := redis.String(conn.Do("PING"))
	if err != nil || r != "PONG" {
		panic("连接失败")
	}
	return conn, nil
}
