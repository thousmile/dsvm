-- ------------------------------------------------------------------------------------ --
-- documentation https://pressly.github.io/goose/reference/annotations/
-- ------------------------------------------------------------------------------------ --

-- +goose Up

INSERT INTO t_user (username)
VALUES ('Tom');

INSERT INTO t_role (id, role_name)
VALUES (1, 'Admin');


-- ------------------------------------------------------------------------------------ --

-- +goose Down
DELETE FROM t_user;

DELETE FROM t_role WHERE id = 1;