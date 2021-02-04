/**
 * @Author: JYD
 * @Description:config相关配置，使用viper
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2021/2/4 13:15
 */

package redisgo

type Config struct {
	Host      string
	Port      int64
	Db        int64
	Password  string
	MaxIdle   int //default 10  最大连接数
	MaxActive int //default 10000 连接池最大数
	Wait      bool
}

/**
 * @Description: 根据路径获取配置文件
 * @param path
 * @return *Config
 * @return error
 */
func getConfigWithFile(path string) (*Config, error) {
	return &Config{}, nil
}
