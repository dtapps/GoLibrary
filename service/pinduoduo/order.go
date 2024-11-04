package pinduoduo

var (
	OrderStatusType = []int64{0, 1, 2, 3, 4, 5, 10}
	OrderStatusDesc = []string{"已支付", "已成团", "确认收货", "审核失败(不可提现)", "已经结算", "已处罚"}
)

// GetOrderStatusDesc 订单状态
func GetOrderStatusDesc(Type int64) (desc string) {
	for i, v := range OrderStatusType {
		if v == Type {
			desc = OrderStatusDesc[i]
			break
		}
	}
	return
}

var (
	OrderSubsidyType = []int64{0, 1, 2, 3, 4, 5, 8}
	OrderSubsidyDesc = []string{"非补贴订单", "千万补贴", "社群补贴", "多多星选", "品牌优选", "千万神券", "拼团享多多"}
)

// GetOrderSubsidyDesc 订单补贴类型
func GetOrderSubsidyDesc(Type int64) (desc string) {
	for i, v := range OrderSubsidyType {
		if v == Type {
			desc = OrderSubsidyDesc[i]
			break
		}
	}
	return
}

var (
	OrderTypeType = []int64{0, 1, 4, 7, 8, 9, 77, 94, 101, 103, 104, 105}
	OrderTypeDesc = []string{"单品推广", "红包活动推广", "多多进宝商城推广", "今日爆款", "品牌清仓", "1.9包邮", "刮刮卡活动推广", "充值中心", "品牌黑卡", "百亿补贴频道", "内购清单频道", "超级红包"}
)

// GetOrderTypeDesc 下单场景类型
func GetOrderTypeDesc(Type int64) (desc string) {
	for i, v := range OrderTypeType {
		if v == Type {
			desc = OrderTypeDesc[i]
			break
		}
	}
	return
}

var (
	OrderBandanRiskConsultType = []int64{-1, 0, 1}
	OrderBandanRiskConsultDesc = []string{"未出结果", "不是代购订单", "是代购订单"}
)

// GetOrderBandanRiskConsultDesc 预判断是否为代购订单
func GetOrderBandanRiskConsultDesc(Type int64) (desc string) {
	for i, v := range OrderBandanRiskConsultType {
		if v == Type {
			desc = OrderBandanRiskConsultDesc[i]
			break
		}
	}
	return
}

var (
	OrderIsDirectType = []int64{0, 1}
	OrderIsDirectDesc = []string{"否", "是"}
)

// GetOrderIsDirectDesc 下单场景类型
func GetOrderIsDirectDesc(Type int64) (desc string) {
	for i, v := range OrderIsDirectType {
		if v == Type {
			desc = OrderIsDirectDesc[i]
			break
		}
	}
	return
}

var (
	OrderPriceCompareStatusType = []int64{0, 1}
	OrderPriceCompareStatusDesc = []string{"正常", "比价"}
)

// GetOrderPriceCompareStatusDesc 下单场景类型
func GetOrderPriceCompareStatusDesc(Type int64) (desc string) {
	for i, v := range OrderPriceCompareStatusType {
		if v == Type {
			desc = OrderPriceCompareStatusDesc[i]
			break
		}
	}
	return
}