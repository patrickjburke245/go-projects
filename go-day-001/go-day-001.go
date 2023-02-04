//It's common to develop a piece of code to be reused across many applications/other pieces of code.
//This shared piece of code is called a library or package.
//The main package tells the Golang compiler that this piece of code should become an executable program upon compilation, not a shared library.
package main

import "fmt"

func main() {
	fmt.Println("Welcome to Club PB & Jammin'!!!")

	//Acquire attendee's first name
	fmt.Print("Enter your first name: ")
	var firstName string
	fmt.Scanln(&firstName)

	//Acquire attendee's last name
	fmt.Print("Enter your last name: ")
	var lastName string
	fmt.Scanln(&lastName)
	
	//Direct attendee into club
	fmt.Println("Right this way, " + firstName + " " + lastName + ". Your friends will be delighted to see you.")
}