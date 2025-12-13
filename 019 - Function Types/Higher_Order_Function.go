package main

import "fmt"


func add (num1 int, num2 int) {
	fmt.Println(num1 + num2)
}

func processOperation (a int, b int, op func(p int, q int)) {
	op(a, b)
}

func main () {
	processOperation(2, 3, add)		// 5
}