package main

import "fmt"

func main() {
	var age int
	fmt.Println("please input age")
	fmt.Scan(&age)
	if age <= 0 {
		fmt.Println("未出生")
		return
	}
	if age <= 18 {
		fmt.Println("未成年")
		return
	}

}
