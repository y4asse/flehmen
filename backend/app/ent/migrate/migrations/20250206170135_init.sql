-- Create "next_actions" table
CREATE TABLE `next_actions` (`id` bigint NOT NULL AUTO_INCREMENT, `score_min` bigint NOT NULL, `score_max` bigint NOT NULL, `action` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "sukipis" table
CREATE TABLE `sukipis` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `liked_at` timestamp NOT NULL, `weight` double NULL, `height` double NULL, `x_id` varchar(255) NULL, `hobby` varchar(255) NULL, `birthday` timestamp NULL, `shoes_size` double NULL, `family` varchar(255) NULL, `nearly_station` varchar(255) NULL, `mbti` varchar(255) NULL, `created_at` timestamp NOT NULL, `sukipi_user` bigint NULL, PRIMARY KEY (`id`), INDEX `sukipis_users_user` (`sukipi_user`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "universities" table
CREATE TABLE `universities` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `deviation_lower_value` bigint NOT NULL, `deviation_upper_value` bigint NOT NULL, `abbreviation` varchar(255) NOT NULL, `prefecture` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `university_name_abbreviation` (`name`, `abbreviation`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `weight` double NOT NULL, `height` double NOT NULL, `clerk_id` varchar(255) NOT NULL, `is_male` bool NOT NULL, `created_at` timestamp NOT NULL, `user_sukipis` bigint NULL, PRIMARY KEY (`id`), UNIQUE INDEX `clerk_id` (`clerk_id`), INDEX `users_sukipis_sukipis` (`user_sukipis`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Modify "sukipis" table
ALTER TABLE `sukipis` ADD CONSTRAINT `sukipis_users_user` FOREIGN KEY (`sukipi_user`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "users" table
ALTER TABLE `users` ADD CONSTRAINT `users_sukipis_sukipis` FOREIGN KEY (`user_sukipis`) REFERENCES `sukipis` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
