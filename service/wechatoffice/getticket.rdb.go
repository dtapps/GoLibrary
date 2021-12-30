package wechatoffice

import (
	"fmt"
	"gopkg.in/dtapps/go-library.v3/utils/goredis"
)

// GetJsapiTicketRDb 获取api_ticket
func (app *App) GetJsapiTicketRDb() string {
	cacheName := fmt.Sprintf("wechat_jsapi_ticket:%v", app.AppId)
	redis := goredis.App{
		Rdb: app.RDb,
	}
	newCache := redis.NewSimpleStringCache(redis.NewStringOperation(), 7000)
	newCache.DBGetter = func() string {
		token := app.GetTicket(app.GetAccessTokenRDb(), "jsapi")
		return token.Result.Ticket
	}
	return newCache.GetCache(cacheName)
}