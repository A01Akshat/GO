package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// EncodeJson()
	DecodeJson()
}

type course struct {
	Name     string `json:"Coursename"` // aliasing
	Price    int
	Platform string `json:"website"`
	Password string
	Tags     []string
}

// Marshal = Go → JSON

// Unmarshal = JSON → Go

func EncodeJson() {
	lcourse := []course{
		{"Java", 199, "YT", "6490", []string{"lang", "spring"}},
		{"Go", 299, "Go.com", "8899", []string{"lang", "cloud"}},
		{"Data", 499, "YT", "3211", nil},
	}
	// PACKAGE THE DATA AS JSON

	finaljson, _ := json.MarshalIndent(lcourse, "", "\t") // For beautifying result
	fmt.Printf("%s", finaljson)
}

func DecodeJson() {
	jsondatafromweb := []byte(` 
	 {
                "Coursename": "Go",
                "Price": 299,
                "website": "Go.com",
                "Password": "8899",
                "Tags": [
                        "lang",
                        "cloud"
                ]
        }`)
	var lcourse course

	checkvalid := json.Valid(jsondatafromweb)
	if checkvalid {
		fmt.Println("VALID")
		json.Unmarshal(jsondatafromweb, &lcourse)
		fmt.Printf("%#v\n", lcourse)
	} else {
		fmt.Println("INVALID")
	}

	var mydata map[string]interface{} // value is not gurantee, it can be string, int or anything hence used interface
	json.Unmarshal(jsondatafromweb, &mydata)
	fmt.Println(mydata)

}
