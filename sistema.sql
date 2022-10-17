/*
 Navicat Premium Data Transfer

 Source Server         : BasesDatos
 Source Server Type    : MySQL
 Source Server Version : 100902
 Source Host           : localhost:3306
 Source Schema         : sistema

 Target Server Type    : MySQL
 Target Server Version : 100902
 File Encoding         : 65001

 Date: 16/10/2022 19:37:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for empleado
-- ----------------------------
DROP TABLE IF EXISTS `empleado`;
CREATE TABLE `empleado`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `correo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of empleado
-- ----------------------------
INSERT INTO `empleado` VALUES (38, 'raul antonio de la cruz hernandez', 'detretaty@gmail.com');

-- ----------------------------
-- Table structure for usuario
-- ----------------------------
DROP TABLE IF EXISTS `usuario`;
CREATE TABLE `usuario`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `correo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `usuario` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `contraseña` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of usuario
-- ----------------------------
INSERT INTO `usuario` VALUES (4, 'raul antonio de la cruz hernandez', 'detretaty@gmail.com', '0000', '0000');
INSERT INTO `usuario` VALUES (5, '', '', '', '');

SET FOREIGN_KEY_CHECKS = 1;
