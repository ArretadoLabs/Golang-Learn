package main

//Importing methods of library fmt
import (
	"fmt"
)

func main() {
	// 1 - Calling function = helloWorld()
	helloWorld()
	// 2 - Calling function = printNumber()
	printNumber()
	// 3 - calling function = sumTwoNumbers()
	sumTwoNumbers()
	// 4 - calling function = fourNotesAverage()
	fourNotesAverage()
	// 5 - calling function = converterSizeMetersCentimeters()
	converterSizeMetersCentimeters()

}

// =========================================== section of functinos for calling in main function ===========================================

// Exercise 1
func helloWorld() {
	fmt.Println("Hello World, Francisco!!")
}

// Exercise 2
func printNumber() {
	var number int
	fmt.Println("Enter with your number: ")
	fmt.Scan(&number)

	fmt.Printf("Your number is %v\n", number)
}

// Exercise 3
func sumTwoNumbers() {
	var numberOne int
	var numberTwo int

	fmt.Println("Enter with a number one: ")
	fmt.Scan(&numberOne)
	fmt.Println("Enter with a number two: ")
	fmt.Scan(&numberTwo)

	var sumTwoNumbers = numberOne + numberTwo

	fmt.Printf("The sum of numbers is: %v\n", sumTwoNumbers)
}

// Exercise 4
func fourNotesAverage() {
	var noteOne uint
	var noteTwo uint
	var noteThree uint
	var noteFour uint
	var averageNotes uint

	//Input data user and storing in variable
	fmt.Println("Note one: ")
	fmt.Scan(&noteOne)

	fmt.Println("Note two: ")
	fmt.Scan(&noteTwo)

	fmt.Println("Note three: ")
	fmt.Scan(&noteThree)

	fmt.Println("Note four: ")
	fmt.Scan(&noteFour)

	// calculating average notes and divided four
	averageNotes = (noteOne + noteTwo + noteThree + noteFour) / 4

	fmt.Printf("The you average notes is: %v\n", averageNotes)

}

// Exercise 5
func converterSizeMetersCentimeters() {
	var sizeMeters float32

	fmt.Println("What size in meters: ")
	fmt.Scan(&sizeMeters)

	var converterSizeInCentimeters = sizeMeters * 100

	fmt.Printf("The size %v meters converted in centimeters is: %v cm(s)\n", sizeMeters, converterSizeInCentimeters)

}
