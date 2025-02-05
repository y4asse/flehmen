-- Create "next_actions" table
CREATE TABLE `next_actions` (`id` bigint NOT NULL AUTO_INCREMENT, `score_min` bigint NOT NULL, `score_max` bigint NOT NULL, `action` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
