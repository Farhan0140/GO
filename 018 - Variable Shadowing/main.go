package main

import "fmt"

func main () {
	x := 100

	if 10 < 20 {
		x := 17
		fmt.Println(x)	// 17
	}

	fmt.Println(x)	// 100
}