package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func TestRedisSet(conn redis.Conn) {
	_, err := conn.Do("set", "name", "tom")
	if err != nil {
		fmt.Println("set fail = ", err)
		return
	}

	val, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get fail = ", err)
		return
	}
	fmt.Println("the value of name = ", val)

}

func TestRedisHash(conn redis.Conn) {
	_, err := conn.Do("Hset", "stu1", "name", "jack")
	if err != nil {
		fmt.Println("hset stu1 name fail = ", err)
		return
	}

	_, err = conn.Do("Hset", "stu1", "age", 18)
	if err != nil {
		fmt.Println("hset stu1 age fail = ", err)
		return
	}

	name, err := redis.String(conn.Do("hget", "stu1", "name"))
	if err != nil {
		fmt.Println("hegt stu1 name fail = ", err)
	}
	age, err := redis.Int(conn.Do("hget", "stu1", "age"))
	if err != nil {
		fmt.Println("hegt stu1 age fail = ", err)
	}

	fmt.Printf("stu1 name = %s, age = %d\n", name, age)

}

func TsetRedisMHash(conn redis.Conn) {
	// 用hset也可以批量插入
	_, err := conn.Do("Hmset", "stu2", "name", "tom", "age", 20)
	if err != nil {
		fmt.Println("hset stu2 fail = ", err)
		return
	}

	rev, err := redis.Strings(conn.Do("Hmget", "stu2", "name", "age"))
	if err != nil {
		fmt.Println("hget stu2 fail = ", err)
	}

	for i, value := range rev {
		fmt.Printf("stu2[%d] = %s\n", i, value)
	}

}

func main() {
	conn, err := redis.Dial("tcp", "8.131.242.137:6379")
	if err != nil {
		fmt.Println("redis.Dial err = ", err)
		return
	}
	defer conn.Close()
	fmt.Println("redis connect success")

	// go中的redis所有操作和在命令行无差别，在conn.Do()函数中填入命令行的操作即可，区别在于是用逗号隔开

	// set测试
	TestRedisSet(conn)

	// hash测试
	TestRedisHash(conn)

	// 批量插入查询测试
	TsetRedisMHash(conn)
}
