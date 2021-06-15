package main

import (
	"go_project/chat_room/server/model"
	"time"

	"github.com/garyburd/redigo/redis"
)

// 定义一个Redis Pool全局变量
var pool *redis.Pool

func initPool(addr string, maxIdle, maxActive int, idleTimeout time.Duration) {

	pool = &redis.Pool{
		MaxIdle:     maxIdle,     // 连接池中的最大连接数
		MaxActive:   maxActive,   // 同时可以有几个连接
		IdleTimeout: idleTimeout, // 超时时间，300s连接没活动后连接会被放回连接池中
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}

// 初始化创建一个UserDao实例
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
