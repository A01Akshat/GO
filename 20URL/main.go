package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://google.com:3000/learn?course=golang&paymentid=njdbchbd"

func main() {
	fmt.Println(myurl)

	res, _ := url.Parse(myurl) // Parse returns a pointer to a URL struct, which contains all the parts of the URL
	// fmt.Println(res.Scheme)
	// fmt.Println(res.Host)
	// fmt.Println("PATH:", res.Path)
	// fmt.Println(res.Port())

	queryparam := res.Query() // Query returns a map of query parameters, where the key is the parameter name and the value is a slice of strings
	fmt.Printf("param:%T\n", queryparam) 
	fmt.Println(queryparam["course"])   
	fmt.Println(queryparam["paymentid"]) 

	// CREATING NEW URL
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "yt.com",
		Path:    "/videos",
		RawPath: "user=Aksh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
