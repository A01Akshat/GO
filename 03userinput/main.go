package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter ur age")

	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n') // to take input till enter is pressed 
	fmt.Println("Your age :", inp)
	

}


// inp, _ Instead of _ we could write inp,err
//That means that if any error would hae been there then that err will catch it just like try catch block