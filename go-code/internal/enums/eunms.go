package enums

// order type
const (
	OrderListing int = iota + 1
	OrderCanceled
	OrderMatched
	OrderInValid
)

// history type
const (
	HistoryNFTMint         int = iota + 1 //铸造
	HistoryNFTTransfered                  // 转移
	HistoryNFTBurned                      // 销毁
	HistoryOrderMatchFixed                // 撮合成功
	HistoryOrderCanceled                  // 取消订单
	HistoryOrderCancelAll                 // 取消全部订单
	HistoryOrderList                      // 挂单
	HistoryOrderExpires                   // 订单过期
)

// sort type
const (
	DefaultType int = iota
	PriceLowToHigth
	PriceHigthToLow
	NewList
	NewSale
)

// buy now  type
const (
	BuyNowStatus int = iota + 1
	NoList
)

// Price Token
const (
	BuyEthToken int = iota + 1
	BuyusdToken
)

// Attributes  Parameters
const (
	OneParameters int = iota + 1
	TwoParameters
	ThreeParameters
	FourParameters
	FiveParameters
	SixParameters
	SevenParameters
	EightParameters
	NineParameters
	TenParameters
)

// upload image
const (
	DefaultFileName int64 = iota
	UserLogoFileName
	UserBackgorundFileName
	CollectionLogoFileName
	CollectionCoverFileName
	CollectionFeaturedFileName
)

// personal nft status
const (
	PersonalNftStatusList int = iota + 1
	PersonalNftStatusNoList
)

// personal nft sort
const (
	PersonalNftSale int = iota + 1
	PersonalNftSell
	PersonalNftLowToHight
	PersonalNftHightToLow
)

// personcal nft price
const (
	PersonalNftEthPay int = iota + 1
	PersonalNftUsdPay
)
