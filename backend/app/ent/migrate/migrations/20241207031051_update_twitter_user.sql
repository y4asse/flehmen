-- Create "twitter_users" table
CREATE TABLE `twitter_users` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `username` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Modify "tweets" table
ALTER TABLE `tweets` ADD COLUMN `reply_twitter_user_id` bigint NULL, ADD INDEX `tweets_twitter_users_replies` (`reply_twitter_user_id`), ADD CONSTRAINT `tweets_twitter_users_replies` FOREIGN KEY (`reply_twitter_user_id`) REFERENCES `twitter_users` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
