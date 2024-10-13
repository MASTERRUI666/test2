package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	defer f(a)
	a++
	fmt.Println(a)

}

func f(s int) {
	fmt.Println("函数里面的a：", s)
}
