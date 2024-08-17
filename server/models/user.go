package models

import (
	"database/sql"
	"errors"
	"main/forms"
)

/**
CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	user_name VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
*/

type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserModel struct{}

func (m UserModel) Register(user forms.RegisterForm, db *sql.DB) error {
	// Check if user already exists
	check, err := db.Query("SELECT * FROM users WHERE user_name = ?", user.UserName)

	if err != nil {
		return err
	}

	if check.Next() {
		return errors.New("User already exists")
	}

	// Create a new user
	_, err = db.Query("INSERT INTO users (user_name, password) VALUES (?, ?)", user.UserName, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (m UserModel) Delete(id int, db *sql.DB) error {
	_, err := db.Query("DELETE FROM users WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
