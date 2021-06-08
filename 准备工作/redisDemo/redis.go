package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "8.131.242.137:6379")
	if err != nil {
		fmt.Println("redis.Dial err = ", err)
		return
	}
	fmt.Println("redis connect success", conn)
}
