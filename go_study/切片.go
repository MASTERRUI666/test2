package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	s1 := []int{}
	s1 = arr[1:4]
	fmt.Println(s1)
	fmt.Printf("****************************** \n")
	slice := make([]int, 10)
	fmt.Println(slice)
}
