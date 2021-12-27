package pinduoduo

// GoodsCatsGetResult 返回参数
type GoodsCatsGetResult struct {
	GoodsCatsGetResponse struct {
		GoodsCatsList []struct {
			CatId       int    `json:"cat_id"`        // 商品类目ID
			CatName     string `json:"cat_name"`      // 商品类目名称
			Level       int    `json:"level"`         // 类目层级，1-一级类目，2-二级类目，3-三级类目，4-四级类目
			ParentCatID int    `json:"parent_cat_id"` // id所属父类目ID，其中，parent_id=0时为顶级节点
		} `json:"goods_cats_list"`
	} `json:"goods_cats_get_response"`
}

// GoodsCatsGet 商品标准类目接口 https://open.pinduoduo.com/application/document/api?id=pdd.goods.cats.get
func (app *App) GoodsCatsGet(parentOptId int) (body []byte, err error) {
	// 参数
	param := NewParams()
	param.Set("parent_cat_id", parentOptId)
	params := NewParamsWithType("pdd.goods.cats.get", param)
	// 请求
	body, err = app.request(params)
	return
}