package main

// NOT CORRECT
import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.co.in/"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("TYPE: %T", resp)
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	fmt.Println(string(data))
}
