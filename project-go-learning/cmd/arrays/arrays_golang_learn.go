package main

import "fmt"

func main() {
	// 1 - Calling function fiverNumbersArray()
	fiverNumbersArray()

	// 2 - calling function numbersReverseSort()
	numbersReverseSort()

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

func numbersReverseSort() {
	fmt.Println("--------- Reversing sort of numbers in list -----------")

	var sizeArray int = 5
	var array [5]int

	// creating iteration loop for store data in list
	for i := 0; i < sizeArray; i++ {
		fmt.Printf("User, enter with in value index %d: ", i)
		fmt.Scan(&array[i])
	}

	fmt.Println("Printing reverse numbers")
	for i := len(array) - 1; i >= 0; i-- { // i := len(array) initialize i with last index element in array
		fmt.Println(array[i])
	}
}
