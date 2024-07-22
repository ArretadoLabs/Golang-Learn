package main

import (
	"fmt"
)

func main() {
	// 1 - calling function = two bigger numbers
	biggerTwoNumbers()
	// 2 - calling function = numberPositiveOrNegative()
	numberPositiveOrNegative()
	// 3 - calling function = letterGenderPeople()
	letterGenderPeople()
	// 4 - calling function = letterConsonantVogal()
	letterConsonantVogal()
	// 5 - calling function = StudentStatusCollege()
	StudentStatusCollege()
	// 6 - calling function = threeBigNumbers()
	threeBigNumbers()

}

func biggerTwoNumbers() {
	fmt.Println("================== Two biggers numbers ==================")
	var numberOne int
	var numberTwo int

	fmt.Println("Enter with number one: ")
	fmt.Scan(&numberOne)

	fmt.Println("Enter with number two: ")
	fmt.Scan(&numberTwo)

	if numberOne > numberTwo {
		fmt.Printf("The number one is %v is bigger\n", numberOne)
	} else {
		fmt.Printf("The number two is %v is bigger\n", numberTwo)
	}
}

func numberPositiveOrNegative() {
	fmt.Println("============= Number positive or negative ==============")
	var number int

	fmt.Println("The input number: ")
	fmt.Scan(&number)

	if number >= 0 {
		fmt.Printf("The number %v is positive\n", number)
	} else {
		fmt.Printf("The number %v is negative\n", number)
	}
}

func letterGenderPeople() {
	fmt.Println("=============== Letter for identify gender people ===============")

	//section of variables
	var letter string

	//Input data user
	fmt.Println("Enter with a letter anyway: ")
	fmt.Scan(&letter)

	//Structure decision
	if letter == "f" || letter == "F" {
		fmt.Println("The gender is: Female")
	} else if letter == "m" || letter == "M" {
		fmt.Println("The gender is: Male")
	} else {
		fmt.Println("Gender is invalid, try again!")
	}
}

func letterConsonantVogal() {
	fmt.Println("=============== Identify letter if vogal or consonante ===============")

	//section of variables
	var letter string

	//Input data user
	fmt.Println("Input a letter: ")
	fmt.Scan(&letter)

	// structure logical
	if letter != "a" && letter != "e" && letter != "i" && letter != "o" && letter != "u" {
		fmt.Println("You letter is consonant")
	} else {
		fmt.Println("You letter is vogal")
	}
}

func StudentStatusCollege() {
	fmt.Println("============== Printing situation of student college ================")

	// Section of variables
	var noteOne float32
	var noteTwo float32

	// Input data
	fmt.Println("The enter with note one: ")
	fmt.Scan(&noteOne)

	fmt.Println("The enter with note two: ")
	fmt.Scan(&noteTwo)

	averageNote := (noteOne + noteTwo) / 2

	// Structure logical
	if averageNote == 10 {
		fmt.Println("Approved with distinction")
	} else if averageNote < 10 && averageNote >= 7 {
		fmt.Println("Approved")
	} else if averageNote < 7 {
		fmt.Println("Reproved")
	}

}

func threeBigNumbers() {
	fmt.Println("============ Identify bigger of three numbers =============")

	// section variables
	var numberOne int
	var numberTwo int
	var numberThree int

	// input data user
	fmt.Println("Number one: ")
	fmt.Scan(&numberOne)

	fmt.Println("Number two: ")
	fmt.Scan(&numberTwo)

	fmt.Println("Number three: ")
	fmt.Scan(&numberThree)

	if numberOne > numberTwo && numberOne > numberThree {
		fmt.Printf("Number one is bigger, a number: %v", numberOne)
	} else if numberTwo > numberThree {
		fmt.Printf("number two is bigger, a number: %v", numberTwo)
	} else {
		fmt.Printf("Number three is bigger, a number: %v", numberThree)
	}
	fmt.Println("Testing function Max in module Math")

}
