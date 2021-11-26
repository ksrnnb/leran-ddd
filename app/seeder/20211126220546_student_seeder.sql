-- +goose Up
-- +goose StatementBegin
INSERT INTO students (name) VALUES
-- 名前は以下サイトからランダムで生成
-- https://namegen.jp/
("伊藤 徹"),
("高橋 裕司"),
("吉田 亮介"),
("鈴木 紀子"),
("中村 千恵"),
("村上 詩織");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
