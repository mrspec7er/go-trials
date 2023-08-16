package main

import (
	"fmt"
	"strings"
)

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