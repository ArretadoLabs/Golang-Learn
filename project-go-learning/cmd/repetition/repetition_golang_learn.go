package main

import "fmt"

func main() {
	// 1 - calling function for print number between zero and ten
	numberBetweenZeroTen()

}

// functions auxiliar for calling em function "main()"

func numberBetweenZeroTen() {
	fmt.Println("Inputing a number between zero and ten")

	// Create variable with value for forcing enter in loop iteration
	number := 11

	for number < 0 || number > 10 {
		fmt.Println("The enter with a number: ")
		fmt.Scan(&number)
	}

	fmt.Printf("The number %v is valid\n", number)
}
