-- +migrate Up
-- +migrate StatementBegin

ALTER TABLE transactions ALTER expired_at TYPE timestamptz;
ALTER TABLE transactions ADD stock_retreived BOOLEAN DEFAULT FALSE;

-- +migrate StatementEnd