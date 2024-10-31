package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age int
}

func (p Person) FullName() string {
	return p.firstName + " " + p.lastName
}

func (p Person) CanVote() bool {
	return p.age >= 18
}

func main() {
	person := Person{firstName: "Some", lastName: "thing", age: 21}
	fullName := person.FullName()

	fmt.Println(fullName)
	fmt.Println(person.CanVote())
}