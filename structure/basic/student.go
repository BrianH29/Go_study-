package basic

import "fmt"

//Student structure
type Student struct {
	name  string
	class int

	grade Grade // students grade
}

//Grade structure
type Grade struct {
	name  string
	grade string
}

func (s Student) viewGrade() {
	fmt.Println(s.grade)
}

func (s Student) inputGrade(name string, grade string) {
	s.grade.name = name
	s.grade.grade = grade
}

func viewGrade(s Student) {
	fmt.Println(s.grade)
}

func inputGrade(s Student, name string, grade string) {
	s.grade.name = name
	s.grade.grade = grade
}

//SendInfo to main
func SendInfo() {
	var s Student
	s.name = "Jack"
	s.class = 1

	s.grade.name = "Math"
	s.grade.grade = "C"

	s.viewGrade()
	viewGrade(s)

	s.inputGrade("History", "A")
	s.viewGrade()

	inputGrade(s, "History", "A")
	viewGrade(s)
}
