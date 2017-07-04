package main

import "fmt"

func main() {
	// create Pet that uses extension of Person struct
	spike := &Pet{
	Person: &Person{"Spike", 12, nil},
	Speed: 21,
	}
	spike.Person.SayHi()
	// can also do:
	// spike.SayHi()

	tom := &Person {
		Name: "Bart",
		Age: 10,
		Father: &Person {
			Name: "Homer",
			Age: 35,
			Father: nil,
		},
	}
	tom.SayHi()
}


type Person struct {
	Name string
	Age int
	Father *Person
}

func (p *Person) SayHi() {
	fmt.Printf("Hi... I'm %s\n", p.Name)
}

type Pet struct {
	*Person
	Speed int
}


// Structures don't have Constructors. Just Methods that initialise the given type. 
func NewPerson(name string, age int) *Person {
	return &Person {
		Name: name,
		Age: age,
	}
}
