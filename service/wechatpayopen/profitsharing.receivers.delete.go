package wechatpayopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ProfitSharingReceiversDeleteResponse struct {
	SubMchid string `json:"sub_mchid"` // 子商户号
	Type     string `json:"type"`      // 分账接收方类型
	Account  string `json:"account"`   // 分账接收方账号
}

type ProfitSharingReceiversDeleteResult struct {
	Result ProfitSharingReceiversDeleteResponse // 结果
	Body   []byte                               // 内容
	Http   gorequest.Response                   // 请求
}

func newProfitSharingReceiversDeleteResult(result ProfitSharingReceiversDeleteResponse, body []byte, http gorequest.Response) *ProfitSharingReceiversDeleteResult {
	return &ProfitSharingReceiversDeleteResult{Result: result, Body: body, Http: http}
}

// ProfitSharingReceiversDelete 删除分账接收方API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_9.shtml
func (c *Client) ProfitSharingReceiversDelete(ctx context.Context, Type, account string, notMustParams ...*gorequest.Params) (*ProfitSharingReceiversDeleteResult, ApiError, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	params.Set("appid", c.GetSpAppid())      // 应用ID
	params.Set("sub_appid", c.GetSubAppid()) // 子商户应用ID
	params.Set("type", Type)                 // 分账接收方类型
	if Type == MERCHANT_ID {
		params.Set("account", account) // 商户号
	}
	if Type == PERSONAL_OPENID {
		params.Set("account", account) // 个人openid
	}
	if Type == PERSONAL_SUB_OPENID {
		params.Set("account", account) // 个人sub_openid
	}

	// 请求
	var response ProfitSharingReceiversDeleteResponse
	var apiError ApiError
	request, err := c.request(ctx, "v3/profitsharing/receivers/delete", params, http.MethodPost, &response, &apiError)
	return newProfitSharingReceiversDeleteResult(response, request.ResponseBody, request), apiError, err
}
