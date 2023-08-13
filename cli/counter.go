package main

import "fmt"

func main()  {
	var passangerCount = 0;
	var passanger []string // slice type is like an array but have dynamic value
	var newPassanger = "";
	var stillWaiting = true;

	for stillWaiting {

		fmt.Println("Input new passanger");
		fmt.Scan(&newPassanger);
	
		passanger = append(passanger, newPassanger);
		passangerCount++;
	
		fmt.Printf("New passanger with username %v entered the room, %v passanger ready to go!\n", newPassanger, passangerCount)
	
		fmt.Println("Waiting for new passanger?");
		fmt.Scan(&stillWaiting)
	}

	fmt.Println("Passanger list: ", passanger)

}