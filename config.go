package dingtalk

func (c *Client) Config(secret, accessToken string) *Client {
	c.config.secret = secret
	c.config.accessToken = accessToken
	return c
}
