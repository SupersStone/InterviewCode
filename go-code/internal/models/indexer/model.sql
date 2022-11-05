create table `event_definition` (
    `id` bigint unsigned auto_increment not null not null comment 'id',
    `created_at` timestamp default CURRENT_TIMESTAMP not null comment '保存时间',
    `updated_at` timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `contract_address` varchar(42) default '' not null comment '合约地址',
    `event_signature` varchar(66) default '' not null comment '事件签名',
    `contract_type` varchar(20) default '' not null comment '合约类型',
    `event_name` varchar(30) default '' not null comment '事件方法名',
    `chain_id` varchar(20) default '' not null comment 'chain id',
    `chain_name` varchar(50) default '' not null comment '链名',
    PRIMARY KEY `pk_t_event_definition` (`id`) USING BTREE,
    UNIQUE KEY `uniq_definition` (
        `chain_name`,
        `chain_id`,
        `contract_address`,
        `event_signature`
    ) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '事件定义';

create table `process_log` (
    `id` bigint unsigned auto_increment not null comment 'id',
    `created_at` timestamp default CURRENT_TIMESTAMP not null comment '保存时间',
    `updated_at` timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `event_definition_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '事件定义表id',
    `block_hash` varchar(66) NOT NULL DEFAULT '' COMMENT '区块hash',
    `transaction_hash` varchar(66) NOT NULL DEFAULT '' COMMENT '交易hash',
    `block_number` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '区块链高度',
    `transaction_index` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '交易编号',
    `block_timestamp` bigint unsigned default 0 not null comment '区块时间戳',
    `log_index` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '日志编号',
    `status` tinyint unsigned default null comment '1:成功, 2:失败',
    `log_json` text default null comment 'event json',
    `msg` varchar(300) default '' comment '处理消息',
    PRIMARY KEY `pk_t_process_log` (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '事件处理日志';

CREATE TABLE `fixed_price_matched_event` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '保存时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `event_definition_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '事件定义表id',
    `block_hash` varchar(66) NOT NULL DEFAULT '' COMMENT '区块hash',
    `transaction_hash` varchar(66) NOT NULL DEFAULT '' COMMENT '交易hash',
    `block_number` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '区块链高度',
    `transaction_index` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '交易编号',
    `log_index` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '日志编号',
    `block_timestamp` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '区块时间戳',
    `maker` varchar(42) NOT NULL DEFAULT '' COMMENT '订单maker address',
    `taker` varchar(42) NOT NULL DEFAULT '' COMMENT '订单taker address',
    `order_hash` varchar(66) NOT NULL DEFAULT '' COMMENT 'make订单hash',
    `order_bytes` text COMMENT 'make订单eip712标准订单内容',
    `royalty_recipient` varchar(42) NOT NULL DEFAULT '' COMMENT '版税接收人',
    `royalty_rate` varchar(66) NOT NULL DEFAULT '' COMMENT '合约版税',
    `start_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单开始时间戳',
    `expire_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单过期时间戳',
    `maker_nonce` varchar(66) NOT NULL DEFAULT '' COMMENT 'dex合约用户nonce',
    `taker_get_nft` tinyint(2) NOT NULL DEFAULT '1' COMMENT '是否是taker获取的nft,1-> 是，0-> 否',
    `nft` varchar(42) NOT NULL DEFAULT '' COMMENT 'nft 地址',
    `ft` varchar(42) NOT NULL DEFAULT '' COMMENT 'erc20 合约地址',
    `nft_id` varchar(66) NOT NULL DEFAULT '0' COMMENT 'nft id',
    `nft_amount` varchar(66) DEFAULT NULL COMMENT 'nft数量',
    `ft_amount` varchar(66) DEFAULT NULL COMMENT '付款erc20的数量',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_log` (
        `event_definition_id`,
        `block_hash`,
        `transaction_hash`,
        `log_index`
    ) USING BTREE,
    KEY `idx_tx` (
        `event_definition_id`,
        `block_hash`,
        `transaction_hash`
    ) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '  固定价格订单成交事件';

create table `erc721_transfer_event` (
    `id` bigint unsigned auto_increment not null comment 'id',
    `created_at` timestamp default CURRENT_TIMESTAMP not null comment '保存时间',
    `updated_at` timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `event_definition_id` bigint unsigned default 0 not null comment '事件定义表id',
    `block_hash` varchar(66) default '' not null comment '区块hash',
    `transaction_hash` varchar(66) default '' not null comment '交易hash',
    `block_number` bigint unsigned default 0 not null comment '区块链高度',
    `transaction_index` bigint unsigned default 0 not null comment '交易编号',
    `log_index` bigint unsigned default 0 not null comment '日志编号',
    `block_timestamp` bigint unsigned default 0 not null comment '区块时间戳',
    `from` varchar(42) default '' not null comment 'from address',
    `to` varchar(42) default '' not null comment 'to address',
    `token_id` varchar(66) default '' not null comment 'token id',
    PRIMARY KEY `pk_t_erc721_transfer_event` (`id`) USING BTREE,
    UNIQUE KEY `uniq_log` (
        `event_definition_id`,
        `block_hash`,
        `transaction_hash`,
        `log_index`
    ) USING BTREE,
    KEY `idx_event_seq` (
        `event_definition_id`,
        `block_number`,
        `transaction_index`,
        `log_index`
    ) USING BTREE,
    KEY `idx_token_id` (
        `event_definition_id`,
        `token_id`
    ) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment 'erc721 transfer事件';

create table `token721` (
    `id` bigint unsigned auto_increment not null comment 'id',
    `created_at` timestamp default CURRENT_TIMESTAMP not null comment '保存时间',
    `updated_at` timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `event_definition_id` bigint unsigned default 0 not null comment '事件定义表id',
    `block_number` bigint unsigned default 0 not null comment '区块链高度',
    `transaction_index` bigint unsigned default 0 not null comment '交易编号',
    `log_index` bigint unsigned default 0 not null comment '日志编号',
    `block_timestamp` bigint unsigned default 0 not null comment '区块时间戳',
    `score` decimal(40, 0) default 0 not null comment '用于判断先后',
    `token_id` varchar(66) default '' not null comment 'erc721 token id ',
    `owner` varchar(42) default '' not null comment 'owner address',
    PRIMARY KEY `pk_t_token721`(`id`) USING BTREE,
    UNIQUE KEY `uniq_token_id` (
        `event_definition_id`,
        `token_id`
    ) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment 'erc721 token 所有记录';

create table `nftex_order_canceled_event` (
    `id` bigint unsigned auto_increment not null comment 'id',
    `created_at` timestamp default CURRENT_TIMESTAMP not null comment '保存时间',
    `updated_at` timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `event_definition_id` bigint unsigned default 0 not null comment '事件定义表id',
    `block_hash` varchar(66) default '' not null comment '区块hash',
    `transaction_hash` varchar(66) default '' not null comment '交易hash',
    `block_number` bigint unsigned default 0 not null comment '区块链高度',
    `transaction_index` bigint unsigned default 0 not null comment '交易编号',
    `log_index` bigint unsigned default 0 not null comment '日志编号',
    `block_timestamp` bigint unsigned default 0 not null comment '区块时间戳',
    `maker` varchar(42) default '' not null comment '订单maker address',
    `order_hash` varchar(66) default '' not null comment '订单hash',
    UNIQUE KEY `uniq_log` (
        `event_definition_id`,
        `block_hash`,
        `transaction_hash`,
        `log_index`
    ) USING BTREE,
    PRIMARY KEY `pk_t_order_cancelled_event` (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '  取消订单事件';

create table `nftex_all_order_canceled_event` (
    `id` bigint unsigned auto_increment not null comment 'id',
    `created_at` timestamp default CURRENT_TIMESTAMP not null comment '保存时间',
    `updated_at` timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `event_definition_id` bigint unsigned default 0 not null comment '事件定义表id',
    `block_hash` varchar(66) default '' not null comment '区块hash',
    `transaction_hash` varchar(66) default '' not null comment '交易hash',
    `block_number` bigint unsigned default 0 not null comment '区块链高度',
    `transaction_index` bigint unsigned default 0 not null comment '交易编号',
    `log_index` bigint unsigned default 0 not null comment '日志编号',
    `block_timestamp` bigint unsigned default 0 not null comment '区块时间戳',
    `maker` varchar(42) default '' not null comment '订单maker address',
    `nonce` varchar(66) default '' not null comment 'dex合约用户nonce',
    UNIQUE KEY `uniq_log` (
        `event_definition_id`,
        `block_hash`,
        `transaction_hash`,
        `log_index`
    ) USING BTREE,
    PRIMARY KEY `pk_t_all_order_cancelled_event` (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '  取消全部订单事件';

create table `token_worship_history` (
    `id` bigint unsigned auto_increment not null comment 'id',
    `created_at` timestamp default CURRENT_TIMESTAMP not null comment '保存时间',
    `updated_at` timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `event_definition_id` bigint unsigned default 0 not null comment '事件定义表id',
    `block_hash` varchar(66) default '' not null comment '区块hash',
    `transaction_hash` varchar(66) default '' not null comment '交易hash',
    `block_number` bigint unsigned default 0 not null comment '区块链高度',
    `transaction_index` bigint unsigned default 0 not null comment '交易编号',
    `log_index` bigint unsigned default 0 not null comment '日志编号',
    `block_timestamp` bigint unsigned default 0 not null comment '区块时间戳',
    `release_timestamp` bigint unsigned default 0 not null comment '释放时间戳',
    `votary` varchar(66) default '' not null comment 'votary',
    `redeemer` varchar(66) default '' not null comment 'owner address',
    `type` tinyint(4) default 0 not null comment 'type 1 offer 2 redeeme',
    `token_id` varchar(66) default '' not null comment 'erc721 token id ',
    `nft_address` varchar(66) default '' not null comment 'nft_address',
    UNIQUE KEY `uniq_log` (
        `event_definition_id`,
        `block_hash`,
        `transaction_hash`,
        `log_index`
    ) USING BTREE,
    PRIMARY KEY `pk_t_nftex_all_order_canceled_event` (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment 'worship NFT 供奉历史记录';

create table `token_worship_status` (
    `id` bigint unsigned auto_increment not null comment 'id',
    `chain_id` varchar(20) default '0' not null comment 'chain id',
    `contract_address` varchar(42) default '' not null comment '合约地址',
    `created_at` timestamp default CURRENT_TIMESTAMP not null comment '保存时间',
    `updated_at` timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `start_time` bigint unsigned default 0 not null comment '开始时间',
    `end_time` bigint unsigned default 0 not null comment '结束时间',
    `token_id` varchar(66) default '' not null comment 'erc721 token id ',
    `worship` tinyint default 0 not null comment '区块时间戳',
    `score` decimal(40, 0) default 0 not null comment '用于判断先后',
    UNIQUE KEY `uniq_token_id` (`token_id`) USING BTREE,
    PRIMARY KEY `pk_t_token_worship_status` (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment 'worship NFT 供奉状态';

INSERT INTO
    `event_definition` (
        `contract_address`,
        `event_signature`,
        `contract_type`,
        `event_name`,
        `chain_id`,
        `chain_name`
    )
VALUES
    (
        '0xdFC304EC128C72066a4b329584EF582C9FfD7a47',
        '0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985',
        'dao-dex',
        'OrderCancelled',
        '80001',
        'polygon'
    ),
    (
        '0xdFC304EC128C72066a4b329584EF582C9FfD7a47',
        '0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985',
        'dao-dex',
        'OrderCancelled',
        '80001',
        'polygon'
    ),
;