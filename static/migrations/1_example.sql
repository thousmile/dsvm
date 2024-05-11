-- ------------------------------------------------------------------------------------ --
-- documentation https://pressly.github.io/goose/reference/annotations/
-- ------------------------------------------------------------------------------------ --

-- +goose Up
CREATE TABLE `t_user`
(
    `id`       int          NOT NULL AUTO_INCREMENT COMMENT 'user id',
    `username` varchar(255) NOT NULL COMMENT 'user name',
    PRIMARY KEY (`id`)
) COMMENT='user';


CREATE TABLE `t_role`
(
    `id`        int          NOT NULL COMMENT 'role id',
    `role_name` varchar(255) NOT NULL COMMENT 'role name',
    PRIMARY KEY (`id`)
) COMMENT='role';


-- ------------------------------------------------------------------------------------ --

-- +goose Down
DROP TABLE t_user;
DROP TABLE t_role;
