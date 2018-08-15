package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	result, error := client.Ping().Result()
	fmt.Println(result, error) // PONG <nil>

	set_result, set_error := client.Set("key", "value", 0).Result()
	fmt.Println(set_result, set_error)
	if set_error != nil {
		panic(set_error)
	}

	get_result, get_error := client.Get("key").Result()
	if get_error != nil {
		panic(get_error)
	}

	if get_error == redis.Nil {
		fmt.Println("key does not exist")
	} else if get_error != nil {
		panic(get_error)
	} else {
		fmt.Println(get_result, get_error)
	}
}
