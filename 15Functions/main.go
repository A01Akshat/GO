package main

import "fmt"

func main() {
	x:= add(3, 5)
	fmt.Println(x);

	y,q:=proAdd(2,6,9,1)
	fmt.Println(y,q)

}

func add(a int, b int) int {
	return a + b
}

// We use this when we dont know how many values are passed 
// "values" will be passed ...int of int type

func proAdd(values ...int) (int, string){
	total:=0;
	for _,val:=range values{
		total+=val
	}
	return total,"VALUES ADDED"
}