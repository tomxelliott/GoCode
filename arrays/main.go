package main

import "fmt"

func main() {
	// eg. the following creates a slice with a length of 0 but capacity of 10:
	scores4 := make([]int, 0, 10)
	scores4 = append(scores4, 5)
	fmt.Println(scores4) // prints 5 which we added to array

	scores5 := make([]int, 0, 10)
	scores5 = scores5[0:8]
	scores5[7] = 666
	fmt.Println(scores5)

	c := cap(scores5)
	fmt.Println(c)

	for i := 0; i < 25; i++ {
		scores5 = append(scores5, i)

		//grow array to accept the new data
		if cap(scores5) != c {
			c = cap(scores5)
			fmt.Println(c)

			// expands the array from 5 to 10 to 20 to 40.
		}
	}

	fmt.Println(scores5)
	// removes the last element from the array.
	scores5 = scores5[:len(scores5)-1]
	fmt.Println(scores5)

	scores := []int{1,2,3,4,5}
	scores = removeAtIndex(scores,2)
	fmt.Println(scores)
}

/*
swaps the last element in the array with the index stated
then removes that item from the array
*/
func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	//swap the last value and the value we want to remove
	//using pattern matching essentially
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}

/*
append is a special function because if the array is full,
it will create a new larger array and copy the values over.
much like a dynamic array in Python, PHP, Ruby et al.


// the following are ways to initialize slices:
names := []string{"oliver", "leno", "noah"}
checks := make([]bool, 10)
var names2 []string
scores6 := make([]int, 0, 20)


// hold indexes scores[0] up to scores[9]
var scores7 [10]int
scores7[0] = 330

// can also initialize array with:
ages := [4]int{27,25,30,11}

for index, value := range scores {

}

// Slices are a lightweight structure
// You don't declare a length
scores2 := []int{1,4,293,4,9}

// can also create using "make":
scores3 := make([]int, 10)
// this initializes the array length and capacity as 10
*/
