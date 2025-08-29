CREATE TABLE `t_airdrop` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `address` varchar(128) DEFAULT '' COMMENT '地址',
    `hash` varchar(128) DEFAULT '0' COMMENT '交易哈希',
    `amount` decimal(50,0) DEFAULT NULL COMMENT '数量',
    `status` int(11) default '0' COMMENT '0 默认值, 1 成功 2 失败',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='空投';