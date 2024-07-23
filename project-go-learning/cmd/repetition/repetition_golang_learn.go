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
	// 3 - calling function = validationInformationPersona()
	validationInformationPersona()

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

func validationInformationPersona() {
	fmt.Println("------ Validation information personal ------")

	// section of variables
	var name string
	var age int
	var salary float32
	var gender string
	var stateCivil string

	//Input data and validation information
	fmt.Println("Enter with your name: ")
	fmt.Scan(&name)
	for len(name) <= 3 {
		fmt.Println("Size name need most three character")
		fmt.Println("Enter with a name, please: ")
		fmt.Scan(&name)
	}
	fmt.Println("Next step = Input your age...")
	fmt.Println("Enter with your age: ")
	fmt.Scan(&age)
	for age < 0 || age > 150 {
		fmt.Println("Age invalid, try again!")
		fmt.Println("Your age, please: ")
		fmt.Scan(&age)
	}
	fmt.Println("Next step = input your salary with value > 0")
	fmt.Println("Enter with you salary, please: ")
	fmt.Scan(&salary)
	for salary < 0 {
		fmt.Println("Your value of salary is invalid, try again!")
		fmt.Scan(&salary)
	}
	fmt.Println("Next step = input your gender")
	fmt.Println("Enter with your gender: F/f - Female or M/m - Male")
	fmt.Scan(&gender)
	for gender != "f" && gender != "F" && gender != "m" && gender != "M" {
		fmt.Println("Gender is invalid, try again!")
		fmt.Println("What your gender: ")
		fmt.Scan(&gender)
	}
	fmt.Println("Next step = input your state civil")
	fmt.Println("Enter with your state civil: ")
	fmt.Scan(&stateCivil)
	for stateCivil != "s" && stateCivil != "c" && stateCivil != "v" && stateCivil != "d" {
		fmt.Println("State civil is invalid, try again!")
		fmt.Println("Enter with your state civil: ")
		fmt.Scan(&stateCivil)
	}
}
