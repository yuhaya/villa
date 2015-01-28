/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50621
 Source Host           : localhost
 Source Database       : villa

 Target Server Type    : MySQL
 Target Server Version : 50621
 File Encoding         : utf-8

 Date: 01/27/2015 09:01:25 AM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `villa_admin`
-- ----------------------------
DROP TABLE IF EXISTS `villa_admin`;
CREATE TABLE `villa_admin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `pwd` varchar(40) NOT NULL,
  `date` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `villa_admin`
-- ----------------------------
BEGIN;
INSERT INTO `villa_admin` VALUES ('1', 'yuhaya', '6116afedcb0bc31083935c1c262ff4c9', '2015-01-26 23:27:36');
COMMIT;

-- ----------------------------
--  Table structure for `villa_currency`
-- ----------------------------
DROP TABLE IF EXISTS `villa_currency`;
CREATE TABLE `villa_currency` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(12) NOT NULL,
  `exchange_rate` decimal(7,4) NOT NULL,
  `update_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_destination`
-- ----------------------------
DROP TABLE IF EXISTS `villa_destination`;
CREATE TABLE `villa_destination` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `currency_id` int(10) unsigned NOT NULL,
  `name` varchar(30) DEFAULT NULL,
  `name_en` varchar(60) DEFAULT NULL,
  `level` tinyint(3) unsigned NOT NULL,
  `left` smallint(5) unsigned NOT NULL,
  `right` smallint(5) unsigned NOT NULL,
  `sort_num` smallint(5) unsigned NOT NULL,
  `memo` varchar(256) NOT NULL,
  `tag` tinyint(3) unsigned NOT NULL,
  `status` tinyint(3) unsigned NOT NULL,
  `show` tinyint(3) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `villa_destination_currency_id` (`currency_id`),
  KEY `villa_destination_level` (`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_destination_img`
-- ----------------------------
DROP TABLE IF EXISTS `villa_destination_img`;
CREATE TABLE `villa_destination_img` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `destination_id` int(10) unsigned NOT NULL,
  `img` varchar(255) NOT NULL,
  `type` tinyint(3) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `villa_destination_img_destination_id` (`destination_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_inquire`
-- ----------------------------
DROP TABLE IF EXISTS `villa_inquire`;
CREATE TABLE `villa_inquire` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` int(10) unsigned NOT NULL,
  `destination_id` int(10) unsigned NOT NULL,
  `uid` int(10) unsigned NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `status` tinyint(3) unsigned NOT NULL,
  `user_ip` varchar(20) DEFAULT NULL,
  `add_date` datetime NOT NULL,
  `update_date` datetime NOT NULL,
  `user_message` varchar(250) DEFAULT NULL,
  `memo` varchar(512) DEFAULT NULL,
  `lock` tinyint(3) unsigned NOT NULL,
  `source` tinyint(3) unsigned NOT NULL,
  `adults` tinyint(3) unsigned NOT NULL,
  `kids` tinyint(3) unsigned NOT NULL,
  `bedroom` tinyint(3) unsigned NOT NULL,
  `kid_add_beds` tinyint(3) unsigned NOT NULL,
  `adults_add_beds` tinyint(3) unsigned NOT NULL,
  `invoice` tinyint(3) unsigned NOT NULL,
  `coupon` tinyint(3) unsigned NOT NULL,
  `pay_date` date DEFAULT NULL,
  `tax_formula` tinyint(3) unsigned NOT NULL,
  `deprecated_reson` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `villa_inquire_product_id` (`product_id`),
  KEY `villa_inquire_destination_id` (`destination_id`),
  KEY `villa_inquire_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_inquire_currency`
