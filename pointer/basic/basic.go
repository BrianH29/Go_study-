package basic

import "fmt"

//Basic pointer example
func Basic() {
	var a int
	var b int
	var p *int // p is currently a's memory address *p has the value

	p = &a
	a = 3
	b = 2

	// & : memory address , * : value
	fmt.Println("a:", a)
	fmt.Println("p:", p)
	fmt.Println("*p:", *p)

	p = &b
	fmt.Println("b:", b)
	fmt.Println("p:", p)
	fmt.Println("*p:", *p)
}
