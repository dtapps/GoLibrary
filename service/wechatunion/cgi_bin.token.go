package wechatunion

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 返回参数
type authGetAccessTokenResult struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
	Errcode     int    `json:"errcode"`      // 错误码
	Errmsg      string `json:"errmsg"`       // 错误信息
}

// AuthGetAccessToken
// 接口调用凭证
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (app *App) AuthGetAccessToken() (accessToken string, err error) {
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", app.AppId, app.AppSecret), map[string]interface{}{}, http.MethodGet)
	// 定义
	var result authGetAccessTokenResult
	err = json.Unmarshal(body, &result)
	return result.AccessToken, err
}