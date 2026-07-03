-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS patients (
	id TEXT PRIMARY KEY,
	nik TEXT NOT NULL UNIQUE,
	name TEXT NOT NULL,
	age INTEGER NOT NULL,
	gender TEXT NOT NULL,
	birth_date TEXT NOT NULL,
	address TEXT NOT NULL,
	phone TEXT NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS patients;
-- +goose StatementEnd
