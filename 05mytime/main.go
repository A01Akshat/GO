package main

import (
	"fmt"
	"time"
)

func main() {
	presenttime := time.Now()
	fmt.Println(presenttime)
	fmt.Println((presenttime.Format("01-02-2006 Monday")))
	// If you want date and day in that format always use that given date above and day as Monday
	
}