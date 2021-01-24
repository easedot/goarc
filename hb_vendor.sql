CREATE DATABASE  IF NOT EXISTS `huobi_vendors` ;
USE `huobi_vendors`;


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255)  NOT NULL,
  `type` smallint(6) NOT NULL DEFAULT 0, /*1:供应商 2:管理员 3:运营人员*/
  `state` smallint(6) NOT NULL, /*1:正常 2:禁用*/
  `email` varchar(80)  NOT NULL DEFAULT '',
  `password` varchar(100)  NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

LOCK TABLES `user` WRITE;
INSERT INTO `user` VALUES (1,'user1',0,1,'user1@email.com','user1');
INSERT INTO `user` VALUES (2,'user1',0,2,'user2@email.com','user2');
INSERT INTO `user` VALUES (3,'user2',1,1,'user3@email.com','user3');
INSERT INTO `user` VALUES (4,'user3',2,1,'user4@email.com','user4');
UNLOCK TABLES;

DROP TABLE IF EXISTS `vendor`;
CREATE TABLE `vendor` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `state` smallint(6) NOT NULL,     /*1:待审核 2:审核中 3:通过 4:拒绝 5:补充*/
  `type` smallint(6)  NOT NULL,     /*1:普通 2:临时*/
  `name` varchar(200)  NOT NULL,
  `address` varchar(800)  NOT NULL  DEFAULT '',
  `phone` varchar(30)  NOT NULL  DEFAULT '',
  `country_code` varchar(100)  NOT NULL DEFAULT 'CN',/*CN  https://countrycode.org/ */
  `web_link` varchar(300)  NOT NULL DEFAULT '',
  `registered_address` varchar(800)  NOT NULL  DEFAULT '',
  `registered_capital` int(11) NOT NULL DEFAULT 0,
  `registered_no` varchar(100)  NOT NULL  DEFAULT '',
  `registered_date` varchar(30)  NOT NULL  DEFAULT '',
  `registered_type` smallint(6) NOT NULL  DEFAULT 0, /*1私企、2国企、3事业、4外资、5其他 */
  `tax_no` varchar(100)  NOT NULL DEFAULT '',/**/
  `employee_count` int(11) NOT NULL DEFAULT 0,/**/
  `market_staff_count` int(11) NOT NULL  DEFAULT 0,/**/
  `technical_staff_count` int(11) NOT NULL  DEFAULT 0,/**/
  `bank_name` varchar(200)  NOT NULL DEFAULT '',/**/
  `bank_account` varchar(50)  NOT NULL DEFAULT '',/**/
  `referrer` varchar(50)  NULL DEFAULT '',/**/
  `referrer_reason` varchar(200)  NULL DEFAULT '',/**/
  `success_case_documents` varchar(200)  NULL DEFAULT '',/**/
  `main_product` varchar(200)  NULL DEFAULT '',/**/
  `channel_level` varchar(50)  NULL DEFAULT '',/**/
  `is_all_country` tinyint(1)  NOT NULL DEFAULT 0,/**/
  `boss_name` varchar(50)  NOT NULL DEFAULT '', /*负责人*/
  `boss_email` varchar(80)  NOT NULL DEFAULT '',/*负责人邮件*/
  `boss_phone` varchar(30)  NOT NULL DEFAULT '',/*负责人电话*/
  `boss_tel` varchar(30)  NOT NULL DEFAULT '',/*负责人固话*/
  `contact_name` varchar(50)  NOT NULL DEFAULT '',/**/
  `contact_email` varchar(80)  NOT NULL DEFAULT '',/**/
  `contact_phone` varchar(30)  NOT NULL DEFAULT '',/**/
  `contact_tel` varchar(30)  NOT NULL DEFAULT '',/**/
  `qualification_documents` varchar(200)  NULL DEFAULT '',/**/
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`country_code`,`name`),
  UNIQUE KEY `registered_info` (`country_code`,`registered_no`),
  UNIQUE KEY `tax_info` (`country_code`,`tax_no`),
  KEY `vendor_state` (`state`),
  KEY `vendor_phone` (`phone`),
  KEY `vendor_boss_phone` (`boss_phone`),
  KEY `vendor_boss_email` (`boss_email`),
  KEY `vendor_contact_phone` (`contact_phone`),
  KEY `vendor_contact_email` (`contact_email`),
  KEY `vendor_registered_no` (`registered_no`),
  KEY `vendor_tax_no` (`tax_no`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

LOCK TABLES `vendor` WRITE;
INSERT INTO `vendor` VALUES (1,1,2,1,'test_vendor');
UNLOCK TABLES;


