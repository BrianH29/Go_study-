package basic

import "fmt"

//Person structure
type Person struct {
	name string
	age  int
}

func (p Person) printName() {
	fmt.Println(p.name)
}

//PrintPerson list of person
func PrintPerson() {
	var p Person
	p1 := Person{"Jane", 15}
	p2 := Person{name: "Jack", age: 22}
	p3 := Person{name: "Sam"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)

	p.name = "Smith"
	p.age = 25

	fmt.Println(p)

	p.printName()
}
