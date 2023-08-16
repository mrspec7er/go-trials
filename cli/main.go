package main

import "fmt"

func main()  {

	var passangers, total = addPassanger()
	var passangerName = getPassangerName(passangers)

	fmt.Println("Passanger list: ", passangers)
	fmt.Println("Total passanger: ", total)
	fmt.Println("Passanger name: ", passangerName)

	getSatisfyingLevel("Speaker")

}