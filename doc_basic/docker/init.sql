-- ----------------------------
-- 创建数据库
-- ----------------------------
-- create schema gf_study collate utf8mb4_bin;
CREATE DATABASE IF NOT EXISTS gf_study CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_bin;


-- ----------------------------
-- 创建表 && 数据初始化
-- ----------------------------
use gf_study;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Drop Table
-- ----------------------------

drop table if exists user;


-- ----------------------------
-- Create Table
-- ----------------------------
CREATE TABLE `user` (
                        `uid` int(11) NOT NULL AUTO_INCREMENT,
                        `name` varchar(255) DEFAULT NULL,
                        `site` varchar(255) DEFAULT NULL,
                        PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 ;

SET FOREIGN_KEY_CHECKS = 1;
