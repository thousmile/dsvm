-- +goose Up

CREATE TABLE `cms_device` (
                              `device_id` bigint NOT NULL COMMENT '设备ID',
                              `device_name` varchar(255) NOT NULL COMMENT '设备名称',
                              `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 【0.禁用 1.正常 2.锁定 】',
                              `project_id` bigint NOT NULL COMMENT '项目ID',
                              `create_time` datetime NOT NULL COMMENT '创建时间',
                              `create_user` bigint NOT NULL COMMENT '创建人',
                              `last_update_time` datetime NOT NULL COMMENT '最后一次修改时间',
                              `last_update_user` bigint NOT NULL COMMENT '最后一次修改人',
                              PRIMARY KEY (`device_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='[ 项目 ] 设备表';

-- +goose Down
DROP TABLE cms_device;
