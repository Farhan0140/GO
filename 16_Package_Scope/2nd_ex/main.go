package main

import (
	"fmt"
	custompackage "math_mod/CustomPackage"
)

func main () {

	// custompackage.add(100, 400)	//  for use this i have to initialize module
	// after initializing

	fmt.Printf("%f, %T\n", custompackage.Pi, custompackage.Pi)	// 3.141590, float64

	custompackage.SayHello("Nadim")		//	Hi! Nadim
}