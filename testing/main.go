package main

import "fmt"

func main () {
	fmt.Println("Hello There!");

	const constantVariable = "This is how to define constant variable in GO!"

	var mutateableVariable = "And this is the changeble variable"

	mutateableVariable = "This variable can be change and redifined"

	const totalTicket = 50

	var ticketLeft = totalTicket - 7

	fmt.Println(constantVariable, "And", mutateableVariable);

	// variable inside string only work in printf
	fmt.Printf("Only %v tickets left\n", ticketLeft)
}