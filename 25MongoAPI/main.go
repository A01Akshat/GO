package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/A01Akshat/mongoapi/router"
)

func main() {
	fmt.Println("MONGOAPI")
	r := router.Router()
	fmt.Println("SERVER GETTING STARTED...")
	log.Fatal(http.ListenAndServe(":7000", r))
	fmt.Println("LISTENING AT PORT 7000...")
}
