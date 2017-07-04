package main 

import( 
	"fmt"
	"os"
	"strconv"
	"errors" // errors package
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	n,err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("not a valid number")
		}	else {
			fmt.Println(n)
		}
	}

func process(count int) error {
	if count < 1 {
		return errors.New("Invalid Count")
	}
	return nil
}

// Common pattern of using error variabes.
// eg. in io package EOF variable: 
// var EOF = errors.New("EOF")

// Go does have panic and recover functions. 
// panic is like throwing an exception 
// recover is like catch; 
// they are rarely used.

/*

Can create own error type,
The only constraint is that it fulfills the contract of the built-in error interface:

type error interface {
	Error() string
}

*/
