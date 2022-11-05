CREATE TABLE `block_scan_contract` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '合约名称',
  `address` varchar(42) NOT NULL DEFAULT '' COMMENT '合约地址',
  `chain` varchar(30) NOT NULL DEFAULT '' COMMENT '链名',
  `chain_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '链ID',
  `start_height` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '开始扫块高度',
  `end_height` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '新增合约的时候，主线程开始扫块高度',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否扫描 0不扫描 1扫描',
  `abi` text COMMENT '合约abi',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT = '合约地址表';

CREATE TABLE `block_scan_height` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `task_name` varchar(30) NOT NULL DEFAULT '' COMMENT '扫块任务名称',
  `chain` varchar(10) NOT NULL DEFAULT '' COMMENT '链名',
  `height` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '当前高度',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_task_chain` (`task_name`, `chain`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT = '任务扫块高度表';