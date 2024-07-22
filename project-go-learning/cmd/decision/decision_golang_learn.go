package main

import "fmt"

func main() {
	// 1 - calling function: two bigger numbers
	biggerTwoNumbers()
	// 2 - calling function: numberPositiveOrNegative()
	numberPositiveOrNegative()

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
		fmt.Printf("The number %v is negative", number)
	}
}
