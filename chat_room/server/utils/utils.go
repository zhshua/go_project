/* 存放通用工具方法 */

package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_project/chat_room/common/message"
	"net"
)

// 定义传输数据类
type Transfer struct {
	Conn net.Conn //连接
	Buf  []byte   //缓冲区
}

func (tf *Transfer) ReadPkg() (msg message.Message, err error) {
	// 1. conn.Read返回的是读取到的字节数
	// 2. conn,Read会把读取的[]byte结果存储到传进的buf切片中
	_, err = tf.Conn.Read(tf.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read head err = ", err)
		return
	}

	// 把字节序转换成uint32类型的数字
	pkgLen := binary.BigEndian.Uint32(tf.Buf[:4])
	_, err = tf.Conn.Read(tf.Buf[:pkgLen])
	if err != nil {
		fmt.Println("conn.Read body err = ", err)
		return
	}

	// 反序列化结构体
	err = json.Unmarshal(tf.Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
	}
	return
}

func (tf *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度给对方
	var buf [4]byte // 一个uint32类型是4个字节
	pkgLen := uint32(len(data))
	// 先发送消息的长度,发送前需要用下面函数把数字类型转换成字节序, 转换结果存在buf
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := tf.Conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write head err = ", err)
		return
	}

	// 发送消息体
	n, err = tf.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write body err = ", err)
		return
	}
	return
}
