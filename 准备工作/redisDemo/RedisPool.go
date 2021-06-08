package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 连接池中的最大连接数
		MaxActive:   0,   // 同时可以有几个连接
		IdleTimeout: 100, // 超时时间，100s连接没活动后连接会被放回连接池中
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "8.131.242.137:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "test", "haha")
	if err != nil {
		fmt.Println("conn.Do err = ", err)
		return
	}

	value, err := redis.String(conn.Do("get", "test"))
	if err != nil {
		fmt.Println("conn.Do get err = ", err)
	}

	fmt.Printf("value = %s", value)

	pool.Close() // 关闭连接池，连接池关闭后不可再从池中取连接
}
