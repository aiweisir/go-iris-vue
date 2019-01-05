/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50714
Source Host           : localhost:3306
Source Database       : casbin

Target Server Type    : MYSQL
Target Server Version : 50714
File Encoding         : 65001

Date: 2019-01-05 14:55:13
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  KEY `IDX_casbin_rule_v3` (`v3`),
  KEY `IDX_casbin_rule_v4` (`v4`),
  KEY `IDX_casbin_rule_v5` (`v5`),
  KEY `IDX_casbin_rule_p_type` (`p_type`),
  KEY `IDX_casbin_rule_v0` (`v0`),
  KEY `IDX_casbin_rule_v1` (`v1`),
  KEY `IDX_casbin_rule_v2` (`v2`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', 'alice', '/dataset1/*', 'get', '.*', null, null);
INSERT INTO `casbin_rule` VALUES ('p', '3', '*', 'POST', '.*', null, null);
INSERT INTO `casbin_rule` VALUES ('p', '2', '/a/*', 'GET|POST', '.*', null, null);
INSERT INTO `casbin_rule` VALUES ('p', '1', '/*', 'ANY', '.*', '', '');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `appid` varchar(255) NOT NULL DEFAULT '',
  `secret` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `phone` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `userface` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'root', 'x04jpoIrc8/mvNRqAG59Wg==', '', '', '', '', '', '', '2019-01-05 14:31:29', null);
INSERT INTO `user` VALUES ('2', 'yhm01', 'x04jpoIrc8/mvNRqAG59Wg==', '', '', '', '', '', '', '2019-01-04 16:45:18', null);
INSERT INTO `user` VALUES ('3', 'yhm02', 'x04jpoIrc8/mvNRqAG59Wg==', '', '', '', '', '', '', '2019-01-02 11:59:15', null);
INSERT INTO `user` VALUES ('4', 'yhm03', 'x04jpoIrc8/mvNRqAG59Wg==', '', '', '', '', '', '', '2019-01-05 13:03:50', null);
INSERT INTO `user` VALUES ('5', 'yhm04', 'x04jpoIrc8/mvNRqAG59Wg==', '', '', '', '', '', '', '2019-01-05 13:06:37', null);
