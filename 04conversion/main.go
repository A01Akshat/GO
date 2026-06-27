package main

import (
	"bufio"

	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	// Now we want to add 1 to input,hence type conversion

	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input) // Atoi means ASCII to integer

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New number:", number+1)
	}

}
