package main

import "fmt"

type Human interface {
	Name() string
	Profession() string
}

type User1 struct {
	Age    int
	Gender string
}

type User2 struct {
	Age      int
	Gender   string
	Location string
}

func (u1 User1) Name() string {
	return "Harry"
}

func (u1 User1) Profession() string {
	return "Frontend Dev"
}

func (u2 User2) Name() string {
	return "Ron"
}

func (u2 User2) Profession() string {
	return "Backend Dev"
}

func displayInformation(h Human) {
	fmt.Printf("My name is %v and my profession is %v.\n", h.Name(), h.Profession())
}

func main() {
	harry := User1{
		Age: 21,
		Gender: "Male",
	}
	ron := User2{
		Age: 23,
		Gender: "Female",
		Location: "Somewhere",
	}
	displayInformation(harry)
	displayInformation(ron)
}