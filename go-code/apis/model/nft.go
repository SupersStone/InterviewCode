package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// NftBaseInfo  nft基础信息资产表
type NftBaseInfo struct {
	ID              int64     `gorm:"column:id" db:"id" json:"id" form:"id"`                                                         //  主键
	ContractAddress string    `gorm:"column:contract_address" db:"contract_address" json:"contract_address" form:"contract_address"` //  nft合约地址
	ContractType    string    `gorm:"column:contract_type" db:"contract_type" json:"contract_type" form:"contract_type"`             //  nft合约地址类型，721， 1155等协议类型
	ContractCreator string    `gorm:"column:contract_creator" db:"contract_creator" json:"contract_creator" form:"contract_creator"` //  nft合约地址的创建者
	TokenId         string    `gorm:"column:token_id" db:"token_id" json:"token_id" form:"token_id"`                                 //  nft合约token id
	TokenAmount     string    `gorm:"column:token_amount" db:"token_amount" json:"token_amount" form:"token_amount"`                 //  nft token id数量，0 代表是erc721协议，非0兼容erc1155协议
	Name            string    `gorm:"column:name" db:"name" json:"name" form:"name"`                                                 //  nft 名字，合集名字#token_id组合形成
	CollectionId    int64     `gorm:"column:collection_id" db:"collection_id" json:"collection_id" form:"collection_id"`             //  当前nft token id 是属于那个合集
	MetadataUrl     string    `gorm:"column:metadata_url" db:"metadata_url" json:"metadata_url" form:"metadata_url"`                 //  ipfs的url地址
	ImageUrl        string    `gorm:"column:image_url" db:"image_url" json:"image_url" form:"image_url"`                             //  aws 的 url 图片存储地址
	ChainId         int64     `gorm:"column:chain_id" db:"chain_id" json:"chain_id" form:"chain_id"`                                 //  支持公链的id, 默认以太坊主网
	CreatorRateFee  float64   `gorm:"column:creator_rate_fee" db:"creator_rate_fee" json:"creator_rate_fee" form:"creator_rate_fee"` //  创建者手续费比例
	NftVerify       int64     `gorm:"column:nft_verify" db:"nft_verify" json:"nft_verify" form:"nft_verify"`                         //  nft合约地址是否验证开源，默认为0，未开源， 1代表开源
	Deleted         int64     `gorm:"column:deleted" db:"deleted" json:"deleted" form:"deleted"`                                     //  逻辑删除,  0:未删除, 1:已删除
	CreatedAt       time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`                         //  创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`                         //  修改时间
}

func (NftBaseInfo) TableName() string {
	return "nft_base_info"
}

const DefaultPageSize = 10
const DefaultPageSizeMax = 100
const DefaultPageNumber = 1

// nft的订单表（同步合约中定义的字段）
type NftOrder struct {
	Id               uint64          `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:主键" json:"id"`
	Nft              string          `gorm:"column:nft;type:varchar(255);default:0;comment:nft合约地址;NOT NULL" json:"nft"`
	NftId            string          `gorm:"column:nft_id;type:varchar(500);default:0;comment:nft编号的token id;NOT NULL" json:"nft_id"`
	NftAmount        int64           `gorm:"column:nft_amount;type:bigint(20);default:0;comment:nft token id数量，0 代表是erc721协议，非0兼容erc1155协议;NOT NULL" json:"nft_amount"`
	Ft               string          `gorm:"column:ft;type:varchar(255);default:0;comment:支付代币的合约地址;NOT NULL" json:"ft"`
	FtAmount         decimal.Decimal `gorm:"column:ft_amount;type:decimal(64);comment:售卖价格，支付代币的数量" json:"ft_amount"`
	Maker            string          `gorm:"column:maker;type:varchar(255);default:0;comment:订单创建者地址，maker地址或者是taker地址，也是签名地址;NOT NULL" json:"maker"`
	RoyaltyRecipient string          `gorm:"column:royalty_recipient;type:varchar(255);default:0;comment:nft实际接受者，针对于oferrs订单来说，可定制买单接受者地址;NOT NULL" json:"royalty_recipient"`
	ServiceFee       decimal.Decimal `gorm:"column:service_fee;type:decimal(64);default:0;comment:平台手续费占比,前端从合约中获取该参数" json:"service_fee"`
	RoyaltyRate      decimal.Decimal `gorm:"column:royalty_rate;type:decimal(64);default:0;comment:创作者版权费用，创建合集的时候确定版权占比" json:"royalty_rate"`
	StartAt          int64           `gorm:"column:start_at;type:bigint(20);default:0;comment:订单开始时间;NOT NULL" json:"start_at"`
	ExpireAt         int64           `gorm:"column:expire_at;type:bigint(20);default:0;comment:订单过期时间;NOT NULL" json:"expire_at"`
	MakerNonce       int64           `gorm:"column:maker_nonce;type:bigint(20);comment:订单nonce，默认是0，如果是点击批量取消订单后，默认的nonce加1" json:"maker_nonce"`
	TakerGetNft      int             `gorm:"column:taker_get_nft;type:int(11);comment:订单类型，0=taker卖方挂市价单，1=maker买方挂限价单;NOT NULL" json:"taker_get_nft"`
	Sig              string          `gorm:"column:sig;type:varchar(500);default:0;comment:签名信息;NOT NULL" json:"sig"`
	OrderHash        string          `gorm:"column:order_hash;type:varchar(500);default:0;comment:订单hash(订单id), 后台系统需要关联history表;NOT NULL" json:"order_hash"`
	ChainId          int             `gorm:"column:chain_id;type:int(11);default:1;comment:支持公链的ID, 默认以太坊主网;NOT NULL" json:"chain_id"`
	ContractType     string          `gorm:"column:contract_type;type:varchar(255);default:0;comment:nft合约地址类型;NOT NULL" json:"contract_type"`
	CollectionId     int64           `gorm:"column:collection_id;type:bigint(20);default:0;comment:当前nft token id 是属于那个合集" json:"collection_id"`
	Status           int             `gorm:"column:status;type:int(11);comment:当前订单状态，1=Listing(挂单)，2=Cancel(取消挂单)，3=Match(撮合成功)，4=Invalid(订单失效)" json:"status"`
	Deleted          int             `gorm:"column:deleted;type:tinyint(1);default:0;comment:逻辑删除,  0:未删除, 1:已删除;NOT NULL" json:"deleted"`
	CreatedAt        time.Time       `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt        time.Time       `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:修改时间;NOT NULL" json:"updated_at"`
}

