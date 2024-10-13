package main

import "fmt"

func main() {
	var usermap = map[int]string{
		1: "ze",
		2: "zr",
		4: "",
	}
	usermap[1] = "zradla"
	fmt.Println(usermap)
	delete(usermap, 4)
	fmt.Println(usermap)

	/*var Zrmap map[string]string
	Zrmap["name"] = "AZ"
	fmt.Println(Zrmap)*/

}
