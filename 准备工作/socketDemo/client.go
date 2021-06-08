package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "172.30.12.32:8888")
	if err != nil {
		fmt.Println("dial err = ", err)
		return
	}

	for {
		var buf string
		fmt.Scan(&buf)
		if buf == "exit" {
			fmt.Println("client exit")
			return
		}

		_, err = conn.Write([]byte(buf))
		if err != nil {
			fmt.Println("conn.Write err = ", err)
		}

	}

}
