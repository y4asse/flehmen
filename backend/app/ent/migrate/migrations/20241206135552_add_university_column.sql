-- Modify "universities" table
ALTER TABLE `universities` ADD COLUMN `deviation_lower_value` bigint NOT NULL, ADD COLUMN `deviation_upper_value` bigint NOT NULL, ADD COLUMN `prefecture` varchar(255) NOT NULL;
