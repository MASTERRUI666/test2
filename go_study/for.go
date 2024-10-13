package main

import (
	"fmt"
	"time"
)

func main() {
	/*var sum int
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)*/
	for true {
		fmt.Println(time.Now())
		time.Sleep(1 * time.Second)
	}
}
