package process

import "fmt"

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

var userMgr *UserMgr

// 初始化userMgr
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对OnlineUsers的添加操作
func (um *UserMgr) AddOnlineUsers(up *UserProcess) {
	um.onlineUsers[up.UserId] = up
}

// 完成对OnlineUsers的删除操作
func (um *UserMgr) DeleteOnlineUsers(userId int) {
	delete(um.onlineUsers, userId)
}

// 返回id对应的UserProcess
func (um *UserMgr) GetNolineUsersById(UserId int) (up *UserProcess, err error) {
	up, ok := um.onlineUsers[UserId]
	if !ok {
		// 没找到, 说明查找的用户不在线
		err = fmt.Errorf("用户 %d 不存在！", UserId)
	}
	return
}

// 返回所有OnlineUsers
func (um *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return um.onlineUsers
}
