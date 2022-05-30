package pinduoduo

import (
	"encoding/json"
	"go.dtapp.net/library/gorequest"
)

type MemberAuthorityQueryResponse struct {
	AuthorityQueryResponse struct {
		Bind      int    `json:"bind"`
		RequestId string `json:"request_id"`
	} `json:"authority_query_response"`
}

type MemberAuthorityQueryResult struct {
	Result MemberAuthorityQueryResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
	Err    error                        // 错误
}

func NewMemberAuthorityQueryResult(result MemberAuthorityQueryResponse, body []byte, http gorequest.Response, err error) *MemberAuthorityQueryResult {
	return &MemberAuthorityQueryResult{Result: result, Body: body, Http: http, Err: err}
}

// MemberAuthorityQuery 查询是否绑定备案
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.search
func (app *App) MemberAuthorityQuery(notMustParams ...Params) *MemberAuthorityQueryResult {
	// 参数
	params := NewParamsWithType("pdd.ddk.member.authority.query", notMustParams...)
	params.Set("pid", app.Pid)
	// 请求
	request, err := app.request(params)
	// 定义
	var response MemberAuthorityQueryResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewMemberAuthorityQueryResult(response, request.ResponseBody, request, err)
}