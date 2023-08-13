package main

import "fmt"

func main()  {
	var username string
	var password string

	fmt.Println("Input Username: ")
	fmt.Scan(&username)

	fmt.Println("Input password: ")
	fmt.Scan(&password)

	fmt.Printf("Successfully register with username: %v and password %v\n", username, password)
}