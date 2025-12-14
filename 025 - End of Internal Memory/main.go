package main

import "fmt"

const a = 10
var p = 100


func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}

	add(5, 7)
	add(a, p)
}

func main () {
	call()

	fmt.Println(a)
}

func init() {
	fmt.Println("From init function")
}