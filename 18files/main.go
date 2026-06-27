package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is a demo content for the File"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("LENGTH:", length)
	defer file.Close() // Because we want it to execute at end
	readFile("./myfile.txt")
}

func readFile(filepath string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
