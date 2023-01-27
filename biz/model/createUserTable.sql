CREATE DATABASE "hertz_demo";
CREATE TABLE `user`
(
    `user_id`     bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'UserId',
    `user_name`   varchar(64) NOT NULL DEFAULT '' COMMENT 'UserName',
    `password`    varchar(64) NOT NULL DEFAULT '' COMMENT 'PassWord',
    PRIMARY KEY (`user_id`),
    KEY           `idx_name` (`user_id`) COMMENT 'User id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User information table'