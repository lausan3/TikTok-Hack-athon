package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type Token struct {
	Token   string `json:"token"`
	Expires int64  `json:"expires"`
}

func GenerateToken(username string) (Token, error) {
	fmt.Printf("Token expiration: %v\n", viper.GetDuration("API_JWT_EXPIARY"))
	expirayTime := time.Now().Add(time.Hour * viper.GetDuration("API_JWT_EXPIARY")).Unix()

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = username
	claims["exp"] = expirayTime
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(viper.GetString("API_JWT_SECRET_KEY")))
	if err != nil {
		return Token{}, err
	}

	return Token{
		Token:   tokenString,
		Expires: expirayTime,
	}, nil
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(viper.GetString("API_JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		return err
	}

	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}