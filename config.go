package dingtalk

import "go.dtapp.net/golog"

func (c *Client) Config(secret, accessToken string) *Client {
	c.config.secret = secret
	c.config.accessToken = accessToken
	return c
}

// ConfigApiClientFun 日志配置
func (c *Client) ConfigApiClientFun(apiClientFun golog.ApiClientFun) {
	apiClient := apiClientFun()
	if apiClient != nil {
		c.log.client = apiClient
		c.log.status = true
	}
}
