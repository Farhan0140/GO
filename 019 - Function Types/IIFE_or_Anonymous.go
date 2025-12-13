package main

import "fmt"

func main () {

	// Anonymous function
	func(num1 int, num2 int) {
		fmt.Println(num1 + num2)
	} (20, 50) // IIFE - Immediately Invoke Function Expression
}