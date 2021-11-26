-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS students (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'PK',
    name VARCHAR(255) NOT NULL COMMENT '生徒名'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS students;
-- +goose StatementEnd
