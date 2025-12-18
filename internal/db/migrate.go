package db

import "database/sql"

// MustApplyMigrations создает таблицу, если её нет
func MustApplyMigrations(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS notes (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		content TEXT NOT NULL
	);`
	if _, err := db.Exec(query); err != nil {
		panic("failed to apply migrations: " + err.Error())
	}
}