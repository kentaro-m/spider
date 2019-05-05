USE `spiderdb`;

DROP TABLE IF EXISTS `sites`;

CREATE TABLE `sites` (
   `id` varchar(36) NOT NULL,
   `title` mediumtext,
   `url` mediumtext,
   `created_at` datetime DEFAULT NULL,
   `updated_at` datetime DEFAULT NULL,
   PRIMARY KEY (`id`),
   UNIQUE KEY `id_UNIQUE` (`id`),
   INDEX `idx_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `articles`;

CREATE TABLE `articles` (
  `id` varchar(36) NOT NULL,
  `title` varchar(512),
  `url` mediumtext,
  `thumbnail_url` mediumtext,
  `pub_date` datetime DEFAULT NULL,
  `site_id` varchar(36) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `title_UNIQUE` (`title`),
  INDEX `idx_pub_date` (`pub_date`),
  INDEX `idx_site_id` (`site_id`),
  CONSTRAINT `fk_site_id`
    FOREIGN KEY (`site_id`)
    REFERENCES `sites` (`id`)
    ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;