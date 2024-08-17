package models

import (
	"database/sql"
	"errors"
	"main/forms"
)

/**
CREATE TABLE IF NOT EXISTS posts (
	id INT AUTO_INCREMENT PRIMARY KEY,
	user_id INT,
	title VARCHAR(255) NOT NULL,
	content TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id)
);
*/

type Post struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type PostModel struct{}

func (m PostModel) Create(post forms.CreatePostForm, db *sql.DB) error {
	// Create a new post
	// check if user exists
	check, err := db.Query("SELECT * FROM users WHERE id = ?", post.UserID)
	if err != nil {
		return err
	}

	if !check.Next() {
		return errors.New("User does not exist")
	}

	_, err = db.Query("INSERT INTO posts (user_id, title, content) VALUES (?, ?, ?)", post.UserID, post.Title, post.Content)

	if err != nil {
		return err
	}

	return nil
}
