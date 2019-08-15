package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ping, err := client.Ping().Result()
	fmt.Println(ping, err)
	if err != nil {
		return nil
	}

	return client
}

func main() {
	client := RedisClient()
	if client == nil {
		fmt.Println("error")
		return
	}

	val, err := client.Get("mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key: ", val)

	set, err := client.Set("tsst", "fuck12", 10*time.Second).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(set)

	client.Close()
}
