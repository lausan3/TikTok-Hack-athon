package controllers

import (
	"database/sql"
	"encoding/json"
	"main/forms"
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type PostController struct{}

var postModel = new(models.PostModel)
var postForm = new(forms.PostForm)

func (p *PostController) CreatePost(c *gin.Context, db *sql.DB, redisClient *redis.Client) (post models.Post, err error) {
	var postFormCreate forms.CreatePostForm
	if validatorErr := c.ShouldBindJSON(&postFormCreate); validatorErr != nil {
		message := postForm.Create(validatorErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": message,
		})
		return post, validatorErr
	}

	post, err = postModel.Create(c, postFormCreate, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return post, err
	}

	id := strconv.Itoa(post.ID)

	postJSON, err := json.Marshal(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return post, err
	}

	// add post to redis
	err = redisClient.Set(c, "post:"+id, postJSON, 0).Err()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return post, err
	}

	// Create a notification
	notification := models.PostNotification{
		PostID:    post.ID,
		Type:      "new_post",
		UserName:  post.UserName,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
	}

	notificationJson, err := json.Marshal(notification)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return post, err
	}

	// Publish the notification to "notifications:all" channel
	err = redisClient.Publish(c, "notifications:all", notificationJson).Err()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return post, err
	}

	c.JSON(http.StatusOK, post)

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

func (p *PostController) GetAllPosts(c *gin.Context, db *sql.DB) (posts []models.PostResponse, err error) {
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
