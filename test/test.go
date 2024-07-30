package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := createPerson("John", 20)
	fmt.Println(*person)
	changePersonName("Doe", person)
	fmt.Println(*person)
	changePersonAge(30, person)
	fmt.Println(*person)
}

func createPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func changePersonName(name string, p *Person) {
	p.name = name
}

func changePersonAge(age int, p *Person) {
	p.age = age
}
