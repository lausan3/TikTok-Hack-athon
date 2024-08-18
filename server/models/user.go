package models

import (
	"database/sql"
	"errors"
	"main/forms"
	utils "main/infra/utils"
	validator "main/infra/utils/validator"
)

type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserModel struct{}

func (m UserModel) Register(user forms.RegisterForm, db *sql.DB) error {
	// check if user exists
	check, err := utils.CheckIfUserExists(user.UserName, db)

	if err != nil {
		return err
	}

	if check {
		return errors.New("User already exists")
	}

	// hash password
	hashedPassword, err := validator.GeneratePasswordHash(user.Password)

	if err != nil {
		return err
	}

	// Create a new user
	_, err = db.Query("INSERT INTO users (user_name, password) VALUES (?, ?)", user.UserName, hashedPassword)

	if err != nil {
		return err
	}

	return nil
}

func (m UserModel) Get(userName string, db *sql.DB) (User, error) {
	var user User

	check, err := utils.CheckIfUserExists(userName, db)

	if err != nil {
		return user, err
	}

	if !check {
		return user, errors.New("User does not exist")
	}

	// get user
	err = db.QueryRow("SELECT * FROM users WHERE user_name = ?", userName).Scan(&user.ID, &user.UserName, &user.Password, &user.CreatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (m UserModel) Delete(userName string, db *sql.DB) error {
	// check if user exists
	check, err := utils.CheckIfUserExists(userName, db)

	if err != nil {
		return err
	}

	if !check {
		return errors.New("User does not exist")
	}

	// delete all posts by user
	_, err = db.Query("DELETE FROM posts WHERE user_id = (SELECT id FROM users WHERE user_name = ?)", userName)

	if err != nil {
		return err
	}

	// delete user
	_, err = db.Query("DELETE FROM users WHERE user_name = ?", userName)

	if err != nil {
		return err
	}

	return nil
}

func (m UserModel) Login(user forms.LoginForm, db *sql.DB) (validator.Token, error) {
	var token validator.Token

	check, err := utils.CheckIfUserExists(user.UserName, db)

	if err != nil {
		return token, err
	}

	if !check {
		return token, errors.New("User does not exist")
	}

	var hashedPassword string

	err = db.QueryRow("SELECT password FROM users WHERE user_name = ?", user.UserName).Scan(&hashedPassword)

	if err != nil {
		return token, err
	}

	err = validator.VerifyPassword(user.Password, hashedPassword)

	if err != nil {
		return token, errors.New("Invalid password")
	}

	token, err = validator.GenerateToken(user.UserName)

	if err != nil {
		return token, err
	}

	return token, nil
}
