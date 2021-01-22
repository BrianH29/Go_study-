# Go_study-

#Pointer

```go
func Number() {
	a := 1

	for i := a; i < 10; i++ {
		increase(&a)
		fmt.Print(a, " ")
	}

}

func increase(x *int) {
	// x++ --> without point outcome 1 1 1 1 1 1 1 1 1
	*x = *x + 1
}
```

- **a** will have memory box with the value of 1
- x will have seperate memory box with **a**'s memory address
- \*x will have value of **a'**s memory address
- *x = *x +1 will change the value of the **a** which was 1 to 2

#another study example

```go
type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintGrade() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputGrade(class string, grade string) {
	s.class = class
	s.grade = grade
}

func Result() {
	var s Student = Student{name: "Brian", age: 20, class: "영어", grade: "A"}

	s.InputGrade("국어", "B")
	s.PrintGrade()
}
```

1. without pointer
   - without pointer all the **s**' will create different memory adress and it will only copy the value of the original **s**
   - So even if you create a func to change the value it will not change because original value memory address is different
   - In order to CHANGE the originial **s** you need to create the memory address that is same to the original which is the use of the pointer
   - after the run of Result func InputGrade func will create a same memory address as the Result var s Student
