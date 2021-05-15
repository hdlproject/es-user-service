CREATE TABLE IF NOT EXISTS users (
	id serial primary key,
	balance bigint not null default 0,
	date_created timestamptz not null,
	date_updated timestamptz not null,
	date_deleted timestamptz
);
