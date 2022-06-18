package wechatqy

import (
	"encoding/json"
	"fmt"
	gorequest2 "go.dtapp.net/library/utils/gorequest"
)

type WebhookSendResponse struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

type WebhookSendResult struct {
	Result WebhookSendResponse // 结果
	Body   []byte              // 内容
	Http   gorequest2.Response // 请求
	Err    error               // 错误
}

func NewWebhookSendResult(result WebhookSendResponse, body []byte, http gorequest2.Response, err error) *WebhookSendResult {
	return &WebhookSendResult{Result: result, Body: body, Http: http, Err: err}
}

// WebhookSend 发送应用消息
// https://developer.work.weixin.qq.com/document/path/90372
func (app *App) WebhookSend(notMustParams ...gorequest2.Params) *WebhookSendResult {
	// 参数
	params := gorequest2.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s&type=%s", app.key, "text"), params)
	// 定义
	var response WebhookSendResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewWebhookSendResult(response, request.ResponseBody, request, err)
}