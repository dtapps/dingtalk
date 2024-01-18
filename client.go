package dingtalk

import (
	"go.dtapp.net/golog"
)

// ClientConfig 实例配置
type ClientConfig struct {
	Secret      string
	AccessToken string
}

// Client 实例
type Client struct {
	config struct {
		secret      string
		accessToken string
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
	mongoLog struct {
		status bool            // 状态
		client *golog.ApiMongo // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.secret = config.Secret
	c.config.accessToken = config.AccessToken

	return c, nil
}
