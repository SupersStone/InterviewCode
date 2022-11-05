package common

// the status of opr
const (
	Success int = iota + 1
	Fail
)

// OrderStatusEnum alias of int
type OrderStatusEnum int

// Order Status
const (
	OrderListing        OrderStatusEnum = 1  //("挂单", 1),
	OrderPayed          OrderStatusEnum = 2  //("已支付", 2),
	OrderRefunding      OrderStatusEnum = 3  //("退款中", 3),
	OrderCanceled       OrderStatusEnum = 4  //("已取消", 4),
	OrderFinished       OrderStatusEnum = 5  // ("已完成", 5),
	OrderFinishedIn     OrderStatusEnum = 6  // ("买入", 6),
	OrderBurnFinished   OrderStatusEnum = 7  // ("链上已销毁", 7),
	OrderExpire         OrderStatusEnum = 8  // ("订单已过期", 8),  // 本期没有
	OrderTransferred    OrderStatusEnum = 9  //("NFT已转移", 9),
	OrderBuyFail        OrderStatusEnum = 11 //("购买失败", 11),
	OrderSharing        OrderStatusEnum = 12 //("分账中", 12),
	OrderSharingSuccess OrderStatusEnum = 13 //("分账成功", 13),

	SellOrderListing      OrderStatusEnum = 100 //("挂单", 100),
	SellOrderCancel       OrderStatusEnum = 101 //("已取消", 101),
	SellOrderLoseEfficacy OrderStatusEnum = 102 //("失效", 102),
	SellOrderSuccess      OrderStatusEnum = 103 //("卖出", 103),
	BuyOrderPayed         OrderStatusEnum = 200 //("已支付", 200),
	BuyOrderFinished      OrderStatusEnum = 201 //("买入", 201),
	BuyOrderFail          OrderStatusEnum = 202 //("购买失败", 202),
	OrderPending          OrderStatusEnum = 300 //("Pending", 300),
)

// Order Status
const (
	Buy int = iota + 1
	Sell
	Cancel
)

// pay style
const (
	PayWithAll int = iota
	PayWithToken
	PayWithSpot
	PayWithTokenOrSpot
	PayWithOtc
	PayWithTokenOrOtc
	PayWithSpotOrOtc
	PayWithTokenOrSpotOrOtc
)

// history type
const (
	NFTMint int = iota + 1
	OrderTransfer
	OrderBurn
	OrderMatchFixed
	OrderCancel
	OrderCancelAll
)

// history type
const (
	HistoryNFTMint int = iota + 1
	HistoryNFTTransfered
	HistoryNFTBurned
	HistoryOrderMatchFixed
	HistoryOrderCanceled
	HistoryOrderCancelAll
)
