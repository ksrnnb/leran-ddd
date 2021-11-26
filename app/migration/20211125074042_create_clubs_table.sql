-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS clubs (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'PK',
    name VARCHAR(255) NOT NULL COMMENT '部活名',
    status_id INT NOT NULL COMMENT '部活ステータス',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
    FOREIGN KEY (status_id) REFERENCES club_statuses (id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clubs;
-- +goose StatementEnd
