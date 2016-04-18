DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `member_id` varchar(255) DEFAULT NULL,
  `service` varchar(255) DEFAULT NULL,
  `app` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `service_member_id_idx` (`service`, `member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `word_chain`;
CREATE TABLE `word_chain` (
  `account_id` int(10) DEFAULT NULL,
  `last_word` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
