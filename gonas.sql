/*
 Navicat Premium Data Transfer

 Source Server         : gonas
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : localhost:3307
 Source Schema         : gonas

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 01/06/2020 16:09:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(3) unsigned NOT NULL DEFAULT '0',
  `name` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
  `describe` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for devices
-- ----------------------------
DROP TABLE IF EXISTS `devices`;
CREATE TABLE `devices` (
  `id` int(1) unsigned NOT NULL AUTO_INCREMENT,
  `display` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
  `memo` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '备注',
  `mac` varchar(255) NOT NULL DEFAULT '' COMMENT 'mac地址',
  `last_ip` int(1) unsigned NOT NULL DEFAULT '0' COMMENT 'ip地址',
  `type` tinyint(1) NOT NULL COMMENT '设备类型',
  `created_at` timestamp(1) NOT NULL DEFAULT CURRENT_TIMESTAMP(1),
  `updated_at` timestamp(1) NOT NULL DEFAULT CURRENT_TIMESTAMP(1),
  `deleted_at` timestamp(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for directories
-- ----------------------------
DROP TABLE IF EXISTS `directories`;
CREATE TABLE `directories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `display` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of directories
-- ----------------------------
BEGIN;
INSERT INTO `directories` VALUES (1, 1, '', 0, '2020-05-18 13:35:53', '2020-05-18 13:35:53', NULL);
COMMIT;

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `display` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
  `hash` varchar(255) NOT NULL DEFAULT '',
  `size` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '单位Byte',
  `type` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `original_name` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '原始文件名',
  `device_id` tinyint(4) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of files
-- ----------------------------
BEGIN;
INSERT INTO `files` VALUES (1, '文件1', 'xxx', 10, 0, '', 1, '2020-05-08 08:58:46', '2020-05-08 08:58:46', NULL);
COMMIT;

-- ----------------------------
-- Table structure for log_users
-- ----------------------------
DROP TABLE IF EXISTS `log_users`;
CREATE TABLE `log_users` (
  `id` int(1) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `resource_id` int(10) unsigned NOT NULL DEFAULT '0',
  `type` tinyint(1) NOT NULL COMMENT '日志类型',
  `extra` varchar(1000) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
  `created_at` timestamp(1) NOT NULL DEFAULT CURRENT_TIMESTAMP(1),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for user_category
-- ----------------------------
DROP TABLE IF EXISTS `user_category`;
CREATE TABLE `user_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `category_id` int(10) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for user_files
-- ----------------------------
DROP TABLE IF EXISTS `user_files`;
CREATE TABLE `user_files` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `file_id` int(10) unsigned NOT NULL DEFAULT '0',
  `directory_id` int(10) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQUE_USER_FILE` (`user_id`,`file_id`),
  KEY `USER_ID_INDEX` (`user_id`) USING HASH,
  KEY `USER_DIRECTORY_INDEX` (`directory_id`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of user_files
-- ----------------------------
BEGIN;
INSERT INTO `user_files` VALUES (1, 1, 1, 1, '2020-05-18 12:25:37', '2020-05-18 12:25:37', NULL);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(1) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户表',
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '登录名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `device_id` tinyint(1) NOT NULL DEFAULT '0',
  `login_count` int(1) unsigned NOT NULL DEFAULT '0' COMMENT '登录次数',
  `last_login_at` timestamp(1) NOT NULL DEFAULT CURRENT_TIMESTAMP(1) COMMENT '上次登录时间',
  `last_login_ip` int(1) unsigned NOT NULL DEFAULT '0' COMMENT '上次登录IP',
  `created_at` timestamp(1) NOT NULL DEFAULT CURRENT_TIMESTAMP(1),
  `updated_at` timestamp(1) NOT NULL DEFAULT CURRENT_TIMESTAMP(1),
  `deleted_at` timestamp(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'meshren', '2d38c40a16397893cbcba76d174afbb5', 0, 0, '2020-05-16 09:25:46.6', 0, '2020-05-16 09:25:46.6', '2020-05-16 09:25:46.6', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
