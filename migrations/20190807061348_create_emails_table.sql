-- +goose Up
-- +goose StatementBegin
CREATE TABLE emails (
    contact_id UUID NOT NULL,
    email_id VARCHAR,
    tag VARCHAR,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY(contact_id,email_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE emails;
-- +goose StatementEnd
