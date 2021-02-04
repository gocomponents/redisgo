/**
 * @Author: JYD
 * @Description:redis连接池
 * @File:  conn
 * @Version: 1.0.0
 * @Date: 2021/2/4 13:30
 */

package redisgo

import (
	log "github.com/sirupsen/logrus"
)

type redisConn struct {
	String stringRD
	Lock   lockRD
}

var RedisConn = new(redisConn)

func NewConnectionWithFile(path string) error {
	config, err := getConfigWithFile(path)
	if err != nil {
		log.Error("获取Redis配置信息失败,", path)
		return err
	}
	initPool(config)
	log.Info("redis 配置为: ", config)
	return nil
}
