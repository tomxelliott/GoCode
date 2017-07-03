package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		// len returns the size of a string
		// or the number of values in a dictionary, 
		// or as in this function, the number of elements in an array
		os.Exit(1)
	}	
	fmt.Println("It's over", os.Args[1])
	// returns the number of arguments passed to the function
	
}
