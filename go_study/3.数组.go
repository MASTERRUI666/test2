package main

import "fmt"

func main() {
	var number1 [5]int
	fmt.Println(number1)
	/*********************/
	/*var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)*/
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	/*访问数组元素*/
	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("%d ", n[j])
	}
}
