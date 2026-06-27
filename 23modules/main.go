package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET") 

	log.Fatal(http.ListenAndServe(":7001", r))
}

func greeter() {
	fmt.Println("HEY")
}

func serveHome(w http.ResponseWriter, r *http.Request)  { 
	w.Write([]byte("<h1>WELCOME</h1>")) 
}
