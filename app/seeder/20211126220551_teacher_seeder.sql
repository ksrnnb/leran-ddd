-- +goose Up
-- +goose StatementBegin
INSERT INTO teachers (name) VALUES
-- 名前は以下サイトからランダムで生成
-- https://namegen.jp/
("鈴木 諒"),
("大塚 由香");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
