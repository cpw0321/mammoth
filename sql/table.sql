
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64)  NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(64)  NOT NULL DEFAULT '' COMMENT '密码',
    `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
    `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';