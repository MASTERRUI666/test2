package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务器已启动，正在监听端口 8080...")

	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}
		fmt.Println("客户端已连接：", conn.RemoteAddr())
		// 处理客户端请求
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		// 读取客户端发送的数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		clientMsg := string(buf[:n])
		fmt.Println("收到客户端消息：", clientMsg)

		// 回复消息给客户端
		serverMsg := "服务器已收到消息：" + clientMsg
		_, err = conn.Write([]byte(serverMsg))
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}
