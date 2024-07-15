package main

//Importing methods of library fmt
import "fmt"

func main() {

	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Print("Printing type of data store in variables\n")

	//Using %T for discover of type data store in variable
	fmt.Printf("ConferenceTickets is %T, remainingTickets %T and conferenceName %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println("============= Printing values with on format interpolation =============")

	fmt.Printf("Welcome to %v booking application\n", conferenceName) //Concept of interpolation and explict new line "\n"
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Println("================= manipulation variables =================")

	//declaring variables
	var userName string
	var userTickets int

	userName = "Francisco"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
