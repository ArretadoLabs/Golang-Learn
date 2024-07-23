package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1 - calling function = numberBetweenZeroTen()
	numberBetweenZeroTen()
	// 2 - calling function = nameLoginAndPasswordNotEquals()
	nameLoginAndPasswordNotEquals()

}

func numberBetweenZeroTen() {
	fmt.Println("------------ Printing value number valid (between zero or ten) ---------------")
	var number int

	// input data user
	fmt.Println("Enter with a number valid: ")
	fmt.Scan(&number)

	// loop iteration for valid number
	for number < 0 || number > 10 {
		fmt.Println("Number invalid, try again!")
		fmt.Scan(&number)
	}

	// printing value final number valid
	fmt.Printf("Number is: %v\n", number)
}

func nameLoginAndPasswordNotEquals() {
	fmt.Println("------------ Name login and password is not equals -------------")

	var nameUser, password string

	for {
		fmt.Println("Input a name user: ")
		fmt.Scan(&nameUser)

		fmt.Println("Input a password: ")
		fmt.Scan(&password)

		if strings.EqualFold(nameUser, password) {
			fmt.Println("Name and password is equais, try again!")
		} else {
			break
		}

	}

	fmt.Printf("The values of name login is: Login:%v\n", nameUser)
	fmt.Printf("The values of password is: password:%v\n", password)

}
