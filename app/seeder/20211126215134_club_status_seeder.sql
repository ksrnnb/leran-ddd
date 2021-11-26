-- +goose Up
-- +goose StatementBegin
INSERT INTO club_statuses (name) VALUES
("未承認"),
("承認済み");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
