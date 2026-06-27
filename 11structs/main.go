package main

import "fmt"

// NO INHERITANCE, NO SUPER,PARENT IN GO
func main() {
    Akshat:=User{"Aksh","ak@test",true,22}
	fmt.Println((Akshat))
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int	
}

// LOOPS

// days:=[]string{"Mon","Tues","Wed"}

// for i:= range(days){
// 	fmt.Println(days[i])
// }
