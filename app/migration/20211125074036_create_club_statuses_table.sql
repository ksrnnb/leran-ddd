-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS club_statuses (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'PK',
    name VARCHAR(255) NOT NULL COMMENT 'ステータス名'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS club_statuses;
-- +goose StatementEnd
