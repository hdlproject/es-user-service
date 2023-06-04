CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	balance BIGINT NOT NULL DEFAULT 0,
	date_created timestamptz NOT NULL,
	date_updated timestamptz NOT NULL,
	date_deleted timestamptz
);
