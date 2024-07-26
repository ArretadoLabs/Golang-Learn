package main

import "fmt"

func main() {
	//1 - Calling function fiverNumbersArray()
	fiverNumbersArray()

}

func fiverNumbersArray() {
	fmt.Print("-------- store and printing five numbers in array -----------\n")

	//creating array in Golang with 5 values
	var sizeArray int = 5
	var array [5]int

	//Loop for iteration and inserting data in array
	for i := 0; i < sizeArray; i++ {
		fmt.Printf("Insert a value in index %d: ", i)
		fmt.Scan(&array[i])
	}

	// Printing values store in array
	fmt.Println("The values inserted in array is: ")
	for i := 0; i < sizeArray; i++ {
		fmt.Printf("Value in index %d : %d \n", i, array[i])
	}

}
