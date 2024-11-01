package main

import "fmt"

type Describer interface {
	Describe() string
}

type Person struct {
	Name string
}

func (p Person) Describe() string {
	return "Hi, I am a person with name " + p.Name
}

type Spider struct {
	Name string
}

func (s Spider) Describe() string {
	return "Hi, I am a Spider with name " + s.Name
}

func displayDescription(d Describer) {
	fmt.Println(d.Describe())
}

func main() {
	person := Person{
		Name: "FLASH",
	}
	spider := Spider {
		Name: "SOMEONE",
	}
	displayDescription(person)
	displayDescription(spider)
}