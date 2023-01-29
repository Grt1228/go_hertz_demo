package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.IndexFunc("BCDAEFG", isA))
	fmt.Println(strings.IndexFunc("B1DAEFG", unicode.IsNumber))

	fmt.Println(strings.Map(removeSpace,"i am tomi."))
	fmt.Println(callback(10, Add))
}

func removeSpace(x rune) rune {
	if x == ' '{
		return 0
	}
	return x
}

func isA(x int32) bool {
	if x == 65 {
		return true
	}
	return false
}

func Add(a int, b int) int {
	return a + b
}

func callback(x int, f func(a int, b int) int) int {
	i := f(x, 10000)
	return i
}
