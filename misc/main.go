package main 

import( 
	"fmt"
)

// functions are first-class types
// can be used anywhere: as a field, parameter or return value
type Add func(a int, b int) int

func main() {
	fmt.Println("String to Bytes arrays...")
	count := 2
	// Strings and byte arrays are closely related
	// The following is a conversion between the two.
	// When performing this action we are creating a copy as strings are immutable
	stra := "I will be compartmentalised into an array"
	byts := []byte(stra)
	strb := string(byts)	
	//int32 and int64 can be easily converted.
	//eg.
	fmt.Println(strb)
	fmt.Println(int64(count))

	fmt.Println(process(func(a int, b int) int {
		return a + b
		}))
}

// Strings are made of runes which are unicode code points.
// If you take the length of a string, you might not get what you expect
// Print results etc could be not as expected
// If you iterate over a string using range, you'll get runes not bytes.
// When you turn a string into a []byte you'll get the correct data


// using functions like this helps decouple code much like with interfaces
func process(adder Add) int {
	return adder(1,2)
}
