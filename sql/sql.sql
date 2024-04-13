create database rose;
use rose;

CREATE TABLE `user` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `create_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `user_name` varchar(255) NOT NULL COMMENT '用户名',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_uname` (`user_name`),
    KEY `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';