package order

import (
	"time"

	"github.com/shopspring/decimal"
)

// NftOrder  nft的订单表（同步合约中定义的字段）
type NftOrder struct {
	ID               int64     `gorm:"column:id" db:"id" json:"id" form:"id"`                                                             //  主键
	Nft              string    `gorm:"column:nft" db:"nft" json:"nft" form:"nft"`                                                         //  nft合约地址
	NftID            string    `gorm:"column:nft_id" db:"nft_id" json:"nft_id" form:"nft_id"`                                             //  nft编号的token id
	NftAmount        int64     `gorm:"column:nft_amount" db:"nft_amount" json:"nft_amount" form:"nft_amount"`                             //  nft token id数量，0 代表是erc721协议，非0兼容erc1155协议
	Ft               string    `gorm:"column:ft" db:"ft" json:"ft" form:"ft"`                                                             //  支付代币的合约地址
	FtAmount         string    `gorm:"column:ft_amount" db:"ft_amount" json:"ft_amount" form:"ft_amount"`                                 //  售卖价格，支付代币的数量
	Maker            string    `gorm:"column:maker" db:"maker" json:"maker" form:"maker"`                                                 //  订单创建者地址，maker地址或者是taker地址，也是签名地址
	RoyaltyRecipient string    `gorm:"column:royalty_recipient" db:"royalty_recipient" json:"royalty_recipient" form:"royalty_recipient"` //  nft实际接受者，针对于oferrs订单来说，可定制买单接受者地址
	ServiceFee       float64   `gorm:"column:service_fee" db:"service_fee" json:"service_fee" form:"service_fee"`                         //  平台手续费占比,前端从合约中获取该参数
	RoyaltyRate      float64   `gorm:"column:royalty_rate" db:"royalty_rate" json:"royalty_rate" form:"royalty_rate"`                     //  创作者版权费用，创建合集的时候确定版权占比
	StartAt          time.Time `gorm:"column:start_at" db:"start_at" json:"start_at" form:"start_at"`                                     //  订单开始时间
	ExpireAt         time.Time `gorm:"column:expire_at" db:"expire_at" json:"expire_at" form:"expire_at"`                                 //  订单过期时间
	MakerNonce       int64     `gorm:"column:maker_nonce" db:"maker_nonce" json:"maker_nonce" form:"maker_nonce"`                         //  订单nonce，默认是0，如果是点击批量取消订单后，默认的nonce加1
	TakerGetNft      int64     `gorm:"column:taker_get_nft" db:"taker_get_nft" json:"taker_get_nft" form:"taker_get_nft"`                 //  订单类型，0=taker卖方挂市价单，1=maker买方挂限价单
	Sig              string    `gorm:"column:sig" db:"sig" json:"sig" form:"sig"`                                                         //  签名信息
	OrderHash        string    `gorm:"column:order_hash" db:"order_hash" json:"order_hash" form:"order_hash"`                             //  订单hash(订单id), 后台系统需要关联history表
	ChainID          int64     `gorm:"column:chain_id" db:"chain_id" json:"chain_id" form:"chain_id"`                                     //  支持公链的id, 默认以太坊主网
	ContractType     string    `gorm:"column:contract_type" db:"contract_type" json:"contract_type" form:"contract_type"`                 //  nft合约地址类型
	CollectionID     int64     `gorm:"column:collection_id" db:"collection_id" json:"collection_id" form:"collection_id"`                 //  当前nft token id 是属于那个合集
	Status           int       `gorm:"column:status" db:"status" json:"status" form:"status"`                                             //  当前订单状态，1=listing(挂单)，2=cancel(取消挂单)，3=match(撮合成功)，4=invalid(订单失效)
	Deleted          int64     `gorm:"column:deleted" db:"deleted" json:"deleted" form:"deleted"`                                         //  逻辑删除,  0:未删除, 1:已删除
	CreatedAt        time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`                             //  创建时间
	UpdatedAt        time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`                             //  修改时间
}

func (NftOrder) TableName() string {
	return "nft_order"
}

// NftOperationHistory  nft的订单操作历史
type NftOperationHistory struct {
	ID              int64           `gorm:"column:id" db:"id" json:"id" form:"id"`                                                         //  主键
	ContractAddress string          `gorm:"column:contract_address" db:"contract_address" json:"contract_address" form:"contract_address"` //  nft合约地址
	ContractType    string          `gorm:"column:contract_type" db:"contract_type" json:"contract_type" form:"contract_type"`             //  nft合约地址类型
	TokenID         string          `gorm:"column:token_id" db:"token_id" json:"token_id" form:"token_id"`                                 //  nft编号的token id
	SellAmount      int64           `gorm:"column:sell_amount" db:"sell_amount" json:"sell_amount" form:"sell_amount"`                     //  nft token id数量，0 代表是erc721协议，非0兼容erc1155协议
	Price           decimal.Decimal `gorm:"column:price" db:"price" json:"price" form:"price"`                                             //  售卖价格，支付代币的数量
	PaymentToken    string          `gorm:"column:payment_token" db:"payment_token" json:"payment_token" form:"payment_token"`             //  支付代币的合约地址
	From            string          `gorm:"column:from" db:"from" json:"from" form:"from"`                                                 //  订单from地址
	To              string          `gorm:"column:to" db:"to" json:"to" form:"to"`                                                         //  订单to地址
	OrderHash       string          `gorm:"column:order_hash" db:"order_hash" json:"order_hash" form:"order_hash"`                         //  订单hash(订单id), 后台系统需要关联history表
	EventType       int             `gorm:"column:event_type" db:"event_type" json:"event_type" form:"event_type"`                         //  交易类型：1: mint 铸造， 2: list 挂单，3:cancel 取消挂单， 4: offer 出价单 ， 5: cancel offer 取消报价，6: matched 撮合成功 7 : transfer 转移, 8:expired 订单过期
	Hash            string          `gorm:"column:hash" db:"hash" json:"hash" form:"hash"`                                                 //  交易hash，上架属于链下，没有hash
	ChainID         int64           `gorm:"column:chain_id" db:"chain_id" json:"chain_id" form:"chain_id"`                                 //  支持公链的id, 默认以太坊主网
	Deleted         int64           `gorm:"column:deleted" db:"deleted" json:"deleted" form:"deleted"`
	CreatedAt       time.Time       `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"` //  创建时间
	UpdatedAt       time.Time       `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"` //  修改时间
}

func (NftOperationHistory) TableName() string {
	return "nft_operation_history"
}
