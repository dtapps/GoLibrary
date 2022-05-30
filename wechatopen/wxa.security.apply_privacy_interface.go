package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type WxaSecurityApplyPrivacyInterfaceResponse struct {
	Errcode int    `json:"errcode"`  // 返回码
	Errmsg  string `json:"errmsg"`   // 返回码信息
	AuditId int64  `json:"audit_id"` // 审核单id
}

type WxaSecurityApplyPrivacyInterfaceResult struct {
	Result WxaSecurityApplyPrivacyInterfaceResponse // 结果
	Body   []byte                                   // 内容
	Http   gorequest.Response                       // 请求
	Err    error                                    // 错误
}

func NewWxaSecurityApplyPrivacyInterfaceResult(result WxaSecurityApplyPrivacyInterfaceResponse, body []byte, http gorequest.Response, err error) *WxaSecurityApplyPrivacyInterfaceResult {
	return &WxaSecurityApplyPrivacyInterfaceResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaSecurityApplyPrivacyInterface 申请接口
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/apply_api/apply_privacy_interface.html
func (app *App) WxaSecurityApplyPrivacyInterface(notMustParams ...Params) *WxaSecurityApplyPrivacyInterfaceResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/security/apply_privacy_interface?access_token=%s", app.GetAuthorizerAccessToken()), params, http.MethodPost)
	// 定义
	var response WxaSecurityApplyPrivacyInterfaceResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewWxaSecurityApplyPrivacyInterfaceResult(response, request.ResponseBody, request, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaSecurityApplyPrivacyInterfaceResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 61031:
		return "审核中，请不要重复申请"
	case 61032:
		return "视频格式不对，要传mp4格式的"
	case 61033:
		return "视频下载失败"
	case 61034:
		return "必填的参数没填，检查后重新提交"
	case 61035:
		return "输入的api（api_name严格区分大小写）无需申请，可以直接使用"
	case 61036:
		return "该帐号不可申请，请检查类目是否符合"
	case 61037:
		return "需要以ntf-8的编码格式提交"
	}
	return "系统繁忙"
}