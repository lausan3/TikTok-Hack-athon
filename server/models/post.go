package models

import (
	"database/sql"
	"errors"
	"main/forms"
	"main/infra/utils"
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

func (m PostModel) Create(username string, post forms.CreatePostForm, db *sql.DB) error {
	check, err := utils.CheckIfUserExists(username, db)
	if err != nil {
		return err
	}

	if !check {
		return errors.New("User does not exist")
	}

	_, err = db.Query("INSERT INTO posts (user_id, title, content) VALUES ((SELECT id FROM users WHERE user_name = ?), ?, ?)", username, post.Title, post.Content)
	if err != nil {
		return err
	}

	return nil
}

func (m PostModel) GetPostsByUser(username string, db *sql.DB) ([]Post, error) {
	// check if user exists
	check, err := utils.CheckIfUserExists(username, db)
	if err != nil {
		return nil, err
	}

	if !check {
		return nil, errors.New("User does not exist")
	}

	var posts []Post

	postsQuery, err := db.Query("SELECT * FROM posts WHERE user_id = (SELECT id FROM users WHERE user_name = ?)", username)
	if err != nil {
		return nil, err
	}

	for postsQuery.Next() {
		var post Post
		err = postsQuery.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (m PostModel) GetAllPosts(db *sql.DB) ([]Post, error) {
	var posts []Post

	postsQuery, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	for postsQuery.Next() {
		var post Post
		err = postsQuery.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
