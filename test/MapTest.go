package main

import "fmt"

func main() {

	user := make(map[string]string)

	user["name"] = "小A"
	user["age"] = "18"

	fmt.Println(user["name"])
	fmt.Println(len(user))
	delete(user, "age")
	age, exist := user["age"]
	fmt.Println(age)
	fmt.Println(exist)
	fmt.Println(len(user))
	name, exist := user["name"]
	fmt.Println(name)
	fmt.Println(exist)

	for key, value := range user {
		fmt.Printf("key: %s,value：%s\n", key, value)
	}

	sum1, _ := plus2(1, 2)
	fmt.Println(sum1)

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>")
	array1 := []int{1, 1, 2, 3, 10}
	fmt.Println(nums(array1...))
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>")
	s2 := make([]int, 3)
	s2[2] = 100
	fmt.Println(nums(s2...))

}

func plus(a, b int) int {
	return a + b
}

func plus2(a, b int) (int, int) {
	return a + b, b + a
}

func nums(num ...int) int {
	sum := 0
	for _, n := range num {
		sum += n
	}
	return sum
}
