package main

import (
	"fmt"
	"go_code/go_study/version"
)

func hello() {
	//fmt.Println(age)
	fmt.Println("hello world")
}

var age = 12

// age := 12
func main() {
	//首字母大写的变量、函数、方法、属性可在包外访问
	fmt.Println(version.Version)
	//fmt.Println(version.zr)
	//先声明
	var name string
	//在赋值
	name = "zr"

	//声明并赋值
	var name1 string = "zr"
	var name2 = "zr"
	//短变量声明
	name3 := "zr"
	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(name2)

	fmt.Println(name3)
	hello()
	var a1, a2 = 1, 2
	fmt.Println(a1, a2)
	var (
		s1 string = "1"
		s2 string = "2"
	)
	fmt.Println(s1, s2)
}
