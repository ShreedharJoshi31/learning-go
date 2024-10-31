package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	person Person
	salary int
	job string
}

func main() {
	user := Person{name: "Harry", age: 20}
	fmt.Println(user.name, user.age)
	
	var user1 Employee
	user1.person.name = "Flash"
	user1.person.age = 21
	user1.salary = 20000
	user1.job = "Speedster"
	fmt.Println(user1)

	person := struct {
		firstName string
		lastName string
	} {
		firstName: "Some",
		lastName: "thing",
	}
	fmt.Println(person.firstName, person.lastName)
}