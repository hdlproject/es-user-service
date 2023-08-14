CREATE EXTENSION postgis;

CREATE TABLE IF NOT EXISTS user_locations (
	user_id INT PRIMARY KEY,
	geo geometry,
	date_created timestamptz NOT NULL,
	date_updated timestamptz NOT NULL,
	date_deleted timestamptz,
	CONSTRAINT fk_user_location_user FOREIGN KEY (user_id) REFERENCES users (id)
);
