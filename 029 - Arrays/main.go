package main

import "fmt"

func main () {
	var arr [5]int
	fmt.Println(arr)	// [0 0 0 0 0]

	arr[2] = 99
	fmt.Println(arr)	// [0 0 99 0 0]

	arr1 := [5]int {1,2,3,4,5}
	fmt.Println(arr1)	// [1 2 3 4 5]

	// Initialize only specific element
	arr2 := [10]int{1:99, 5:100, 7:3}
	fmt.Println(arr2)	// [0 99 0 0 0 100 0 3 0 0]

	// Length of an array
	fmt.Println( len(arr1 ))	// 5
	fmt.Println( len(arr2 ))	// 10


	for i:=0; i<len(arr2); i++ {
		fmt.Printf("%d ", arr2[i])		// 0 99 0 0 0 100 0 3 0 0
	}
	fmt.Println()


	
	fmt.Println("for arr1")
	for idx, val := range arr1 {
		fmt.Printf("index [%d] -> value [%d]\n", idx, val)
	}

	// for arr1
	// index [0] -> value [1]
	// index [1] -> value [2]
	// index [2] -> value [3]
	// index [3] -> value [4]
	// index [4] -> value [5]
}