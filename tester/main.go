package main 

import (
	"fmt"
)

func main() {
	power := 1000
	fmt.Println("default power rating is", power)

	name, power := "Tom", 5000
	fmt.Println(name,"has the power of", power)
	fmt.Printf("%s's power is over %d\n", name, power)

	name = "Bob"
	fmt.Println(name)
}

func power(name string) (int, bool) {
	value, exists := power("Goku")
	if exists == false {
		fmt.Println("Hello")
	} else {
		fmt.Println("Goodbye")
	}
	return value, exists
}

func add(a,b int) int {
	return a+b
}

func log(message string) {
	fmt.Println(message)
}

func random(name string) (int, bool) {
	var g int
	g = 1000
	_, exists:= random("Tom")
	// underscore character is a blank identifier - 
	// it is special in that the return value isn't actually assigned.
	if exists == false {
		fmt.Println("OK")
	}
	return g, exists
}
