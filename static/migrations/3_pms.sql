-- +goose Up

CREATE TABLE `pms_role` (
                            `role_id` bigint NOT NULL AUTO_INCREMENT,
                            `role_name` varchar(255) NOT NULL COMMENT '角色名',
                            `sort` bigint NOT NULL COMMENT '排序',
                            `description` varchar(255) NOT NULL COMMENT '部门描述',
                            `create_time` datetime NOT NULL COMMENT '创建时间',
                            `create_user` bigint NOT NULL COMMENT '创建人',
                            `last_update_time` datetime NOT NULL COMMENT '最后一次修改时间',
                            `last_update_user` bigint NOT NULL COMMENT '最后一次修改人',
                            PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='[ 权限 ] 角色表';

-- +goose Down
DROP TABLE pms_role;
