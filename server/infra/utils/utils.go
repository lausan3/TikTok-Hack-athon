package utils

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
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

func GeneratePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
