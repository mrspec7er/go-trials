package utility

import (
	"fmt"
)

var GlobalVariable = "This variable can be access in every file in this module!"

func GetSatisfyingLevel(peripherals string) {

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