USE `spiderdb`;

DROP TABLE IF EXISTS `articles`;

CREATE TABLE `articles` (
  `id` varchar(36) NOT NULL,
  `title` mediumtext,
  `url` mediumtext,
  `pub_date` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  INDEX `idx_pub_date` (`pub_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;