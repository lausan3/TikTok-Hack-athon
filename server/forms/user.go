package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type UserForm struct{}

type RegisterForm struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginForm struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u UserForm) UserName(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		return "User name is required"
	default:
		return "User name is invalid"
	}
}

func (u UserForm) Password(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		return "Password is required"
	default:
		return "Password is invalid"
	}
}

func (u UserForm) Register(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Some fields are invalid"
		}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Field() {
			case "UserName":
				return u.UserName(e.Tag())
			case "Password":
				return u.Password(e.Tag())
			}
		}
	default:
		return "Invalid request"
	}

	return "Some fields are invalid"
}
