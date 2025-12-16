package main

import "fmt"


type Person struct {
	Name string
	Age int
	Salary float32
	Dept string
}


func main () {
	// Type-1 instantiation of an instance
	var person1 Person

	// Type-2 instantiation of an instance and initialization
	person2 := Person {
		Name: "Farhan",
		Age: 23,
		Salary: 100.00,
		Dept: "CSE",
	}


	// Initialization

	person1.Name = "Nadim"
	person1.Age = 23
	person1.Salary = 300.00
	person1.Dept = "EEE"


	// or
	var person3 Person
	person3 = Person{
		Name: "Sadi",
		Age: 27,
		Salary: 100000.50,
		Dept: "CSE",
	}

	// Output 1
	fmt.Println(person1)
	fmt.Println()
	
	// Output 2
	fmt.Printf("%v", person2)
	fmt.Println()
	
	// Output 3 with field names
	fmt.Printf("%+v\n", person3)
	fmt.Println()

	// Output 4 Go-syntax representation
	fmt.Printf("%#v\n", person3)
	fmt.Println()


	// Output 4
	fmt.Printf("Name: %s\nAge: %d\nSalary: %.2f\nDepartment: %s", 
				person1.Name, 
				person1.Age, 
				person1.Salary, 
				person1.Dept,
			  )

	



}