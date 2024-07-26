package main

import "fmt"

func main() {
	//1 - Calling function fiverNumbersArray()
	fiverNumbersArray()

}

func fiverNumbersArray() {
	fmt.Print("-------- store and printing five numbers in array -----------\n")

	//Section of variables and arrays
	var sizeArray int

	fmt.Println("What size of you array: ")
	fmt.Scan(&sizeArray)

	array := make([]int, sizeArray)

	fmt.Println("Input you number/element: ")
	for i := 0; i < sizeArray; i++ {
		fmt.Scan(&array[i])
	}

	// printing elements inserting in array
	fmt.Println("The elements in array is: ")
	for _, v := range array {
		fmt.Println(v)
	}

}
