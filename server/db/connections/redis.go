package connections

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func OpenRedisConnection() *redis.Client {
	host := viper.GetString("REDIS_HOST")
	password := viper.GetString("REDIS_PASSWORD")
	port := viper.GetString("REDIS_PORT")

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	ctx := context.Background()

	//send SET command
	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	//send GET command and print the value
	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}

	println(val)

	return client
}
