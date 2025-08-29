CREATE TABLE `airdrop` (
                           `id` int(1) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                           `to_address` varchar(42) NOT NULL COMMENT '被空投的地址',
                           `hash` varchar(66) DEFAULT NULL COMMENT 'tx hash',
                           `airdrop_amount` varchar(255) NOT NULL COMMENT '空投数量',
                           `tx_status` int(1) unsigned zerofill NOT NULL COMMENT 'tx status 0 失败, 1 成功',
                           `airdrop_count` int(1) unsigned zerofill NOT NULL COMMENT '空投次数',
                           `airdrop_status` int(1) unsigned zerofill NOT NULL COMMENT '0 失败, 1 等待验证, 2  验证失败, 3 验证成功, 4. 手动操作',
                           `airdrop_time` bigint(1) unsigned zerofill DEFAULT NULL COMMENT '空投时间戳',
                           PRIMARY KEY (`id`,`to_address`) USING BTREE,
                           UNIQUE KEY `u_to_address` (`to_address`) USING BTREE COMMENT '地址唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=latin1