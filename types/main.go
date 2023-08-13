package main

import "fmt"

func main()  {
	const first string = "First sentence, ";

	const second string = "Second sentence";

	const count int = 5;

	const status bool = true

	fmt.Println(first + second)

	fmt.Println(count)

	fmt.Println(status)

	list:= [3]string{"name", "email", "password"};

	fmt.Printf("List has type of %T, and the value is %v\n", list, list)
}