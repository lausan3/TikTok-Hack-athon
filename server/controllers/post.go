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
	username := c.Param("username")

	if validatorErr := c.ShouldBindJSON(&postFormCreate); validatorErr != nil {
		message := postForm.Create(validatorErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": message,
		})
		return post, validatorErr
	}

	if validatorErr := postModel.Create(username, postFormCreate, db); validatorErr != nil {
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

func (p *PostController) GetPostsByUser(c *gin.Context, db *sql.DB) (posts []models.Post, err error) {
	username := c.Param("username")

	posts, err = postModel.GetPostsByUser(username, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return posts, err
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

	return posts, nil
}

func (p *PostController) GetAllPosts(c *gin.Context, db *sql.DB) (posts []models.Post, err error) {
	posts, err = postModel.GetAllPosts(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return posts, err
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

	return posts, nil
}
