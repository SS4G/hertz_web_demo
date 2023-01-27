CREATE DATABASE hertz_demo;
use hertz_demo;
CREATE TABLE IF NOT EXISTS `doc`
(
    `doc_id`     bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'DocId',
    `image_url`  varchar(1024)  DEFAULT 'https://globaleasterninvestment.com/wp-content/uploads/2020/11/blog-baidu.jpg' COMMENT 'ImageUrl',
    `title`      varchar(256) NOT NULL DEFAULT 'untitled' COMMENT 'Title',
    `text`       varchar(4096) NOT NULL DEFAULT 'default text' COMMENT 'Text',
    PRIMARY KEY (`doc_id`),
    KEY           `idx_name` (`doc_id`) COMMENT 'User id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Doc information table';
