SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'UUID',
  `login_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录名/11111',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `real_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` int(11) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` int(11) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_user_username`(`login_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, '94091b1fa6ac4a27a06c0b92155aea6a', 'admin', '$2a$10$RAQhfHSilINka4OGQRoyqODPqK7qihpwKzfH1UPu7iq0de3wBsCZS', '系统管理员', 1, '2019-12-24 12:01:43', 1, '2017-03-19 20:41:25', 1);
INSERT INTO `sys_user` VALUES (2, '84091b1fa6ac4a27a06c0b92155aea6b', 'test', '$2a$10$6abcq8HdFMUm4Qr6sOBasOOQYotyaOvKfl951I/HdMN9br5q3BGK6', '测试用户', 1, '2019-12-24 12:01:43', 1, '2017-03-19 20:41:25', 1);

SET FOREIGN_KEY_CHECKS = 1;