func (m *NftOrder) TableName() string {
	return "nft_order"
}

// NftOperationHistory  nft的订单操作历史
type NftOperationHistory struct {
	ID              int64           `gorm:"column:id" db:"id" json:"id" form:"id"`                                                         //  主键
	ContractAddress string          `gorm:"column:contract_address" db:"contract_address" json:"contract_address" form:"contract_address"` //  nft合约地址
	ContractType    string          `gorm:"column:contract_type" db:"contract_type" json:"contract_type" form:"contract_type"`             //  nft合约地址类型
	TokenId         string          `gorm:"column:token_id" db:"token_id" json:"token_id" form:"token_id"`                                 //  nft编号的token id
	SellAmount      string          `gorm:"column:sell_amount" db:"sell_amount" json:"sell_amount" form:"sell_amount"`                     //  nft token id数量，0 代表是erc721协议，非0兼容erc1155协议
	Price           decimal.Decimal `gorm:"column:price" db:"price" json:"price" form:"price"`                                             //  售卖价格，支付代币的数量
	PaymentToken    string          `gorm:"column:payment_token" db:"payment_token" json:"payment_token" form:"payment_token"`             //  支付代币的合约地址
	From            string          `gorm:"column:from" db:"from" json:"from" form:"from"`                                                 //  订单from地址
	To              string          `gorm:"column:to" db:"to" json:"to" form:"to"`                                                         //  订单to地址
	OrderHash       string          `gorm:"column:order_hash" db:"order_hash" json:"order_hash" form:"order_hash"`                         //  订单hash(订单id), 后台系统需要关联history表
	EventType       int64           `gorm:"column:event_type" db:"event_type" json:"event_type" form:"event_type"`                         //  交易类型：1: mint 铸造， 2: list 挂单，3:cancel 取消挂单， 4: offer 出价单 ， 5: cancel offer 取消报价，6: matched 撮合成功 7 : transfer 转移, 8:expired 订单过期
	Hash            string          `gorm:"column:hash" db:"hash" json:"hash" form:"hash"`                                                 //  交易hash，上架属于链下，没有hash
	ChainId         int64           `gorm:"column:chain_id" db:"chain_id" json:"chain_id" form:"chain_id"`                                 //  支持公链的id, 默认以太坊主网
	Deleted         int64           `gorm:"column:deleted" db:"deleted" json:"deleted" form:"deleted"`                                     //  逻辑删除,  0:未删除, 1:已删除
	CreatedAt       time.Time       `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`                         //  创建时间
	UpdatedAt       time.Time       `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`                         //  修改时间
}

func (NftOperationHistory) TableName() string {
	return "nft_operation_history"
}

// NftAttribute  nft属性表
type NftAttribute struct {
	ID              int64     `gorm:"column:id" db:"id" json:"id" form:"id"`                                                         //  id
	ChainId         int64     `gorm:"column:chain_id" db:"chain_id" json:"chain_id" form:"chain_id"`                                 //  支持公链的id, 默认以太坊主网
	ContractAddress string    `gorm:"column:contract_address" db:"contract_address" json:"contract_address" form:"contract_address"` //  合约地址
	TokenId         string    `gorm:"column:token_id" db:"token_id" json:"token_id" form:"token_id"`                                 //  nft合约token id
	Name            string    `gorm:"column:name" db:"name" json:"name" form:"name"`                                                 //  属性名称
	Value           string    `gorm:"column:value" db:"value" json:"value" form:"value"`                                             //  属性值
	Rate            float64   `gorm:"column:rate" db:"rate" json:"rate" form:"rate"`                                                 //  该属性占比是多少
	CreatedAt       time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`                         //  创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`                         //  修改时间
}

func (NftAttribute) TableName() string {
	return "nft_attribute"
}