-- ----------------------------
DROP TABLE IF EXISTS `villa_inquire_currency`;
CREATE TABLE `villa_inquire_currency` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `inquire_id` int(10) unsigned NOT NULL,
  `name` varchar(12) NOT NULL,
  `exchange_rate` decimal(7,4) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `villa_inquire_currency_inquire_id` (`inquire_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_inquire_extra`
-- ----------------------------
DROP TABLE IF EXISTS `villa_inquire_extra`;
CREATE TABLE `villa_inquire_extra` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `inquire_id` int(10) unsigned NOT NULL,
  `card_id` varchar(30) DEFAULT NULL,
  `flight_number` varchar(30) DEFAULT NULL,
  `flight_start_time` datetime DEFAULT NULL,
  `flight_end_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `villa_inquire_extra_inquire_id` (`inquire_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_inquire_price_cost`
-- ----------------------------
DROP TABLE IF EXISTS `villa_inquire_price_cost`;
CREATE TABLE `villa_inquire_price_cost` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `inquire_id` int(10) unsigned NOT NULL,
  `invoice_price` decimal(7,4) NOT NULL,
  `coupon_price` decimal(7,4) NOT NULL,
  `flights_in` decimal(7,4) NOT NULL,
  `flights_out` decimal(7,4) NOT NULL,
  `insurance_in` decimal(7,4) NOT NULL,
  `insurance_out` decimal(7,4) NOT NULL,
  `poundage_pay` decimal(7,4) NOT NULL,
  `poundage_remit` decimal(7,4) NOT NULL,
  `adults_bed_rate` decimal(7,4) NOT NULL,
  `kids_bed_rate` decimal(7,4) NOT NULL,
  `rooms_rate_in` decimal(7,4) NOT NULL,
  `rooms_rate_out` decimal(7,4) NOT NULL,
  `airport_transfes_rate` decimal(7,4) NOT NULL,
  `location_rate` decimal(7,4) NOT NULL,
  `service_rate` decimal(7,4) NOT NULL,
  `breakfasts_rate` decimal(7,4) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `villa_inquire_price_cost_inquire_id` (`inquire_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_inquire_villa`
-- ----------------------------
DROP TABLE IF EXISTS `villa_inquire_villa`;
CREATE TABLE `villa_inquire_villa` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `inquire_id` int(10) unsigned NOT NULL,
  `user_passport` varchar(30) DEFAULT NULL,
  `card_id` varchar(30) DEFAULT NULL,
  `basic_facility` varchar(255) DEFAULT NULL,
  `integrated_facility` varchar(255) DEFAULT NULL,
  `action_facility` varchar(255) DEFAULT NULL,
  `nearby_facility` varchar(255) DEFAULT NULL,
  `tags` varchar(255) DEFAULT NULL,
  `service_free` varchar(255) DEFAULT NULL,
  `location_tax_radio` decimal(7,4) NOT NULL,
  `service_tax_radio` decimal(7,4) NOT NULL,
  `cancellation_clause` varchar(255) DEFAULT NULL,
  `kid_clause` varchar(255) DEFAULT NULL,
  `adult_bed_rate` decimal(7,4) NOT NULL,
  `kid_bed_rate` decimal(7,4) NOT NULL,
  `breakfast_rate` decimal(7,4) NOT NULL,
  `airport_transfe_rate` decimal(7,4) NOT NULL,
  `commission` decimal(7,4) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `villa_inquire_villa_inquire_id` (`inquire_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_partner`
-- ----------------------------
DROP TABLE IF EXISTS `villa_partner`;
CREATE TABLE `villa_partner` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `destination_id` int(10) unsigned NOT NULL,
  `manager_contact` varchar(20) DEFAULT NULL,
  `manager_telephone` varchar(20) DEFAULT NULL,
  `manager_email` varchar(20) DEFAULT NULL,
  `reservation_contact` varchar(20) DEFAULT NULL,
  `reservation_telephone` varchar(20) DEFAULT NULL,
  `reservation_email` varchar(20) DEFAULT NULL,
  `created_date` datetime NOT NULL,
  `contract_start_date` date DEFAULT NULL,
  `contract_end_date` date DEFAULT NULL,
  `commission` decimal(7,4) NOT NULL,
  `membership_group` varchar(20) NOT NULL,
  `state` tinyint(3) unsigned NOT NULL,
  `memo` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `website` varchar(100) DEFAULT NULL,
  `availablity_link` varchar(100) DEFAULT NULL,
  `availablity_username` varchar(30) DEFAULT NULL,
  `availablity_password` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_product`
-- ----------------------------
DROP TABLE IF EXISTS `villa_product`;
CREATE TABLE `villa_product` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `destination_id` int(10) unsigned NOT NULL,
  `partner_id` int(10) unsigned DEFAULT NULL,
  `currency_id` tinyint(3) unsigned NOT NULL,
  `name` varchar(30) NOT NULL,
  `name_en` varchar(50) DEFAULT NULL,
  `memo` varchar(50) DEFAULT NULL,
  `location_info` varchar(50) DEFAULT NULL,
  `shower_room` tinyint(3) unsigned NOT NULL,
  `short_description` varchar(100) DEFAULT NULL,
  `description` varchar(500) DEFAULT NULL,
  `longitude` decimal(7,4) NOT NULL,
  `latitude` decimal(7,4) NOT NULL,
  `add_date` datetime NOT NULL,
  `valid_from_date` date NOT NULL,
  `valid_to_date` date NOT NULL,
  `state` tinyint(3) unsigned NOT NULL,
  `checkin_time` varchar(255) DEFAULT NULL,
  `check_out_time` varchar(255) DEFAULT NULL,
  `adult_bed_rate` decimal(7,4) NOT NULL,
  `kid_bed_rate` decimal(7,4) NOT NULL,
  `kid_add_beds_max` tinyint(3) unsigned NOT NULL,
  `adults_add_beds_max` tinyint(3) unsigned NOT NULL,
  `cancellation_clause` varchar(255) DEFAULT NULL,
  `kid_clause` varchar(255) DEFAULT NULL,
  `location_tax_radio` decimal(7,4) NOT NULL,
  `service_tax_radio` decimal(7,4) NOT NULL,
  `airport_transfe_rate` decimal(7,4) NOT NULL,
  `commission` decimal(7,4) NOT NULL,
  `deposit_ratio` decimal(7,4) NOT NULL,
  `tax_formula` tinyint(3) unsigned NOT NULL,
  `website` varchar(100) DEFAULT NULL,
  `reservation_contact` varchar(20) DEFAULT NULL,
  `reservation_telephone` varchar(20) DEFAULT NULL,
  `reservation_email` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_product_img`
-- ----------------------------
DROP TABLE IF EXISTS `villa_product_img`;
CREATE TABLE `villa_product_img` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` int(10) unsigned NOT NULL,
  `name` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_promotion`
-- ----------------------------
DROP TABLE IF EXISTS `villa_promotion`;
CREATE TABLE `villa_promotion` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_rate`
-- ----------------------------
DROP TABLE IF EXISTS `villa_rate`;
CREATE TABLE `villa_rate` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` int(10) unsigned NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `bedroom` tinyint(3) unsigned NOT NULL,
  `mini_stay` tinyint(3) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_tag`
-- ----------------------------
DROP TABLE IF EXISTS `villa_tag`;
CREATE TABLE `villa_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `type` tinyint(3) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_tag_product`
-- ----------------------------
DROP TABLE IF EXISTS `villa_tag_product`;
CREATE TABLE `villa_tag_product` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` int(10) unsigned NOT NULL,
  `tag_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `villa_tag_product_product_id` (`product_id`),
  KEY `villa_tag_product_tag_id` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `villa_user`
-- ----------------------------
DROP TABLE IF EXISTS `villa_user`;
CREATE TABLE `villa_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `password` varchar(40) NOT NULL,
  `email` varchar(40) DEFAULT NULL,
  `phone` varchar(40) DEFAULT NULL,
  `add_date` datetime NOT NULL,
  `user_passport` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
