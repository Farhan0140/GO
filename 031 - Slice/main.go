package main

import "fmt"



func Change_Slice( p []int ) []int {
	p[0] = 100
	p = append(p, 99)

	return p
}

func main () {
	x := []int{1,2,3,4,5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := Change_Slice(a)

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(x[0:8])
}