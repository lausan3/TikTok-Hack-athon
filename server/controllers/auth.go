package controllers

import (
	"database/sql"
	"main/forms"
	validator "main/infra/utils/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (u *AuthController) LoginUser(c *gin.Context, db *sql.DB) (token validator.Token, err error) {
	// get the request body and bind it to the LoginForm struct
	var loginForm forms.LoginForm

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		message := userForm.Register(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": message,
		})
		return token, err
	}

	token, err = userModel.Login(loginForm, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return token, err
	}

	c.JSON(http.StatusOK, token)

	return token, nil
}
