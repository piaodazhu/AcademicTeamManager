-- MySQL dump 10.13  Distrib 8.0.32, for Linux (aarch64)
--
-- Host: localhost    Database: atm
-- ------------------------------------------------------
-- Server version	8.0.32-0ubuntu0.22.04.2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `mail_config`
--

DROP TABLE IF EXISTS `mail_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `output`
--

DROP TABLE IF EXISTS `output`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `output`
--

--
-- Table structure for table `project`
--

DROP TABLE IF EXISTS `project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project`
--

--
-- Table structure for table `student`
--

DROP TABLE IF EXISTS `student`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `student` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) DEFAULT NULL COMMENT '名称',
  `type` tinyint DEFAULT NULL COMMENT '学位',
  `phone` char(12) DEFAULT NULL COMMENT '手机号',
  `email` char(80) DEFAULT NULL COMMENT '邮箱',
  `creator` bigint DEFAULT NULL COMMENT '负责人',
  `lastdis` char(20) DEFAULT '2000-00-00' COMMENT '上次讨论',
  `nextdis` char(20) DEFAULT '2000-00-00' COMMENT '本次讨论',
  `status` tinyint DEFAULT NULL COMMENT '状态',
  `remark` varchar(80) NOT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_creator` (`creator`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `student`
--
DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `type` tinyint DEFAULT '1' COMMENT '账户类型 1-普通 2-订阅',
  `name` char(10) DEFAULT NULL COMMENT '名称',
  `email` char(30) DEFAULT NULL COMMENT '邮箱',
  `password` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '密码',
  `created` bigint DEFAULT NULL COMMENT '创建时间',
  `expired` bigint DEFAULT NULL COMMENT '订阅到期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-03-04 16:21:34
