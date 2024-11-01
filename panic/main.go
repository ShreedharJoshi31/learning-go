package main

func employee(name *string, age int) {
	if age > 65 {
		panic("Age cannot be greater than retirement")
	}
}

func main() {
	empName := "Flash"
	empAge := 66
	employee(&empName, empAge)
}