package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	//GET()
	postform()
}

// func GET() {
// 	const url = "http://localhost:8000/get"

// 	resp, _ := http.Get(url)

// 	defer resp.Body.Close()

// 	fmt.Println("Status:", resp.StatusCode)
// 	fmt.Println("Length:", resp.ContentLength)

// 	content, _ := io.ReadAll(resp.Body)
// 	fmt.Println(string(content))
// }

// func Post() {
// 	const url = "http://localhost:8000/post"
// 	requestBody := strings.NewReader(`
// 	{
// 	    "coursename":"LEARN GO",
// 		"price":100,
// 		"platform":"YT"
// 	}
// 	`)
// 	resp, _ := http.Post(url, "application/json", requestBody)
// 	defer resp.Body.Close()
// 	content, _ := io.ReadAll(resp.Body)
// 	fmt.Println(string(content))
// }

func postform() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("First", "Aksh")
	data.Add("Last", "Sri")
	data.Add("Email", "aksh@test.com")

	resp, _ := http.PostForm(myurl, data)
	content, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(content))

}
