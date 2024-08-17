package migrations

import (
	"database/sql"
	"main/infra/logger"
)

func MigratePostsTable(db *sql.DB) error {
	_, err := db.Query(
		`
		CREATE TABLE IF NOT EXISTS posts (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT,
			title VARCHAR(255) NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`)

	if err != nil {
		logger.Fatalf("Failed to create posts table: %s", err)
	}

	return nil
}
