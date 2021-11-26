-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS club_student (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'PK',
    club_id INT NOT NULL COMMENT '部活ID',
    student_id INT NOT NULL COMMENT '生徒ID',
    FOREIGN KEY (club_id) REFERENCES clubs (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (student_id) REFERENCES students (id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS club_student;
-- +goose StatementEnd
