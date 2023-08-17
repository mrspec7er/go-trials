package main

import (
	"fmt"
	"time"
)

type CustomerType struct {
	id int
	category string
	amount int
	discount float32
}

func main () {
	var categoryInput string;
	var amountInput int;
	var discountInput float32;
	var message string

	customers := []CustomerType{}

	for len(customers) < 2 {

		fmt.Println("Input Category")
		fmt.Scan(&categoryInput)
	
		fmt.Println("Input Amount")
		fmt.Scan(&amountInput)
	
		fmt.Println("Input discount")
		fmt.Scan(&discountInput)
	
		var customerData = CustomerType {
			id: len(customers) + 1,
			category: categoryInput,
			amount: amountInput,
			discount: discountInput,
		}
	
		customers = append(customers, customerData)
	}

	for _,customer := range customers {
		go generateReport(customer)
	}

	fmt.Scan(&message);
	fmt.Println(message);
}

func generateReport(customer CustomerType) {
	time.Sleep(5 * time.Second)
	fmt.Println("______________________")
	fmt.Printf("ID: %v, Category: %v, Total: %v\n",customer.id, customer.category, customer.amount * 2000)
}