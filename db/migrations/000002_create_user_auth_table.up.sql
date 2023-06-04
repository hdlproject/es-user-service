CREATE TABLE IF NOT EXISTS user_auths (
	user_id INT PRIMARY KEY,
	username VARCHAR(20) NOT NULL,
	password VARCHAR(100) NOT NULL,
	date_created timestamptz NOT NULL,
	date_updated timestamptz NOT NULL,
	date_deleted timestamptz,
	CONSTRAINT fk_user_auth_user FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS user_accesses (
	id BIGSERIAL PRIMARY KEY,
	user_id INT PRIMARY KEY,
	refresh_token VARCHAR(100) NOT NULL,
	access_token VARCHAR(100) NOT NULL,
	date_created timestamptz NOT NULL,
	date_updated timestamptz NOT NULL,
	date_deleted timestamptz,
	CONSTRAINT fk_user_access_user FOREIGN KEY (user_id) REFERENCES users (id)
);
