-- +goose Up
-- +goose StatementBegin
CREATE TABLE phnumbers (
    contact_id UUID NOT NULL,
    phnum VARCHAR,
    tag VARCHAR,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY(contact_id,phnum)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE phnumbers;
-- +goose StatementEnd
