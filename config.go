package dingtalk

import (
	"go.dtapp.net/golog"
)

func (c *Client) Config(secret, accessToken string) *Client {
	c.config.secret = secret
	c.config.accessToken = accessToken
	return c
}

// ConfigApiGormFun 接口日志配置
func (c *Client) ConfigApiGormFun(apiClientFun golog.ApiGormFun) {
	client := apiClientFun()
	if client != nil {
		c.gormLog.client = client
		c.gormLog.status = true
	}
}
