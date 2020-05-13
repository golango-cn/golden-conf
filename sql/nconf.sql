/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : 127.0.0.1:3306
 Source Schema         : nconf

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 13/05/2020 20:38:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `admin_id` bigint NOT NULL AUTO_INCREMENT,
  `admin_name` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`admin_id`),
  KEY `idx_admin_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of admin
-- ----------------------------
BEGIN;
INSERT INTO `admin` VALUES (10, 'admin', '2020-04-28 17:15:41', '2020-04-28 17:15:46', NULL);
INSERT INTO `admin` VALUES (20, 'admin2', '2020-04-28 17:28:40', '2020-04-28 17:28:43', NULL);
INSERT INTO `admin` VALUES (30, 'admin3', '2020-04-28 17:28:50', '2020-04-28 17:28:54', NULL);
COMMIT;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `app_id` bigint NOT NULL AUTO_INCREMENT,
  `app_code` varchar(255) DEFAULT NULL,
  `app_name` varchar(255) DEFAULT NULL,
  `app_desc` varchar(255) DEFAULT NULL,
  `admin_id` bigint DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`app_id`),
  KEY `idx_app_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of app
-- ----------------------------
BEGIN;
INSERT INTO `app` VALUES (1, 'member-core', '会员组核心项目', '会员组核心项目', 30, '2020-04-28 17:21:28', '2020-04-28 17:21:28', NULL);
INSERT INTO `app` VALUES (2, 'member-core', '会员组核心项目', '会员组核心项目', 20, '2020-04-28 17:28:00', '2020-04-28 17:28:00', NULL);
INSERT INTO `app` VALUES (3, 'member-core', '会员组核心项目', '会员组核心项目', 10, '2020-04-28 17:28:10', '2020-04-28 17:28:10', NULL);
INSERT INTO `app` VALUES (4, '', '', '', 0, '2020-04-28 21:06:40', '2020-04-28 21:06:40', NULL);
INSERT INTO `app` VALUES (5, '', '', '', 0, '2020-04-29 17:16:28', '2020-04-29 17:16:28', NULL);
INSERT INTO `app` VALUES (6, '', '', '', 0, '2020-04-30 09:28:35', '2020-04-30 09:28:35', NULL);
INSERT INTO `app` VALUES (7, 'member-core333', '会员组核心项目2233', '会员组核心项目22', 12, '2020-04-30 09:35:53', '2020-04-30 09:35:53', NULL);
COMMIT;

-- ----------------------------
-- Table structure for app_config
-- ----------------------------
DROP TABLE IF EXISTS `app_config`;
CREATE TABLE `app_config` (
  `app_config_id` bigint NOT NULL AUTO_INCREMENT,
  `app_id` bigint DEFAULT NULL,
  `app_code` varchar(255) DEFAULT NULL,
  `environment_name` varchar(255) DEFAULT NULL,
  `key` varchar(255) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  `status` bigint DEFAULT NULL,
  `memo` varchar(255) DEFAULT NULL,
  `admin_id` bigint DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`app_config_id`),
  KEY `idx_app_config_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of app_config
-- ----------------------------
BEGIN;
INSERT INTO `app_config` VALUES (1, 1, 'member-core', 'dev', 'sql_driver', 'mysql', 0, '', 1, '2020-04-28 22:08:44', '2020-04-30 15:51:42', NULL);
INSERT INTO `app_config` VALUES (2, 1, 'member-core', 'dev', 'sql_driver1', '123321', 0, '测试备注', 20, '2020-04-28 22:09:01', '2020-04-30 14:08:10', NULL);
INSERT INTO `app_config` VALUES (3, 0, '', '', '', '', 0, '', 0, '2020-04-30 14:27:55', '2020-04-30 14:27:55', NULL);
INSERT INTO `app_config` VALUES (4, 1, 'member-core', 'dev', 'mssql', 'sqlserver', 0, 'testsss', 20, '2020-04-30 14:33:32', '2020-04-30 14:33:32', NULL);
COMMIT;

-- ----------------------------
-- Table structure for app_config_history
-- ----------------------------
DROP TABLE IF EXISTS `app_config_history`;
CREATE TABLE `app_config_history` (
  `app_config_history_id` bigint NOT NULL AUTO_INCREMENT,
  `app_config_id` bigint DEFAULT NULL,
  `app_id` bigint DEFAULT NULL,
  `app_code` varchar(255) DEFAULT NULL,
  `environment_name` varchar(255) DEFAULT NULL,
  `key` varchar(255) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  `status` bigint DEFAULT NULL,
  `memo` varchar(255) DEFAULT NULL,
  `admin_id` bigint DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`app_config_history_id`),
  KEY `idx_app_config_history_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of app_config_history
