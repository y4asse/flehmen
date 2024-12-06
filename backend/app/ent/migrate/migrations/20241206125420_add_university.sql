-- Create "universities" table
CREATE TABLE `universities` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `abbreviation` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
