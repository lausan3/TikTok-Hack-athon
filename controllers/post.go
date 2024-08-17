package controllers

import (
	"database/sql"
	"fmt"
	"main/forms"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

var postModel = new(models.PostModel)
var postForm = new(forms.PostForm)

func (p *PostController) CreatePost(c *gin.Context, db *sql.DB) (post models.Post, err error) {
	var postFormCreate forms.CreatePostForm

	if validatorErr := c.ShouldBindJSON(&postFormCreate); validatorErr != nil {
		message := postForm.Create(validatorErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": message,
		})
		return post, validatorErr
	}

	if validatorErr := postModel.Create(postFormCreate, db); validatorErr != nil {
		fmt.Println(validatorErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validatorErr.Error(),
		})
		return post, validatorErr
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post created successfully",
	})

	return post, nil
}
