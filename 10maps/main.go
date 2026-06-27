package main

import "fmt"

func main() {
	languages := make(map[string]string) // Key Value

	languages["Py"] = "Python"
	languages["Ja"] = "Java"
	languages["Go"] = "Golang"
	fmt.Println(languages)
	delete(languages,"Py")
	fmt.Println(languages)

	for key, value:=range languages{
	    fmt.Printf("Key:%v, Value:%v\n",key,value)
	}
}