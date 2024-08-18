package models

import (
	"database/sql"
	"errors"
	"main/forms"
	"main/infra/utils"
)

type Post struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type PostResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
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

	_, err = db.Query("INSERT INTO posts (user_id, user_name, title, content) VALUES ((SELECT id FROM users WHERE user_name = ?), ?, ?, ?)", username, username, post.Title, post.Content)
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

	postsQuery, err := db.Query("SELECT * FROM posts WHERE user_name = ?", username)
	if err != nil {
		return nil, err
	}

	for postsQuery.Next() {
		var post Post
		err = postsQuery.Scan(&post.ID, &post.UserID, &post.UserName, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (m PostModel) GetAllPosts(db *sql.DB) ([]PostResponse, error) {
	var posts []PostResponse

	postsQuery, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	for postsQuery.Next() {
		var post PostResponse
		err = postsQuery.Scan(&post.ID, &post.UserID, &post.UserName, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
