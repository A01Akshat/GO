package main

import "fmt"

func main() {

	num := 25

	var ptr = &num

	fmt.Println("Ptr's mem adress:",ptr)
	fmt.Println("Value held by ptr:",*ptr)

}