package main

import (
	"fmt"
	"testing/cli/utility"
)

func main()  {

	var passangers, total = addPassanger()
	var passangerName = utility.GetPassangerName(passangers)

	fmt.Println("Passanger list: ", passangers)
	fmt.Println("Total passanger: ", total)
	fmt.Println("Passanger name: ", passangerName)

	utility.GetSatisfyingLevel("Speaker")
	fmt.Println(utility.GlobalVariable)

}