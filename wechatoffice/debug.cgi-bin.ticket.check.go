package wechatoffice

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type DebugCgiBinTicketCheckResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type DebugCgiBinTicketCheckResult struct {
	Result DebugCgiBinTicketCheckResponse // 结果
	Body   []byte                         // 内容
	Http   gorequest.Response             // 请求
	Err    error                          // 错误
}

func NewDebugCgiBinTicketCheckResult(result DebugCgiBinTicketCheckResponse, body []byte, http gorequest.Response, err error) *DebugCgiBinTicketCheckResult {
	return &DebugCgiBinTicketCheckResult{Result: result, Body: body, Http: http, Err: err}
}

// DebugCgiBinTicketCheck 判断ticket是否合法
// https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=jsapisign
func (app *App) DebugCgiBinTicketCheck() *DebugCgiBinTicketCheckResult {
	app.jsapiTicket = app.GetJsapiTicket()
	// 请求
	request, err := app.request(fmt.Sprintf("https://mp.weixin.qq.com/debug/cgi-bin/ticket/check?ticket=%s", app.jsapiTicket), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response DebugCgiBinTicketCheckResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewDebugCgiBinTicketCheckResult(response, request.ResponseBody, request, err)
}