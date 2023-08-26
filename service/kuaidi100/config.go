package kuaidi100

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) Config(customer string) *Client {
	c.config.customer = customer
	return c
}

// ConfigZapClientFun 日志配置
func (c *Client) ConfigZapClientFun(apiZapLogFun golog.ApiZapLogFun) {
	apiZapLog := apiZapLogFun()
	if apiZapLog != nil {
		c.zap.client = apiZapLog
		c.zap.status = true
	}
}