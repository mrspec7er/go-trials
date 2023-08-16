package main

import (
	"fmt"
	"strings"
)


func main()  {

	var passangers, total = addPassanger()
	var passangerName = getPassangerName(passangers)

	fmt.Println("Passanger list: ", passangers)
	fmt.Println("Total passanger: ", total)
	fmt.Println("Passanger name: ", passangerName)

	getSatisfyingLevel("Speaker")

}

func addPassanger() ([]string, int)  {
	var passangers []string // slice type is like an array but have dynamic value
	var passangerCount = 0;
	var newPassanger string;
	var stillWaiting string = "y";

	for stillWaiting == "y" { // in go there is no while loop, for loop can take it all

		fmt.Println("Input new passanger");
		fmt.Scan(&newPassanger);

		if newPassanger == "DRIVER" {
			continue
		}

		// check string contain and string length
		if strings.Contains(newPassanger, "#") || len(newPassanger) < 2 {
			fmt.Println("invalid username");
			continue
		}
	
		passangers = append(passangers, "PSG " + newPassanger);
		passangerCount++;
	
		fmt.Printf("New passanger with username %v entered the room, %v passanger ready to go!\n", newPassanger, passangerCount)
	
		if len(passangers) == 3 { // len() for check length of the slice / array
			break;
		}

		fmt.Println("Waiting for new passanger? (y/n)");
		fmt.Scan(&stillWaiting)

	}

	return passangers, passangerCount // go can return multiple value in one function
}

func getPassangerName(passangers []string) []string  {
	var passangerName []string

	for _, passanger := range passangers { // _to define unused variable in golang, in this for loops it define the index of the slice / array
		var name = strings.Fields(passanger)
		passangerName = append(passangerName, name[1])
	}

	return passangerName
}

func getSatisfyingLevel(peripherals string) {

	switch peripherals {
		case "Monitor":
			fmt.Println("Satisfying level 9");
		case "Keyboard":
			fmt.Println("Satisfying level 7");
		case "Speaker":
			fmt.Println("Satisfying level 10");
		default:
			fmt.Println("Satisfying level 8.5");
	}
}