-- +goose Up

CREATE TABLE `cms_project` (
                               `project_id` bigint NOT NULL AUTO_INCREMENT,
                               `project_name` varchar(255) NOT NULL COMMENT '项目名称',
                               `linkman` varchar(255) NOT NULL COMMENT '联系人名称',
                               `contact_number` varchar(255) NOT NULL COMMENT '联系电话',
                               `area_code` bigint NOT NULL COMMENT '行政区域',
                               `address` varchar(255) NOT NULL COMMENT '联系地址，如：左右云创谷1栋A座',
                               `sort` bigint NOT NULL COMMENT '排序',
                               `password` varchar(255) NOT NULL COMMENT '项目密码，做一些危险操作时，使用',
                               `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 【0.禁用 1.正常 2.锁定 】',
                               `dept_id` bigint NOT NULL COMMENT '部门ID',
                               `create_time` datetime NOT NULL COMMENT '创建时间',
                               `create_user` bigint NOT NULL COMMENT '创建人',
                               `last_update_time` datetime NOT NULL COMMENT '最后一次修改时间',
                               `last_update_user` bigint NOT NULL COMMENT '最后一次修改人',
                               PRIMARY KEY (`project_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='[ 项目 ] 项目表';

-- +goose Down
DROP TABLE cms_project;
