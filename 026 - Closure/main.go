package main

import "fmt"

const a = 10
var p = 100


func Outer() func() {
	money := 100
	age := 30

	fmt.Println("Age: ", age)

	show := func () {
		money = money + a + p
		fmt.Println("Money: ", money)
	}

	return show
}

func Call() {
	inc1 := Outer()
	inc1()
	inc1()

	inc2 := Outer()
	inc2()
	inc2()
}

func main () {
	Call()
}

func init() {
	fmt.Println("--- From init ---")
}