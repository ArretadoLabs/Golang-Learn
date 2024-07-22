package main

import "fmt"

func main() {
	// 1 - calling function = two bigger numbers
	biggerTwoNumbers()
	// 2 - calling function = numberPositiveOrNegative()
	numberPositiveOrNegative()
	// 3 - calling function = letterGenderPeople()
	letterGenderPeople()

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
