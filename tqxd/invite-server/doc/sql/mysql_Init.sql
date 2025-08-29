CREATE TABLE `Basic_Type` (
    `Base_ID`            BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `Base_Type`          INT NOT NULL,
    `Base_Name`          VARCHAR(30)  NOT NULL,
    INDEX `idx_basic_asset` (`Base_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `KafkaTopic_Offset` (
    
    `Topic_Name`         VARCHAR(128) NOT NULL,
    `Topic_Offset`       BIGINT ,
    INDEX `idx_Topic_Name` (`Topic_Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO KafkaTopic_Offset ( Topic_Name, Topic_Offset)VALUES( "balances", -3);
INSERT INTO KafkaTopic_Offset ( Topic_Name, Topic_Offset)VALUES( "balances_as", -3);
INSERT INTO KafkaTopic_Offset ( Topic_Name, Topic_Offset)VALUES( "deals", -3);
INSERT INTO KafkaTopic_Offset ( Topic_Name, Topic_Offset)VALUES( "orders", -3);
INSERT INTO KafkaTopic_Offset ( Topic_Name, Topic_Offset)VALUES( "ckAttribute", -3);
INSERT INTO KafkaTopic_Offset ( Topic_Name, Topic_Offset)VALUES( "ckAction", -3);
INSERT INTO KafkaTopic_Offset ( Topic_Name, Topic_Offset)VALUES( "ckBusiness", -3);