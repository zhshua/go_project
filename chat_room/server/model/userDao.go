package model

import (
	"encoding/json"
	"fmt"
	"go_project/chat_room/common/message"

	"github.com/garyburd/redigo/redis"
)

// 在服务器启动后会创建一个UserDao实例,
// 把它定义为全局变量，需要的时候直接用即可
var (
	MyUserDao *UserDao
)

// 定义一个UserDao结构体，完成对user结构体的各种操作
type UserDao struct {
	Pool *redis.Pool
}

// 创建UserDao实例
func NewUserDao(pool *redis.Pool) (dao *UserDao) {
	dao = &UserDao{
		Pool: pool,
	}
	return
}

// 通过给定id查询是否有这个用户
func (dao *UserDao) getUserById(conn redis.Conn, id int) (user *message.User, err error) {
	// 通过给定id去查询用户
	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXIST
			return
		}
	}

	// 对取到的json进行反序列化,得到user结构体
	user = &message.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return
}

// 登录判断
func (dao *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {

	// 从redis连接池中取出一个连接
	conn := dao.Pool.Get()
	defer conn.Close()

	user, err = dao.getUserById(conn, userId)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

// 注册判断
func (dao *UserDao) Register(user *message.User) (err error) {

	// 从redis连接池中取出一个连接
	conn := dao.Pool.Get()
	defer conn.Close()

	_, err = dao.getUserById(conn, user.UserId)
	// err = nil说明用户在redis里
	if err == nil {
		err = ERROR_USER_EXIST
		return
	}

	// 注册用户,先序列化user
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 存入数据库
	_, err = conn.Do("hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("conn.Do err = ", err)
		return
	}
	return
}
