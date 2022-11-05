






CREATE TABLE `nft_attribute` (
     `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
     `chain_id` int NOT NULL DEFAULT '1' COMMENT '支持公链的ID, 默认以太坊主网',
     `contract_address` varchar(255)  NOT NULL DEFAULT '0' COMMENT '合约地址',
     `token_id` varchar(500) NOT NULL DEFAULT '0'  COMMENT 'NFT合约token ID',
     `name` varchar(100)  DEFAULT NULL COMMENT '属性名称',
     `value` varchar(100)  DEFAULT NULL COMMENT '属性值',
     `rate` decimal(5,2) DEFAULT NULL COMMENT '该属性占比是多少',
     `deleted` bit(1) DEFAULT b'0' COMMENT '0 ：未删除，1：已删除',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
     PRIMARY KEY (`id`),
     KEY `token_id` (`token_id`) USING BTREE COMMENT '查询条件索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='nft属性表';




CREATE TABLE `nft_base_info` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
      `contract_address` varchar(255) NOT NULL  DEFAULT ''  COMMENT 'nft合约地址',
      `contract_type` varchar(255) NOT NULL  DEFAULT ''  COMMENT 'nft合约地址类型，721， 1155等协议类型',
      `contract_creator` varchar(255) NOT NULL DEFAULT ''  COMMENT 'nft合约地址的创建者',
      `token_id` varchar(500) NOT NULL DEFAULT '0'  COMMENT 'NFT合约token ID',
      `token_amount` varchar(500) NOT NULL DEFAULT '0'  COMMENT 'nft token id数量，0 代表是erc721协议，非0兼容erc1155协议',
      `name` varchar(255) NOT NULL  DEFAULT ''  COMMENT 'nft 名字，合集名字#token_id组合形成',
      `collection_id` bigint DEFAULT NULL COMMENT '当前nft token id 是属于那个合集',
      `metadata_url` varchar(255) DEFAULT NULL COMMENT 'ipfs的url地址',
      `image_url` varchar(255) DEFAULT NULL COMMENT 'aws 的 url 图片存储地址',
      `chain_id` int NOT NULL DEFAULT "1" COMMENT '支持公链的ID, 默认以太坊主网',
      `creator_rate_fee` decimal(5,4) DEFAULT NULL COMMENT '创建者手续费比例',
      `nft_verify` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'nft合约地址是否验证开源，默认为0，未开源， 1代表开源',
      `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除,  0:未删除, 1:已删除',
      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
      PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='nft基础信息资产表';






CREATE TABLE `nft_operation_history` (
     `id` bigint(50) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
     `contract_address` varchar(255) NOT NULL COMMENT 'nft合约地址',
     `contract_type` varchar(255) NOT NULL COMMENT 'nft合约地址类型',
     `token_id` varchar(500) NOT NULL DEFAULT '0' COMMENT 'nft编号的token id',
     `sell_amount` varchar(500) NOT NULL DEFAULT '0'  COMMENT 'nft token id数量，0 代表是erc721协议，非0兼容erc1155协议',
     `price` decimal(30,4) DEFAULT NULL COMMENT '售卖价格，支付代币的数量',
     `payment_token` varchar(255) DEFAULT NULL COMMENT '支付代币的合约地址',
     `from` varchar(255) DEFAULT NULL COMMENT '订单from地址',
     `to` varchar(255) DEFAULT NULL COMMENT '订单to地址',
     `order_hash` varchar(500) NOT NULL COMMENT '订单hash(订单id), 后台系统需要关联history表',
     `event_type` int(11) NOT NULL COMMENT '交易类型：1: Mint 铸造， 2: List 挂单，3:Cancel 取消挂单， 4: offer 出价单 ， 5: Cancel offer 取消报价，6: Matched 撮合成功 7 : Transfer 转移, 8:Expired 订单过期',
     `hash` varchar(255) DEFAULT NULL COMMENT '交易hash，上架属于链下，没有hash',
     `chain_id` int NOT NULL DEFAULT "1" COMMENT '支持公链的ID, 默认以太坊主网',
     `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除,  0:未删除, 1:已删除',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='nft的订单操作历史';




CREATE TABLE `nft_order` (
     `id` bigint(50) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
     `nft` varchar(66) NOT NULL DEFAULT '' COMMENT 'nft合约地址',
     `nft_id` varchar(256) NOT NULL DEFAULT '0' COMMENT 'nft编号的token id',
     `nft_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT 'nft token id数量，0 代表是erc721协议，非0兼容erc1155协议',
     `ft` varchar(66) NOT NULL DEFAULT '0' COMMENT '支付代币的合约地址',
     `ft_amount` decimal(40,0) NOT NULL DEFAULT "0.0" COMMENT '售卖价格，支付代币的数量',
     `maker` varchar(66) NOT NULL DEFAULT '' COMMENT '订单创建者地址，maker地址或者是taker地址，也是签名地址',
     `royalty_recipient` varchar(255) NOT NULL DEFAULT '' COMMENT 'nft实际接受者，针对于oferrs订单来说，可定制买单接受者地址',
     `service_fee` decimal(5,4) NOT NULL DEFAULT '0' COMMENT '平台手续费占比,前端从合约中获取该参数',
     `royalty_rate` decimal(5,4) NOT NULL DEFAULT '0' COMMENT '创作者版权费用，创建合集的时候确定版权占比',
     `start_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '订单开始时间',
     `expire_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '订单过期时间',
     `maker_nonce` bigint NOT NULL DEFAULT '0' COMMENT '订单nonce，默认是0，如果是点击批量取消订单后，默认的nonce加1',
     `taker_get_nft` tinyint(2) NOT NULL COMMENT '订单类型，0=taker卖方挂市价单，1=maker买方挂限价单',
     `sig` varchar(500) NOT NULL DEFAULT '' COMMENT '签名信息',
     `order_hash` varchar(66) NOT NULL DEFAULT '' COMMENT '订单hash(订单id), 后台系统需要关联history表',
     `chain_id` int NOT NULL DEFAULT "1" COMMENT '支持公链的ID, 默认以太坊主网',
     `contract_type` varchar(255) NOT NULL DEFAULT '0' COMMENT 'nft合约地址类型',
     `collection_id` bigint(20) DEFAULT NULL DEFAULT '0' COMMENT '当前nft token id 是属于那个合集',
     `status` tinyint(4) DEFAULT NULL COMMENT '当前订单状态，1=Listing(挂单)，2=Cancel(取消挂单)，3=Match(撮合成功)，4=Invalid(订单失效)',
     `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除,  0:未删除, 1:已删除',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='nft的订单表（同步合约中定义的字段）';


