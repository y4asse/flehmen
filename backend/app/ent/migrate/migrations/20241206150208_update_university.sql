-- Modify "universities" table
ALTER TABLE `universities` ADD UNIQUE INDEX `university_name_abbreviation` (`name`, `abbreviation`);
