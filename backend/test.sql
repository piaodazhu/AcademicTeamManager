SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;


DROP TABLE IF EXISTS `project`;
CREATE TABLE `project` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(200) DEFAULT NULL COMMENT '项目名称',
  `begin_time` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '开始时间',
  `over_time` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '结束时间',
  `remark` varchar(80) DEFAULT NULL COMMENT '备注',
  `outputlist` json DEFAULT NULL COMMENT '成果产出',
  `finishrate` decimal(10,2) DEFAULT NULL COMMENT '完成率',
  `creator` bigint DEFAULT NULL COMMENT '创建人',
  PRIMARY KEY (`id`),
  KEY `idx_creator` (`creator`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3;



DROP TABLE IF EXISTS `student`;
CREATE TABLE `student` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) DEFAULT NULL COMMENT '名称',
  `phone` char(12) DEFAULT NULL COMMENT '手机号',
  `email` char(80) DEFAULT NULL COMMENT '邮箱',
  `creator` bigint DEFAULT NULL COMMENT '负责人',
  `lastdis` bigint DEFAULT NULL COMMENT '上次讨论',
  `nextdis` bigint DEFAULT NULL COMMENT '本次讨论',
  `status` tinyint DEFAULT NULL COMMENT '状态',
  `remark` varchar(80) NOT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_creator` (`creator`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb3;

DROP TABLE IF EXISTS `mail_config`;
CREATE TABLE `mail_config` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `stmp` char(50) DEFAULT NULL COMMENT 'ip地址或域名',
  `port` int DEFAULT NULL COMMENT '端口号',
  `auth_code` char(80) DEFAULT NULL COMMENT '授权码',
  `email` char(80) NOT NULL COMMENT '邮箱账号',
  `status` tinyint DEFAULT NULL COMMENT '服务状态',
  `creator` bigint DEFAULT NULL COMMENT '创建人',
  `created` bigint DEFAULT NULL COMMENT '创建时间',
  `updated` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_creator` (`creator`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb3;

DROP TABLE IF EXISTS `output`;
CREATE TABLE `output` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) DEFAULT NULL COMMENT '成果名称',
  `type` tinyint DEFAULT NULL COMMENT '成果类型',
  `weight` decimal(10,2) DEFAULT NULL COMMENT '权重',
  `description` varchar(200) DEFAULT NULL COMMENT '成果描述',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态 1-已完成 2-进行中',
  `creator` bigint DEFAULT NULL COMMENT '创建人',
  PRIMARY KEY (`id`),
  KEY `idx_creator` (`creator`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb3;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `type` tinyint DEFAULT '1' COMMENT '账户类型 1-普通 2-订阅',
  `name` char(10) DEFAULT NULL COMMENT '名称',
  `email` char(80) DEFAULT NULL COMMENT '邮箱',
  `password` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '密码',
  `created` bigint DEFAULT NULL COMMENT '创建时间',
  `expired` bigint DEFAULT NULL COMMENT '订阅到期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb3;

SET FOREIGN_KEY_CHECKS = 1;
