package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now().Unix()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go sushuGoRun(i)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end - start)
}

func main_A() {
	start := time.Now().Unix()

	for i := 1; i <= 300000; i++ {
		sushu(i)
	}
	end := time.Now().Unix()
	fmt.Println(end - start)
}

func sushuGoRun(start int) {
	for i := start * 30000; i < 3000*(start+1); i++ {
		sushu(i)
	}
	wg.Done()
}

func sushu(num int) bool {
	if num < 2 {
		return false
	}
	var isPrime bool = true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return isPrime
}
