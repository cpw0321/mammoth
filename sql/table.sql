CREATE TABLE IF NOT EXISTS `route` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(64)  NOT NULL DEFAULT '' COMMENT '路由名称',
    `method` varchar(255)  NOT NULL DEFAULT '' COMMENT 'api HTTP请求方式',
    `path` varchar(255)  NOT NULL DEFAULT '' COMMENT 'api请求路径',
    `description` varchar(255)  NOT NULL DEFAULT '' COMMENT '描述',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='路由表';

CREATE TABLE IF NOT EXISTS `role` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(64)  NOT NULL DEFAULT '' COMMENT '角色名称',
    `status` int(3)  DEFAULT NULL COMMENT '状态 1-启用 2-禁用',
    `creator_id` int(11)  DEFAULT NULL COMMENT '创建者id',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

CREATE TABLE IF NOT EXISTS `user_role` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_id` int(11) NOT NULL COMMENT '用户id',
    `role_id` int(11) NOT NULL COMMENT '角色id',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关联表';

CREATE TABLE IF NOT EXISTS `user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64)  NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(64)  NOT NULL DEFAULT '' COMMENT '密码',
    `status` int(3)  DEFAULT NULL COMMENT '状态 1-启用 2-禁用',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment = '用户表';
