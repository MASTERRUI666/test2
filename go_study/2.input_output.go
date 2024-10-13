package main

import "fmt"

func main() {
	fmt.Print("请输入名字：")
	var name string
	fmt.Scanln(&name)
	fmt.Println(name)
}
