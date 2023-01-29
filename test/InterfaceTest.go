package main

import "fmt"

func main() {
	user := new(User)
	say := user.say("hello")
	eat := user.eat("book")
	fmt.Printf("say:%s \n", say)
	fmt.Printf("wat:%s", eat)
}

type UserInterface interface {
	say(text string) string
	eat(text string) string
}

type User struct {
}

func (user User) say(text string) string {
	return text
}

func (user User) eat(text string) string {
	return text
}
