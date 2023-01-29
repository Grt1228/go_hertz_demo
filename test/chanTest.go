package main

import (
	"fmt"
	"time"
)

func main() {

	//mainTestGo()
	//mainTestChan1()
	//mainTestChan2()
	mainTestChan3()
}
func f1(in chan int) {
	for true {

		fmt.Println(<-in)
	}
}

func mainTestChan3() {
	out := make(chan int)
	go f2(out)
	go f1(out)
	time.Sleep(1 * 1E9)
}

func f2(out chan int) {
	out <- 2
	out <- 222
	out <- 1
}

func mainTestChan2() {
	ch := make(chan int)
	go pump(ch)
	go suck(ch)
	time.Sleep(1 * 1e9)
	fmt.Println("main end")
}
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
func mainTestChan1() {
	ch := make(chan string)
	go send(ch)
	go get(ch)
	time.Sleep(1 * 1e9)
	fmt.Println("main end")
}

func send(ch chan string) {
	ch <- "A"
	ch <- "B"
	ch <- "C"
}

func get(ch chan string) {
	time.Sleep(2 * 1e9)
	for {
		fmt.Println(<-ch)
	}
}

func mainTestGo() {

	fmt.Println("start main~")
	go getProduct()
	go getStore()
	time.Sleep(2 * 1e9)
	fmt.Println("end main~")
}

func getProduct() {
	time.Sleep(5 * 1e9)
	fmt.Println("return product")
}

func getStore() {
	time.Sleep(3 * 1e9)
	fmt.Println("return store")
}
