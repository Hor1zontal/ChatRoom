package main

import (
	"fmt"
	"net"
	"server/process"
)

func init() {
	//redis初始化

	//常见userDao用于操作用户信息
	//全局唯一的UserDao实例：model.CurrentUserDao
}

func handle(conn net.Conn) {
	defer conn.Close()
	processor := process.Processor{Conn: conn}
	processor.MainProcess()
}


func main() {
	fmt.Printf("Server is already\n")

	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	defer listener.Close()

	if err != nil {
		fmt.Printf("Server error...\nerror is: %v", err)
	}

	for {
		fmt.Printf("Waiting for client...\n")

		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Server error...\nerror is: %v", err)
		}

		go handle(conn)
	}
}