package main

import (
	"fmt"
	"strconv"
)

func main() {
	code := "M"
	code += fmt.Sprintf("%07d", 111111)
	fmt.Println(code)

	code = "M001"

	fmt.Println(strconv.Atoi(code[1:]))
}
