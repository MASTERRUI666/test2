package main

import "fmt"

func sayhello() {
	fmt.Println("hello world")
}
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num) // 输出: 1 2 3 4 5
	}

}
