package main

func main() {
	// Calling function biggerTwoNumbers
	biggerTwoNumbers(10, 20)

}

//section of functions for training

func biggerTwoNumbers(numberOne, numberTwo int) int {

	// var oneNumber int
	// var twoNumber int

	// fmt.Println("Number one: ")
	// fmt.Scan(&oneNumber)

	// fmt.Println("Number two: ")
	// fmt.Scan(&twoNumber)

	var result int

	if numberOne > numberTwo {
		result = numberOne
	} else {
		result = numberTwo
	}

	return result

}
