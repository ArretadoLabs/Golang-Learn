package main

import "fmt"

func main() {
	// 1 - Calling function fiverNumbersArray()
	fiverNumbersArray()

	// 2 - calling function numbersReverseSort()
	numbersReverseSort()

	// 3 - calling function notesAverage()
	notesAverage()

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

func notesAverage() {
	fmt.Println("-------- Store four notes and printing note average ---------")

	// Section of variables
	var notes [4]float64
	var sumNotes float64

	//Reading notes input
	for i := 0; i < len(notes); i++ {
		fmt.Printf("Input your note: %d", i+1)
		fmt.Scan(&notes[i])
		sumNotes += notes[i]
	}

	// Calculating average of notes
	average := sumNotes / float64(len(notes))
	fmt.Printf("The average of notes: %.2f \n", average)
}
