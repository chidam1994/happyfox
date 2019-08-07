-- +goose Up
-- +goose StatementBegin
CREATE TABLE contacts (
    id UUID NOT NULL,
    name VARCHAR UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE contacts;
-- +goose StatementEnd
