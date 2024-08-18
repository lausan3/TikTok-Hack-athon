package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"main/infra/logger"
	validator "main/infra/utils/validators"
	"main/models"
	"main/routers/middleware"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

func RegisterWebSocketRoutes(ctx context.Context, upgrader websocket.Upgrader, router *gin.Engine, redisClient *redis.Client) {
	connections := make(map[string]*websocket.Conn)
	var connectionsMutex sync.RWMutex

	router.GET("/ws", middleware.JwtAuthMiddleware(), func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			logger.Errorf(err.Error())
			return
		}

		userName, err := validator.ExtractTokenMetadata(c)
		fmt.Println("Username", userName)

		if err != nil {
			logger.Errorf(err.Error())
			conn.Close()
			return
		}

		connectionsMutex.Lock()
		connections[userName] = conn
		connectionsMutex.Unlock()

		// add the connected user to the set of connected clients
		check := redisClient.SAdd(ctx, "connectedClients", userName)

		if check.Err() != nil {
			logger.Errorf(check.Err().Error())
			conn.Close()
			return
		}

		defer func() {
			conn.Close()
			connectionsMutex.Lock()
			delete(connections, userName)
			connectionsMutex.Unlock()
			redisClient.SRem(ctx, "connectedClients", userName)
		}()

		// subscribe to the "all" notification channel
		pubsub := redisClient.Subscribe(ctx, "notifications:all")
		// defer pubsub.Close()

		// listen for new messages being published to the "all" channel
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				logger.Errorf(err.Error())
				return
			}

			var notification models.PostNotification
			err = json.Unmarshal([]byte(msg.Payload), &notification)

			if err != nil {
				logger.Errorf(err.Error())
				return
			}

			if notification.Type == "new_post" {
				post, err := getPost(notification.PostID, ctx, redisClient)

				if err != nil {
					logger.Errorf(err.Error())
					continue
				}

				if userName != post.UserName {
					err = conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
					if err != nil {
						logger.Errorf(err.Error())
						return
					}
				}
			}
		}
	})
}

func getPost(postID int, ctx context.Context, redisClient *redis.Client) (models.Post, error) {
	var post models.Post
	id := strconv.Itoa(postID)
	postJSON, err := redisClient.Get(ctx, "post:"+id).Bytes()
	if err != nil {
		return post, err
	}
	err = json.Unmarshal(postJSON, &post)
	return post, err
}
