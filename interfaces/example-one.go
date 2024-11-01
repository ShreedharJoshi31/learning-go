package main

import "fmt"

type Animal interface {
	GetInfo() string
}

type Dog struct {
	Name  string
	Color string
}

func (d Dog) GetInfo() string {
	return fmt.Sprintf("The dog's name is %v and it's color is %v.", d.Name, d.Color)
}

type Cat struct {
	Name string
	Breed string
}

func (c Cat) GetInfo() string {
	return fmt.Sprintf("The cat's name is %v and it's breed is %v.", c.Name, c.Breed)
}

func main() {
	dog := Dog{
		Name: "Sparky",
		Color: "White",
	}
	cat := Cat{
		Name: "Sparkles",
		Breed: "Dont know",
	}

	fmt.Println(dog.GetInfo())
	fmt.Println(cat.GetInfo())
}