-- ----------------------------
BEGIN;
INSERT INTO `app_config_history` VALUES (1, 1, 1, 'member-core', 'dev', 'sql_driver', 'mysql', 0, 'mysql驱动', 1, '2020-04-28 22:08:44', '2020-04-28 22:08:44', NULL);
INSERT INTO `app_config_history` VALUES (2, 2, 1, 'member-core', 'dev', 'sql_driver1', 'mysql', 0, 'mysql驱动', 3, '2020-04-28 22:09:01', '2020-04-28 22:09:01', NULL);
INSERT INTO `app_config_history` VALUES (3, 2, 1, 'member-core', 'dev', 'sql_driver1', '123321', 0, '测试备注', 20, '2020-04-30 14:08:10', '2020-04-30 14:08:10', NULL);
INSERT INTO `app_config_history` VALUES (4, 3, 0, '', '', '', '', 0, '', 0, '2020-04-30 14:27:55', '2020-04-30 14:27:55', NULL);
INSERT INTO `app_config_history` VALUES (5, 4, 1, 'member-core', 'dev', 'mssql', 'sqlserver', 0, 'testsss', 20, '2020-04-30 14:33:32', '2020-04-30 14:33:32', NULL);
INSERT INTO `app_config_history` VALUES (6, 1, 1, 'member-core', 'dev', 'sql_driver', 'mysql', 0, '', 1, '2020-04-30 15:51:42', '2020-04-30 15:51:42', NULL);
COMMIT;

-- ----------------------------
-- Table structure for app_enviroment
-- ----------------------------
DROP TABLE IF EXISTS `app_enviroment`;
CREATE TABLE `app_enviroment` (
  `app_enviroment_id` int unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int unsigned DEFAULT NULL,
  `app_code` varchar(255) DEFAULT NULL,
  `environment_name` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`app_enviroment_id`),
  KEY `idx_app_enviroment_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of app_enviroment
-- ----------------------------
BEGIN;
INSERT INTO `app_enviroment` VALUES (1, 1, 'member-core', 'pro', '2020-04-28 11:00:55', '2020-04-28 11:00:55', NULL);
INSERT INTO `app_enviroment` VALUES (2, 1, 'member-core', 'dev', '2020-04-28 11:00:55', '2020-04-28 11:00:55', NULL);
INSERT INTO `app_enviroment` VALUES (3, 1, 'member-core', 'env', '2020-04-28 17:21:28', '2020-04-28 17:21:28', NULL);
INSERT INTO `app_enviroment` VALUES (4, 1, 'member-core', 'dev', '2020-04-28 17:21:28', '2020-04-28 17:21:28', NULL);
INSERT INTO `app_enviroment` VALUES (5, 2, 'member-core', 'env', '2020-04-28 17:28:00', '2020-04-28 17:28:00', NULL);
INSERT INTO `app_enviroment` VALUES (6, 2, 'member-core', 'dev', '2020-04-28 17:28:00', '2020-04-28 17:28:00', NULL);
INSERT INTO `app_enviroment` VALUES (7, 3, 'member-core', 'env', '2020-04-28 17:28:10', '2020-04-28 17:28:10', NULL);
INSERT INTO `app_enviroment` VALUES (8, 3, 'member-core', 'dev', '2020-04-28 17:28:10', '2020-04-28 17:28:10', NULL);
INSERT INTO `app_enviroment` VALUES (9, 4, '', 'dev', '2020-04-28 21:06:40', '2020-04-28 21:06:40', NULL);
INSERT INTO `app_enviroment` VALUES (10, 4, '', 'pro', '2020-04-28 21:06:40', '2020-04-28 21:06:40', NULL);
INSERT INTO `app_enviroment` VALUES (11, 5, '', 'dev', '2020-04-29 17:16:28', '2020-04-29 17:16:28', NULL);
INSERT INTO `app_enviroment` VALUES (12, 5, '', 'pro', '2020-04-29 17:16:28', '2020-04-29 17:16:28', NULL);
INSERT INTO `app_enviroment` VALUES (13, 6, '', 'dev', '2020-04-30 09:28:35', '2020-04-30 09:28:35', NULL);
INSERT INTO `app_enviroment` VALUES (14, 6, '', 'pro', '2020-04-30 09:28:35', '2020-04-30 09:28:35', NULL);
INSERT INTO `app_enviroment` VALUES (15, 7, 'member-core333', 'dev', '2020-04-30 09:35:53', '2020-04-30 09:35:53', NULL);
INSERT INTO `app_enviroment` VALUES (16, 7, 'member-core333', 'pro', '2020-04-30 09:35:53', '2020-04-30 09:35:53', NULL);
COMMIT;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `category_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for environment
-- ----------------------------
DROP TABLE IF EXISTS `environment`;
CREATE TABLE `environment` (
  `environment_id` int unsigned NOT NULL AUTO_INCREMENT,
  `environment_name` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`environment_id`),
  KEY `idx_environment_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of environment
-- ----------------------------
BEGIN;
INSERT INTO `environment` VALUES (1, 'pro', '2020-04-28 10:43:18', '2020-04-28 10:43:18', NULL);
INSERT INTO `environment` VALUES (2, 'dev', '2020-04-28 10:43:18', '2020-04-28 10:43:18', NULL);
INSERT INTO `environment` VALUES (3, 'stage', '2020-04-30 16:17:31', '2020-04-30 16:17:31', NULL);
COMMIT;

-- ----------------------------
-- Table structure for topics
-- ----------------------------
DROP TABLE IF EXISTS `topics`;
CREATE TABLE `topics` (
  `topic_id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `category_id` int NOT NULL,
  `user_id` bigint DEFAULT NULL,
  PRIMARY KEY (`topic_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of topics
-- ----------------------------
BEGIN;
INSERT INTO `topics` VALUES (1, '测试', 1, 100);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (100, 'adminn100');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
