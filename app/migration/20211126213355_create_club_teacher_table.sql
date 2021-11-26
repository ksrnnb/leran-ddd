-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS club_teacher (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'PK',
    club_id INT NOT NULL COMMENT '部活ID',
    teacher_id INT NOT NULL COMMENT '教員ID',
    FOREIGN KEY (club_id) REFERENCES clubs (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (teacher_id) REFERENCES teachers (id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS club_teacher;
-- +goose StatementEnd
