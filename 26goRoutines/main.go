package main

import (
	"fmt"
	"net/http"
	"sync"
)

// SOME ISSUES WITH CODE i guess
// BASICALLY USED WHEN WE HAVE HAVE MULTIPLE SERVERS AND WE WANT TO DO THINGS FOR EACH SERVER FAST
// BASICALLY TO ENHANCE PARALLELISM/CONCURENCY

var wg sync.WaitGroup //pointer, wg=waitgroup

func main() {
	// go greeter("Hello") // MADE IT A THREAD
	// greeter("World")

	websitelist := []string{
		"https://google.com",
		"https://youtube.com",
		"https://gmail.com",
	}
	for _, web := range websitelist {
		getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS")
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
