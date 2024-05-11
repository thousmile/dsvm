-- ------------------------------------------------------------------------------------ --
-- documentation https://pressly.github.io/goose/reference/annotations/
-- ------------------------------------------------------------------------------------ --

-- +goose Up

-- +goose StatementBegin
CREATE FUNCTION add_numbers (a INT, b INT)
    RETURNS INT
BEGIN
    DECLARE result INT;
    SET result = a + b;
RETURN result;
END;
-- +goose StatementEnd

-- ------------------------------------------------------------------------------------ --

-- +goose Down
DROP FUNCTION IF EXISTS add_numbers;