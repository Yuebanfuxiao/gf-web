/*
 Navicat Premium Data Transfer

 Source Server         : local@192.168.33.10
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : 192.168.33.10:3306
 Source Schema         : gf-web

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 07/02/2021 17:28:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '邮箱',
  `account` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账号',
  `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `nickname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像地址',
  `register_at` datetime(0) NOT NULL COMMENT '注册时间',
  `register_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '注册IP',
  `last_login_at` datetime(0) NOT NULL COMMENT '最后登陆时间',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登陆IP',
  `status` tinyint(3) UNSIGNED NOT NULL COMMENT '状态 1:启用 0:禁用',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES (1, 'fuxiao', '', '', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '2020-11-09 22:56:09', '127.0.0.1', '2020-11-09 22:56:09', '127.0.0.1', 1, '2020-11-09 22:56:09', '2020-11-09 22:56:09', NULL);
INSERT INTO `admin` VALUES (2, 'fuxiao1', '', '', 'e10adc3949ba59abbe56e057f20f883e', '', '', '2020-11-09 23:05:55', '127.0.0.1', '2020-11-09 23:05:55', '127.0.0.1', 1, '2020-11-09 23:05:55', '2020-11-18 11:23:37', NULL);
INSERT INTO `admin` VALUES (3, 'fuxiao2', '', '', '$2a$10$0iv.v.9WmoLKofazBfDnJOobH9lmZguceUa9PYSP0HRxNMxDeWNPG', '', '', '2020-11-09 23:13:05', '127.0.0.1', '2020-11-09 23:13:05', '127.0.0.1', 1, '2020-11-09 23:13:05', '2020-11-09 23:13:05', NULL);
INSERT INTO `admin` VALUES (4, 'fuxiao3', '', '', '$2a$10$lfKofbC40sqAkDqPc599k.7URrKmEPOZzl4M0YFYb49JlXRe38HiC', '', '', '2020-11-09 23:13:29', '127.0.0.1', '2020-12-03 09:56:52', '127.0.0.1', 1, '2020-11-09 23:13:29', '2020-12-03 09:56:52', NULL);
INSERT INTO `admin` VALUES (5, 'fuxiao5', '', '', '$2a$10$t6zvoaruOxfBxBnbSgwM/.E76o4zclmm7MivWCcqVAKSqkMEppDSq', '', '', '2020-11-20 11:25:11', '127.0.0.1', '2020-11-20 11:25:11', '127.0.0.1', 1, '2020-11-20 11:25:11', '2020-11-20 11:25:11', NULL);
INSERT INTO `admin` VALUES (6, 'fuxiao6', '', '', '$2a$10$2LOi4zaD8Xt5hTBrpMEwre0cMnSm/2nWYX9nb2Nj3dWLeND98JPBC', '', '', '2020-11-20 11:27:15', '127.0.0.1', '2020-11-20 11:27:15', '127.0.0.1', 1, '2020-11-20 11:27:15', '2020-11-20 11:27:15', NULL);
INSERT INTO `admin` VALUES (7, 'fuxiao7', '', '', '$2a$10$Tj5g83HR2MKJYeKMfipwTe/2w/sROs21uYLUhQ10pvXeFf2v.RTvm', '', '', '2020-11-20 11:37:03', '127.0.0.1', '2020-11-20 11:37:03', '127.0.0.1', 1, '2020-11-20 11:37:03', '2020-11-20 11:37:03', NULL);
INSERT INTO `admin` VALUES (8, 'fuxiao8', '', '', '$2a$10$ThOokJRDHTIEkDRseLDWDOhvcqV2GiT24Bx4jE2hSOE/S8RhNG.DW', '', '', '2020-11-20 11:38:22', '127.0.0.1', '2020-11-20 11:38:22', '127.0.0.1', 1, '2020-11-20 11:38:22', '2020-11-20 11:38:22', NULL);
INSERT INTO `admin` VALUES (9, 'fuxiao9', '', '', '$2a$10$DbJU/mhGz9hz9Ym/LxcrLuxOBa0yEYG2Tg5QXVlsFdxjL4iTcamWG', '', '', '2020-11-20 11:41:37', '127.0.0.1', '2020-11-20 11:41:37', '127.0.0.1', 1, '2020-11-20 11:41:37', '2020-11-20 11:41:37', NULL);
INSERT INTO `admin` VALUES (10, 'fuxiao9', '', '', '123456', '', '', '2020-11-20 12:22:03', '127.0.0.1', '2020-11-20 12:22:03', '127.0.0.1', 1, '2020-11-20 12:22:03', '2020-11-20 12:22:03', NULL);
INSERT INTO `admin` VALUES (11, 'fuxiao10', '', '', '$2a$10$nSimdSZCZ4Nu.ZzZZIMAt.lxIZ.TrgABv4.K4eOFhqFqowo/vH37i', '', '', '2020-11-20 12:23:13', '127.0.0.1', '2020-11-20 12:23:13', '127.0.0.1', 1, '2020-11-20 12:23:13', '2020-11-20 12:23:13', NULL);
INSERT INTO `admin` VALUES (12, '', '', '', '$2a$10$22bAGrXEE198XK4TuW2wt.0EI/3fUOyt9DHsZw4pvAVVzfGV96Qqq', '', '', '2020-11-20 14:39:00', '127.0.0.1', '2020-11-20 14:39:00', '127.0.0.1', 0, '2020-11-20 14:39:00', '2020-11-20 14:39:00', NULL);
INSERT INTO `admin` VALUES (13, '', '', '', '$2a$10$MBG8nGA38T6x7cKac0Gj0.KD81.64dYEMq0g76Tle.aL1CqabBqay', '', '', '2020-11-20 14:39:02', '127.0.0.1', '2020-11-20 14:39:02', '127.0.0.1', 0, '2020-11-20 14:39:02', '2020-11-20 14:39:02', NULL);
INSERT INTO `admin` VALUES (14, '', '', '', '$2a$10$Nyd9BNXDWC6XDTmSs4seQeXsAs6ZWAOxZPAm8pHNdGmvXP1Vm1vw.', '', '', '2020-11-20 14:39:03', '127.0.0.1', '2020-11-20 14:39:03', '127.0.0.1', 0, '2020-11-20 14:39:03', '2020-11-20 14:39:03', NULL);
INSERT INTO `admin` VALUES (15, '', '', '', '$2a$10$Oh1UFVoG7HwXd5vToROQr.zbEWroELHo5ivS8lN8GbOzJERki11Qi', '', '', '2020-11-20 14:39:25', '127.0.0.1', '2020-11-20 14:39:25', '127.0.0.1', 0, '2020-11-20 14:39:25', '2020-11-20 14:39:25', NULL);
INSERT INTO `admin` VALUES (16, '', '', '', '$2a$10$PK0cVPeQvUcZacGpj/wbveoe5mjKdDzGuxHFHVcgzAQ4stMENK10O', '', '', '2020-11-20 14:41:52', '127.0.0.1', '2020-11-20 14:41:52', '127.0.0.1', 0, '2020-11-20 14:41:52', '2020-11-20 14:41:52', NULL);
INSERT INTO `admin` VALUES (17, '', '', '', '$2a$10$ZqzQKY1sPgLFAK/DbZzb3e/z1Tbmz6z7wg.yd59JN3blzEI2jh4eW', '', '', '2020-11-20 14:43:04', '127.0.0.1', '2020-11-20 14:43:04', '127.0.0.1', 0, '2020-11-20 14:43:04', '2020-11-20 14:43:04', NULL);
INSERT INTO `admin` VALUES (18, '', '', '', '$2a$10$RtfdPCpe5/7QzElAdm/BYuH7rvn0fCsfRMTPX8.MdL8W6Q.R5Ae96', '', '', '2020-11-20 14:43:52', '127.0.0.1', '2020-11-20 14:43:52', '127.0.0.1', 0, '2020-11-20 14:43:52', '2020-11-20 14:43:52', NULL);
INSERT INTO `admin` VALUES (19, '', '', '', '', '拂晓', 'http://www.baidu.com', '2020-11-20 16:20:26', '', '2020-11-20 16:20:26', '', 0, '2020-11-20 15:17:01', '2020-11-23 14:12:49', NULL);
INSERT INTO `admin` VALUES (20, 'fuxiao20', '', '', '$2a$10$2hh63LcvdUMYTiCBF//cAO8UVb9C1DlnbswL2Er5c4JdL1gIK9nSO', '拂晓', 'https://www.baidu.com/img/flexible/logo/pc/result.png', '2020-11-23 11:32:40', '127.0.0.1', '2020-11-23 11:32:40', '127.0.0.1', 1, '2020-11-23 11:32:40', '2020-11-23 11:32:40', NULL);

-- ----------------------------
-- Table structure for admin_login_record
-- ----------------------------
DROP TABLE IF EXISTS `admin_login_record`;
CREATE TABLE `admin_login_record`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `admin_id` int(10) UNSIGNED NOT NULL COMMENT '管理员ID',
  `login_at` datetime(0) NOT NULL COMMENT '登陆时间',
  `login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登陆IP',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员登陆记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_login_record
-- ----------------------------
INSERT INTO `admin_login_record` VALUES (1, 4, '2020-11-09 23:59:32', '127.0.0.1', '2020-11-09 23:59:32', '2020-11-09 23:59:32', NULL);
INSERT INTO `admin_login_record` VALUES (2, 4, '2020-11-09 23:59:49', '127.0.0.1', '2020-11-09 23:59:49', '2020-11-09 23:59:49', NULL);
INSERT INTO `admin_login_record` VALUES (3, 4, '2020-11-10 11:10:15', '127.0.0.1', '2020-11-10 11:10:15', '2020-11-10 11:10:15', NULL);
INSERT INTO `admin_login_record` VALUES (4, 4, '2020-11-10 11:58:22', '127.0.0.1', '2020-11-10 11:58:22', '2020-11-10 11:58:22', NULL);
INSERT INTO `admin_login_record` VALUES (5, 4, '2020-11-10 18:04:00', '127.0.0.1', '2020-11-10 18:04:00', '2020-11-10 18:04:00', NULL);
INSERT INTO `admin_login_record` VALUES (6, 4, '2020-11-10 18:05:52', '127.0.0.1', '2020-11-10 18:05:52', '2020-11-10 18:05:52', NULL);
INSERT INTO `admin_login_record` VALUES (7, 4, '2020-11-11 14:15:50', '127.0.0.1', '2020-11-11 14:15:50', '2020-11-11 14:15:50', NULL);
INSERT INTO `admin_login_record` VALUES (8, 4, '2020-11-11 14:17:34', '127.0.0.1', '2020-11-11 14:17:34', '2020-11-11 14:17:34', NULL);
INSERT INTO `admin_login_record` VALUES (9, 4, '2020-11-11 14:19:11', '127.0.0.1', '2020-11-11 14:19:11', '2020-11-11 14:19:11', NULL);
INSERT INTO `admin_login_record` VALUES (10, 4, '2020-11-11 18:16:22', '127.0.0.1', '2020-11-11 18:16:22', '2020-11-11 18:16:22', NULL);
INSERT INTO `admin_login_record` VALUES (11, 4, '2020-11-12 12:13:52', '127.0.0.1', '2020-11-12 12:13:52', '2020-11-12 12:13:52', NULL);
INSERT INTO `admin_login_record` VALUES (12, 4, '2020-11-12 12:18:24', '127.0.0.1', '2020-11-12 12:18:24', '2020-11-12 12:18:24', NULL);
INSERT INTO `admin_login_record` VALUES (13, 4, '2020-11-12 12:41:38', '127.0.0.1', '2020-11-12 12:41:38', '2020-11-12 12:41:38', NULL);
INSERT INTO `admin_login_record` VALUES (14, 4, '2020-11-12 14:23:27', '127.0.0.1', '2020-11-12 14:23:27', '2020-11-12 14:23:27', NULL);
INSERT INTO `admin_login_record` VALUES (15, 4, '2020-11-12 14:25:47', '127.0.0.1', '2020-11-12 14:25:47', '2020-11-12 14:25:47', NULL);
INSERT INTO `admin_login_record` VALUES (16, 4, '2020-11-12 17:04:28', '127.0.0.1', '2020-11-12 17:04:28', '2020-11-12 17:04:28', NULL);
INSERT INTO `admin_login_record` VALUES (17, 4, '2020-11-12 17:04:34', '127.0.0.1', '2020-11-12 17:04:34', '2020-11-12 17:04:34', NULL);
INSERT INTO `admin_login_record` VALUES (18, 4, '2020-11-12 17:32:27', '127.0.0.1', '2020-11-12 17:32:27', '2020-11-12 17:32:27', NULL);
INSERT INTO `admin_login_record` VALUES (19, 4, '2020-11-12 17:34:59', '127.0.0.1', '2020-11-12 17:34:59', '2020-11-12 17:34:59', NULL);
INSERT INTO `admin_login_record` VALUES (20, 4, '2020-11-12 17:38:16', '127.0.0.1', '2020-11-12 17:38:16', '2020-11-12 17:38:16', NULL);
INSERT INTO `admin_login_record` VALUES (21, 4, '2020-11-13 15:39:56', '127.0.0.1', '2020-11-13 15:39:56', '2020-11-13 15:39:56', NULL);
INSERT INTO `admin_login_record` VALUES (22, 4, '2020-11-13 15:41:43', '127.0.0.1', '2020-11-13 15:41:43', '2020-11-13 15:41:43', NULL);
INSERT INTO `admin_login_record` VALUES (23, 4, '2020-11-13 16:37:39', '127.0.0.1', '2020-11-13 16:37:39', '2020-11-13 16:37:39', NULL);
INSERT INTO `admin_login_record` VALUES (24, 4, '2020-11-13 16:45:49', '127.0.0.1', '2020-11-13 16:45:49', '2020-11-13 16:45:49', NULL);
INSERT INTO `admin_login_record` VALUES (25, 4, '2020-11-13 16:47:29', '127.0.0.1', '2020-11-13 16:47:29', '2020-11-13 16:47:29', NULL);
INSERT INTO `admin_login_record` VALUES (26, 4, '2020-11-13 16:48:00', '127.0.0.1', '2020-11-13 16:48:00', '2020-11-13 16:48:00', NULL);
INSERT INTO `admin_login_record` VALUES (27, 4, '2020-11-13 16:48:33', '127.0.0.1', '2020-11-13 16:48:33', '2020-11-13 16:48:33', NULL);
INSERT INTO `admin_login_record` VALUES (28, 4, '2020-11-13 16:48:49', '127.0.0.1', '2020-11-13 16:48:49', '2020-11-13 16:48:49', NULL);
INSERT INTO `admin_login_record` VALUES (29, 4, '2020-11-13 17:14:28', '127.0.0.1', '2020-11-13 17:14:28', '2020-11-13 17:14:28', NULL);
INSERT INTO `admin_login_record` VALUES (30, 4, '2020-11-13 17:16:34', '127.0.0.1', '2020-11-13 17:16:34', '2020-11-13 17:16:34', NULL);
INSERT INTO `admin_login_record` VALUES (31, 4, '2020-11-17 17:09:04', '127.0.0.1', '2020-11-17 17:09:04', '2020-11-17 17:09:04', NULL);
INSERT INTO `admin_login_record` VALUES (32, 4, '2020-11-17 18:28:30', '127.0.0.1', '2020-11-17 18:28:30', '2020-11-17 18:28:30', NULL);
INSERT INTO `admin_login_record` VALUES (33, 4, '2020-11-18 11:27:51', '127.0.0.1', '2020-11-18 11:27:51', '2020-11-18 11:27:51', NULL);
INSERT INTO `admin_login_record` VALUES (34, 4, '2020-11-18 11:28:45', '127.0.0.1', '2020-11-18 11:28:45', '2020-11-18 11:28:45', NULL);
INSERT INTO `admin_login_record` VALUES (35, 4, '2020-11-18 11:28:59', '127.0.0.1', '2020-11-18 11:28:59', '2020-11-18 11:28:59', NULL);
INSERT INTO `admin_login_record` VALUES (36, 4, '2020-11-23 11:30:22', '127.0.0.1', '2020-11-23 11:30:22', '2020-11-23 11:30:22', NULL);
INSERT INTO `admin_login_record` VALUES (37, 4, '2020-11-30 11:51:17', '127.0.0.1', '2020-11-30 11:51:17', '2020-11-30 11:51:17', NULL);
INSERT INTO `admin_login_record` VALUES (38, 4, '2020-12-03 09:56:52', '127.0.0.1', '2020-12-03 09:56:52', '2020-12-03 09:56:52', NULL);

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`  (
  `admin_id` int(10) UNSIGNED NOT NULL,
  `role_id` int(10) UNSIGNED NOT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_role
-- ----------------------------

-- ----------------------------
-- Table structure for permission_node
-- ----------------------------
DROP TABLE IF EXISTS `permission_node`;
CREATE TABLE `permission_node`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '节点名称',
  `label` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '节点标签',
  `method` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '节点操作类型',
  `path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '节点路径',
  `status` tinyint(3) UNSIGNED NOT NULL COMMENT '节点状态 0：禁用  1：启用',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '节点备注',
  `created_at` datetime(0) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限节点表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permission_node
-- ----------------------------
INSERT INTO `permission_node` VALUES (1, '创建节点', '', 'DELETE', '/backend/permission/create-node', 1, '创建权限控制节点', '2020-12-21 15:10:14', '2020-12-21 15:10:14', NULL);
INSERT INTO `permission_node` VALUES (2, '创建节点', '', 'DELETE', '/backend/permission/create-node', 1, '创建权限控制节点', '2020-12-21 15:10:21', '2020-12-21 15:10:21', NULL);
INSERT INTO `permission_node` VALUES (3, '创建节点', '', 'DELETE', '/backend/permission/create-node', 1, '创建权限控制节点', '2020-12-21 15:10:40', '2020-12-21 15:10:40', NULL);

-- ----------------------------
-- Table structure for permission_policy
-- ----------------------------
DROP TABLE IF EXISTS `permission_policy`;
CREATE TABLE `permission_policy`  (
  `ptype` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v0` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v1` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v2` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v3` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v4` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v5` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限策略表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permission_policy
-- ----------------------------
INSERT INTO `permission_policy` VALUES ('p', '16', '3', '', '', '', '');
INSERT INTO `permission_policy` VALUES ('p', '16', '4', '', '', '', '');
INSERT INTO `permission_policy` VALUES ('g', '4', '16', '', '', '', '');
INSERT INTO `permission_policy` VALUES ('p', '18', '3', '', '', '', '');
INSERT INTO `permission_policy` VALUES ('p', '18', '4', '', '', '', '');
INSERT INTO `permission_policy` VALUES ('p', '19', '3', '', '', '', '');
INSERT INTO `permission_policy` VALUES ('p', '19', '4', '', '', '', '');
INSERT INTO `permission_policy` VALUES ('p', '20', '3', '', '', '', '');
INSERT INTO `permission_policy` VALUES ('p', '20', '4', '', '', '', '');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色备注',
  `status` tinyint(3) UNSIGNED NOT NULL COMMENT '状态 1:启用 0:禁用',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (3, '超级管理员', '系统超级管理员', 1, '2020-11-26 11:04:33', '2020-11-26 11:04:33', '2020-11-26 12:31:23');
INSERT INTO `role` VALUES (4, '超级管理员', '系统超级管理员啊啊啊', 1, '2020-11-26 11:04:50', '2020-11-30 16:00:51', NULL);
INSERT INTO `role` VALUES (5, '超级管理员', '系统超级管理员', 1, '2020-11-26 11:08:19', '2020-11-26 11:08:19', '2020-11-30 11:51:47');
INSERT INTO `role` VALUES (6, '超级管理员', '系统超级管理员', 1, '2020-11-26 11:08:31', '2020-11-26 11:08:31', '2020-11-30 11:51:47');
INSERT INTO `role` VALUES (8, '超级管理员', '系统超级管理员', 1, '2020-11-26 12:31:44', '2020-11-26 12:31:44', '2020-11-26 15:33:26');
INSERT INTO `role` VALUES (9, '超级管理员', '系统超级管理员', 1, '2020-11-26 14:34:25', '2020-11-26 14:34:25', '2020-11-26 15:33:26');
INSERT INTO `role` VALUES (10, '超级管理员', '系统超级管理员', 1, '2020-11-26 15:29:42', '2020-11-26 15:29:42', '2020-11-26 15:30:25');
INSERT INTO `role` VALUES (11, '超级管理员', '系统超级管理员', 1, '2020-11-27 12:13:06', '2020-11-27 12:13:06', '2020-11-30 11:51:47');
INSERT INTO `role` VALUES (12, '超级管理员', '系统超级管理员', 1, '2020-11-27 12:13:15', '2020-11-27 12:13:15', '2020-11-30 11:51:47');
INSERT INTO `role` VALUES (13, '超级管理员', '系统超级管理员', 1, '2020-11-27 12:14:10', '2020-11-27 12:14:10', '2020-11-30 11:51:47');
INSERT INTO `role` VALUES (14, '超级管理员', '系统超级管理员', 1, '2020-11-27 12:18:32', '2020-11-27 12:18:32', '2020-11-27 12:27:19');
INSERT INTO `role` VALUES (15, '超级管理员', '系统超级管理员', 1, '2020-11-27 12:28:54', '2020-11-27 12:28:54', '2020-11-30 11:51:47');
INSERT INTO `role` VALUES (16, '超级管理员', '系统超级管理员', 1, '2020-11-30 12:12:07', '2020-11-30 12:12:07', NULL);
INSERT INTO `role` VALUES (18, '超级管理员', '系统超级管理员', 1, '2020-12-03 09:57:16', '2020-12-03 09:57:16', NULL);
INSERT INTO `role` VALUES (19, '超级管理员', '系统超级管理员', 1, '2020-12-03 09:58:53', '2020-12-03 09:58:53', NULL);
INSERT INTO `role` VALUES (20, '超级管理员', '系统超级管理员', 1, '2020-12-03 10:00:09', '2020-12-03 10:00:09', NULL);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `wechat_user_id` int(10) UNSIGNED NOT NULL COMMENT '微信用户ID',
  `register_at` datetime(0) NOT NULL COMMENT '注册时间',
  `register_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '注册IP',
  `created_at` datetime(0) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------

-- ----------------------------
-- Table structure for wechat_user
-- ----------------------------
DROP TABLE IF EXISTS `wechat_user`;
CREATE TABLE `wechat_user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `openid` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '微信openid',
  `unionid` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '微信授权联合ID',
  `nickname` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '微信昵称',
  `gender` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别(0：未知  1：男  2：女）',
  `avatar` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '微信头像网址',
  `country` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户所在国家',
  `province` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户所在省份',
  `city` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户所在市',
  `created_at` datetime(0) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '软删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '微信用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of wechat_user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
