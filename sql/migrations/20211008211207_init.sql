-- +goose Up
-- +goose StatementBegin
CREATE TABLE links
(
    short_link    varchar(10) PRIMARY KEY NOT NULL,
    original_link text                    NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS links;
-- +goose StatementEnd
