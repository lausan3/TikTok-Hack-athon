package controllers

import (
	"database/sql"
	"main/forms"
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.UserModel)
var userForm = new(forms.UserForm)

func (u *UserController) RegisterUser(c *gin.Context, db *sql.DB) (user models.User, err error) {
	// get the request body and bind it to the RegisterForm struct
	var registerForm forms.RegisterForm

	if err := c.ShouldBindJSON(&registerForm); err != nil {
		message := userForm.Register(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": message,
		})
		return user, err
	}

	if err := userModel.Register(registerForm, db); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return user, err
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})

	return user, nil
}

// Delete a user
func (u *UserController) DeleteUser(c *gin.Context, db *sql.DB) (err error) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return err
	}

	if err := userModel.Delete(id, db); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})

	return nil
}
