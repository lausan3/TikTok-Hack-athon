package forms

import (
	"encoding/json"
	"fmt"
	"main/infra/logger"

	"github.com/go-playground/validator/v10"
)

type PostForm struct{}

type CreatePostForm struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  int    `json:"user_id" binding:"required"`
}

func (p PostForm) Title(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		return "Title is required"
	default:
		return "Title is invalid"
	}
}

func (p PostForm) Content(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		return "Content is required"
	default:
		return "Content is invalid"
	}
}

func (p PostForm) UserID(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		return "User ID is required"
	default:
		return "User ID is invalid"
	}
}

func (o PostForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Some fields are invalid"
		}

		for _, e := range err.(validator.ValidationErrors) {
			switch e.Field() {
			case "Title":
				return o.Title(e.Tag())
			case "Content":
				return o.Content(e.Tag())
			case "UserID":
				return o.UserID(e.Tag())
			}
		}
	default:
		logger.Errorf(fmt.Sprintf("Error: %v", err))
		return "Invalid request"
	}

	return "Some fields are invalid"
}
