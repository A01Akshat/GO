package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	mydefer()
}

func mydefer()  {
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}



// Defer moves the order of exection of that particular line to just before where the curly braces end