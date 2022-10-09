/*
Navicat MySQL Data Transfer

Source Server         : APLNDC
Source Server Version : 100901
Source Host           : 192.168.1.100:3306
Source Database       : aplndc

Target Server Type    : MYSQL
Target Server Version : 100901
File Encoding         : 65001

Date: 2022-08-19 11:24:13
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tareas
-- ----------------------------
DROP TABLE IF EXISTS `tareas`;
CREATE TABLE `tareas` (
  `TareaIdentificaciónNúmero` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'TaskID',
  `UsuarioIdentificaciónNúmero` int(10) unsigned NOT NULL COMMENT 'UserID',
  `Nombre` varchar(64) NOT NULL COMMENT 'Name',
  `Descripción` text NOT NULL COMMENT 'Description',
  `Prioridad` enum('critical','importante','sérieux','annyira nem qva fontos','unwichtig') NOT NULL COMMENT 'Priority',
  `Hecho` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'Done',
  `Plazo` datetime NOT NULL COMMENT 'Deadline',
  `CreadoEn` datetime NOT NULL DEFAULT current_timestamp() COMMENT 'CreatedAt',
  PRIMARY KEY (`TareaIdentificaciónNúmero`),
  KEY `UsuarioIdentificaciónNúmero` (`UsuarioIdentificaciónNúmero`),
  CONSTRAINT `tareas_ibfk_1` FOREIGN KEY (`UsuarioIdentificaciónNúmero`) REFERENCES `usuarios` (`UsuarioIdentificaciónNúmero`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tareas
-- ----------------------------
INSERT INTO `tareas` VALUES ('1', '1', 'APLNDC', 'Megcsinálni ezt a szart', 'critical', '0', '2022-08-17 00:00:00', '2022-08-15 21:13:26');
INSERT INTO `tareas` VALUES ('2', '2', 'Erdély', 'Vissza kell foglalni erdélyt', 'sérieux', '0', '1920-06-04 17:12:00', '2022-08-15 21:15:46');
INSERT INTO `tareas` VALUES ('3', '1', 'prueba', 'Esto es una prueba', 'sérieux', '0', '2020-12-24 12:05:40', '2022-08-15 23:21:59');
INSERT INTO `tareas` VALUES ('4', '1', 'prueba', 'Esto es una prueba', 'sérieux', '0', '2020-12-24 12:05:40', '2022-08-15 23:22:01');
INSERT INTO `tareas` VALUES ('5', '1', 'prueba', 'Esto es una prueba', 'sérieux', '0', '2020-12-24 12:05:40', '2022-08-15 23:22:02');
INSERT INTO `tareas` VALUES ('6', '1', 'prueba', 'Esto es una prueba', 'sérieux', '0', '2020-12-24 12:05:40', '2022-08-15 23:22:03');
INSERT INTO `tareas` VALUES ('7', '1', 'prueba', 'Esto es una prueba', 'sérieux', '0', '2020-12-24 12:05:40', '2022-08-15 23:22:03');
INSERT INTO `tareas` VALUES ('8', '1', 'prueba', 'Esto es una prueba', 'sérieux', '0', '2020-12-24 12:05:40', '2022-08-15 23:22:04');
INSERT INTO `tareas` VALUES ('9', '1', 'prueba', 'Esto es una prueba', 'sérieux', '0', '2020-12-24 12:05:40', '2022-08-15 23:22:05');

-- ----------------------------
-- Table structure for usuarios
-- ----------------------------
DROP TABLE IF EXISTS `usuarios`;
CREATE TABLE `usuarios` (
  `UsuarioIdentificaciónNúmero` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'UserID',
  `Nombre` varchar(128) NOT NULL COMMENT 'Name',
  `sesiónIdentificaciónNúmero` char(16) DEFAULT NULL COMMENT 'SessionID',
  `Contraseña hash` char(64) DEFAULT NULL COMMENT 'Password Hash',
  `Sal` char(8) DEFAULT NULL COMMENT 'Salt',
  PRIMARY KEY (`UsuarioIdentificaciónNúmero`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of usuarios
-- ----------------------------
INSERT INTO `usuarios` VALUES ('1', 'admin', 'csakamechwart123', 'd82494f05d6917ba02f7aaa29689ccb444bb73f20380876cb05d1f37537b7892', 'admin');
INSERT INTO `usuarios` VALUES ('2', 'Carlos', 'asdasddsadsaqwer', 'acfee723e971747155301f688aecb9ed9bdfbbd977bba7c9546605c287a87105', 'pene');
SET FOREIGN_KEY_CHECKS=1;
