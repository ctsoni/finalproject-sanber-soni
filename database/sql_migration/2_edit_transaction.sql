-- +migrate Up
-- +migrate StatementBegin

ALTER TABLE transactions ADD expired_at timestamp;

-- +migrate StatementEnd