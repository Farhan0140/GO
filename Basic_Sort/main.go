package main

import (
	"fmt"
	"cmp"
	"slices" // OR sort
	"sort"
)


type Person struct {
	Name string 
	Age int
	Department string
	Salary float64
}


func main () {
	// Ascending Order
	arr1 := []int{4,1,7,10,66,1,2,3,5,110}
	slices.Sort(arr1)
	fmt.Println(arr1)	// [1 1 2 3 4 5 7 10 66 110]

	// Descending Order type 1
	slices.SortFunc(arr1, func(a, b int) int {
		return cmp.Compare(b, a)
	})
	fmt.Println(arr1)	// [110 66 10 7 5 4 3 2 1 1]

	// OR Descending Order type 2
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] > arr1[j]
	})
	fmt.Println(arr1)	// [110 66 10 7 5 4 3 2 1 1]


	// Sort Struct Ascending Order
	person1 := []Person{
		{Name: "Farhan", Age: 23, Department: "CSE", Salary: 100.00},
		{Name: "Nadim", Age: 25, Department: "BBA", Salary: 120.50},
		{Name: "Safa", Age: 20, Department: "CSE", Salary: 10000.00},
		{Name: "Tamim", Age: 18, Department: "CSE", Salary: 1200.00},
	}

	slices.SortFunc(person1, func(a, b Person) int {
		return cmp.Compare(a.Age, b.Age)
	})
	fmt.Println(person1)
	// [
	// 	{Tamim 18 CSE 1200} 
	// 	{Safa 20 CSE 10000} 
	// 	{Farhan 23 CSE 100} 
	// 	{Nadim 25 BBA 120.5}
	// ]

	
	// Descending Order
	slices.SortFunc(person1, func(a, b Person) int {
		return cmp.Compare(b.Age, a.Age)
	})
	fmt.Println(person1)
	// [
	// 	{Nadim 25 BBA 120.5} 
	// 	{Farhan 23 CSE 100} 
	// 	{Safa 20 CSE 10000} 
	// 	{Tamim 18 CSE 1200}
	// ]


	// using sort Ascending Order
	sort.Slice(person1, func(i, j int) bool {
		return person1[i].Salary < person1[j].Salary
	})
	fmt.Println(person1)
	// [
	// 	{Farhan 23 CSE 100} 
	// 	{Nadim 25 BBA 120.5} 
	// 	{Tamim 18 CSE 1200} 
	// 	{Safa 20 CSE 10000}
	// ]
	
	
	// using sort Descending Order
	sort.Slice(person1, func(i, j int) bool {
		return person1[i].Salary > person1[j].Salary
	})
	fmt.Println(person1)
	// [
	// 	{Safa 20 CSE 10000} 
	// 	{Tamim 18 CSE 1200} 
	// 	{Nadim 25 BBA 120.5} 
	// 	{Farhan 23 CSE 100}
	// ]
}