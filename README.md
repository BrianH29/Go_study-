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
