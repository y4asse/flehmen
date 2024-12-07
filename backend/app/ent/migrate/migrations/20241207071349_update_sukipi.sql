-- Modify "sukipis" table
ALTER TABLE `sukipis` MODIFY COLUMN `weight` double NULL, MODIFY COLUMN `height` double NULL, MODIFY COLUMN `x_id` varchar(255) NULL, MODIFY COLUMN `instagram_id` varchar(255) NULL, MODIFY COLUMN `start_at` timestamp NOT NULL, ADD COLUMN `birthday` timestamp NULL, ADD COLUMN `family` varchar(255) NULL;
