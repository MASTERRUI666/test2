package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var text string
		fmt.Printf("请输入你要输入的内容:")
		fmt.Scan(&text)
		if text == "q" {
			conn.Close()
			break
		}
		conn.Write([]byte(text))

	}

}
