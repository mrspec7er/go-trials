package main

import (
	"fmt"
)

type CustomerType struct {
	category string
	amount int
	discount float32
}

func main () {
	var categoryInput string;
	var amountInput int;
	var discountInput float32;

	customers := []CustomerType{}

	for len(customers) < 3 {

		fmt.Println("Input Category")
		fmt.Scan(&categoryInput)
	
		fmt.Println("Input Amount")
		fmt.Scan(&amountInput)
	
		fmt.Println("Input discount")
		fmt.Scan(&discountInput)
	
		var customerData = CustomerType {
			category: categoryInput,
			amount: amountInput,
			discount: discountInput,
		}
	
		customers = append(customers, customerData)
	}

	fmt.Println("Data: ", customers)
}