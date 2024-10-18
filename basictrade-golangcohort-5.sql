/*
 Navicat Premium Data Transfer

 Source Server         : simpad_kudus
 Source Server Type    : MySQL
 Source Server Version : 100424 (10.4.24-MariaDB)
 Source Host           : localhost:3306
 Source Schema         : basictrade-golangcohort-5

 Target Server Type    : MySQL
 Target Server Version : 100424 (10.4.24-MariaDB)
 File Encoding         : 65001

 Date: 18/10/2024 22:35:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `name` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `email` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `password` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admins
-- ----------------------------
INSERT INTO `admins` VALUES (1, 'b3f29ce3-81d1-457b-9e24-365c630bffe3', 'reza', 'reza@mail.com', '$2a$08$we2OZc8XyU6HsI1KF6Yptu0oRjPWpRHIzkRmzDRmXYvh6L0amCe1O', '2024-10-18 22:31:40.592', '2024-10-18 22:31:40.592');

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `name` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `image_url` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `admin_id` bigint UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_products_admin_id`(`admin_id` ASC) USING BTREE,
  CONSTRAINT `fk_products_admin` FOREIGN KEY (`admin_id`) REFERENCES `admins` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of products
-- ----------------------------
INSERT INTO `products` VALUES (2, 'abd5993f-3fcd-44b6-ac15-0e0a92a5b5ef', 'Asus', 'https://res.cloudinary.com/dks3wiawl/image/upload/v1728766223/golang-cohort-5/Screenshot%20%282%29.png', 1, '2024-10-18 22:32:17.871', '2024-10-18 22:32:36.789');

-- ----------------------------
-- Table structure for variants
-- ----------------------------
DROP TABLE IF EXISTS `variants`;
CREATE TABLE `variants`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `variant_name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `quantity` bigint NOT NULL,
  `product_id` bigint UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_products_variants`(`product_id` ASC) USING BTREE,
  CONSTRAINT `fk_products_variants` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of variants
-- ----------------------------
INSERT INTO `variants` VALUES (1, 'f7db2181-cf2a-4d96-9c84-daacf71f2673', 'Macbook Air M1', 100, 2, '2024-10-18 22:33:30.170', '2024-10-18 22:33:30.170');
INSERT INTO `variants` VALUES (2, '4ee21faf-a13a-4b65-869f-147e8c6b7f1d', 'Lepi', 8, 2, '2024-10-18 22:33:38.946', '2024-10-18 22:33:55.304');
INSERT INTO `variants` VALUES (3, '4d31339e-a514-4c86-a38e-446b89572d21', 'Laptop', 100, 2, '2024-10-18 22:33:59.284', '2024-10-18 22:33:59.284');

SET FOREIGN_KEY_CHECKS = 1;
