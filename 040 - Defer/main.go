/* 
Example 1 ----------------------

package main

import "fmt"

func calculate () (result int) {
	fmt.Println("Inside calculate function 1: ", result)
	
	temp := func () {
		result += 10
		fmt.Println("Inside temp anonymous function: ", result)
	}

	defer temp()

	result = 50
	fmt.Println("Inside calculate function 2: ", result)
	
	return
}

func calculate_temp () int {
	result := 0

	fmt.Println("Inside calculate_temp function 1: ", result)
	
	temp := func () {
		result += 10
		fmt.Println("Inside temp anonymous function: ", result)
	}

	defer temp()

	result = 50
	fmt.Println("Inside calculate_temp function 2: ", result)
	
	return result
}

func main () {
	a := calculate()
	fmt.Println("Inside main 1: ", a)	// Inside main 1:  60

	b := calculate_temp()
	fmt.Println("Inside main 2: ", b)	// Inside main 2: 50
}

*/


package main

import "fmt"

func calculate() (result int) {
	fmt.Println("Inside calculate 1: ", result)

	show := func() {
		result += 10
		fmt.Println("Defer 1: ", result)
	}
	defer show()

	result = 5

	p := func() {
		fmt.Println("Defer 2: ", result)
	}
	defer p()

	defer fmt.Println(result)

	fmt.Println("Inside calculate 2: ", result)

	defer fmt.Println(result)
	
	return
}

func main () {
	a := calculate()
	fmt.Println("Inside main: ", a)
	/*
		Inside calculate 1:  0
		Inside calculate 2:  5
		5
		5
		Defer 2:  5
		Defer 1:  15
		Inside main:  15
	*/
}