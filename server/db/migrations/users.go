package migrations

import (
	"database/sql"
	"main/infra/logger"
)

func MigrateUsersTable(db *sql.DB) {
	_, err := db.Query(
		`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_name VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		logger.Fatalf("Failed to create users table: %s", err)
	}
}
