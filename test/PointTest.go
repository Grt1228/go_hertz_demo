package main

import (
	"fmt"
)

func main() {
	// a := 10
	// c := do1(&a)

	// fmt.Println(*c)
	// do2(a)
	a, b := do3(30, 10)
	fmt.Print(a, b)
}

func do1(a *int) *int {
	fmt.Println(a)
	b := a
	fmt.Println(b)

	return b

}

func do2(a int) {
	fmt.Print(a)
}

func do3(a int, b int) (min int, max int) {

	if a > b {
		min = b
		max = a
	} else {
		min = a
		max = b
	}
	return min, max
}
