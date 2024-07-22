package main

//Importing methods of library fmt
import (
	"fmt"
	"math"
)

/*
 */
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
	// 6 - calling function = areaCircle()
	areaCircle()
	// 7 - calling function = areaSquareExercise()
	areaSquareExercise()
	// 8 - calling function = salaryBrute()
	salaryBrute()
	// 9 - calling function = converterTemperature()
	converterTemperature()
	// 10 - calling function = numbersOperations()
	numbersOperations()
	// 11 - calling function = weightIdealPeople()
	weightIdealPeople()

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

func areaCircle() {
	const PI = 3.14
	var raidCircle float64

	fmt.Println("Enter with value of raid circle: ")
	fmt.Scan(&raidCircle)

	areaCircle := PI * math.Pow(raidCircle, 2)

	fmt.Printf("The area of circle is: %v\n", areaCircle)
}

func areaSquareExercise() {
	var sideSquare float32

	fmt.Println("The input value of side square: ")
	fmt.Scan(&sideSquare)

	//formula for calculate area Square
	areaSquare := sideSquare * sideSquare

	fmt.Printf("The area of square is: %v M²\n", areaSquare)
	fmt.Printf("The value double of area square is: %v M²\n", areaSquare*2)
}

func salaryBrute() {
	var hoursWorking uint
	var moneyHour uint

	fmt.Println("How many hours working in month: ")
	fmt.Scan(&hoursWorking)

	fmt.Println("How much you money for hour work: ")
	fmt.Scan(&moneyHour)

	bruteSalary := hoursWorking * moneyHour

	fmt.Printf("You salary brute in month is: R$ %v\n", bruteSalary)

}

func converterTemperature() {
	var fahrenheit float32
	fmt.Println("Enter with you temperature in Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	var temperatureConverter float32 = 5 * ((fahrenheit - 32) / 9)
	fmt.Printf("The temperature %v for converter temperature %v C°\n", fahrenheit, temperatureConverter)
}

func numbersOperations() {
	fmt.Println("============ Operation with three numbers ============")
	var numberOne int
	var numberTwo int
	var numberThree float64
	// Section of user data input
	fmt.Println("Number one: ")
	fmt.Scan(&numberOne)

	fmt.Println("Number two: ")
	fmt.Scan(&numberTwo)

	fmt.Println("Number three: ")
	fmt.Scan(&numberThree)

	fmt.Printf("Value of operation one is: %v\n", (numberOne*2)*(numberTwo/2))
	fmt.Printf("Value of operation two is: %v\n", (numberOne*3)+int(numberThree))
	fmt.Printf("Value of operation three is: %v\n", math.Pow((numberThree), 3))

}

func weightIdealPeople() {
	var heightPeople float32

	fmt.Println("Enter with you height: ")
	fmt.Scan(&heightPeople)

	// formula for calculate weight ideal in people
	var idealWeightPeople = (72.7 * heightPeople) - 58

	fmt.Printf("The people with height %v cm(s) have weight ideal is: %v\n", heightPeople, idealWeightPeople)
}
