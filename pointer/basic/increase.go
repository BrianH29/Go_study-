package basic

import "fmt"

//Number format
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
