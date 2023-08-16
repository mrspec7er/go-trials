package utility

import (
	"strings"
)

func GetPassangerName(passangers []string) []string  {
	var passangerName []string

	for _, passanger := range passangers { // _to define unused variable in golang, in this for loops it define the index of the slice / array
		var name = strings.Fields(passanger)
		passangerName = append(passangerName, name[1])
	}

	return passangerName
}