-- Modify "sukipis" table
ALTER TABLE `sukipis` RENAME COLUMN `instagram_id` TO `shoes_size`, DROP COLUMN `is_male`, RENAME COLUMN `start_at` TO `liked_at`, ADD COLUMN `nearly_station` varchar(255) NULL;
