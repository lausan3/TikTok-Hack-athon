package utils

import (
	"database/sql"
)

func CheckIfUserExists(userName string, db *sql.DB) (bool, error) {
	check, err := db.Query("SELECT * FROM users WHERE user_name = ?", userName)

	if err != nil {
		return false, err
	}

	if check.Next() {
		return true, nil
	}

	return false, nil
}
