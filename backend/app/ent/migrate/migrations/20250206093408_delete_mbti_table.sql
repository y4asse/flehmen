-- Modify "sukipis" table
ALTER TABLE `sukipis` RENAME COLUMN `sukipi_mbti` TO `sukipi_user`, ADD COLUMN `mbti` varchar(255) NULL, DROP INDEX `sukipis_mbtis_mbti`, ADD INDEX `sukipis_users_user` (`sukipi_user`), DROP FOREIGN KEY `sukipis_mbtis_mbti`;
-- Modify "users" table
ALTER TABLE `users` RENAME COLUMN `user_mbti` TO `user_sukipis`, DROP INDEX `users_mbtis_mbti`, ADD INDEX `users_sukipis_sukipis` (`user_sukipis`), DROP FOREIGN KEY `users_mbtis_mbti`;
-- Modify "sukipis" table
ALTER TABLE `sukipis` ADD CONSTRAINT `sukipis_users_user` FOREIGN KEY (`sukipi_user`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "users" table
ALTER TABLE `users` ADD CONSTRAINT `users_sukipis_sukipis` FOREIGN KEY (`user_sukipis`) REFERENCES `sukipis` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
-- Drop "mbtis" table
DROP TABLE `mbtis`;
