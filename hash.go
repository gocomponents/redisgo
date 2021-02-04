/**
 * @Author: JYD
 * @Description:list类型
 * @File:  list
 * @Version: 1.0.0
 * @Date: 2021/2/4 15:12
 */

package redisgo

type hashRD struct{}

/**
 * @Description: hash set
 * @receiver l
 * @param key
 * @param filed
 * @param value
 * @param exist
 * @return *Reply
 */
func (h *hashRD) HSet(key string, filed, value interface{}, exist ...bool) *Reply {
	conn := pool.Get()
	defer conn.Close()
	if len(exist) > 0 && exist[0] {
		return getReply(conn.Do("hsetex", key, filed, value))
	}
	return getReply(conn.Do("hset", key, filed, value))
}

/**
 * @Description: 获取指定字段值
 * @receiver h
 * @param key
 * @param filed
 * @return *Reply
 */
func (h *hashRD) HGet(key string, filed interface{}) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("hget", key, filed))
}

/**
 * @Description: 获取所有字段及值
 * @receiver h
 * @param key
 * @return *Reply
 */
func (h *hashRD) HGetAll(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("hgetall", key))
}
