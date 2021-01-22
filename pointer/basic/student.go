package basic

import "fmt"

//Student VO
type Student struct {
	name string
	age  int

	grade string
	class string
}

//PrintGrade result
func (s *Student) PrintGrade() {
	fmt.Println(s.class, s.grade)
}

//InputGrade input grade
func (s *Student) InputGrade(class string, grade string) {
	s.class = class
	s.grade = grade
}

//Result of outcome
func Result() {
	var s Student = Student{name: "Brian", age: 20, class: "영어", grade: "A"}

	s.InputGrade("국어", "B")
	s.PrintGrade()
}
