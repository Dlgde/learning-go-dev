package main

import (
	"fmt"
)

func test(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("send data:", i)
	}
}

func main() {
	ch01 := make(chan int, 20)
	ch02 := make(chan string, 20)

	go test(ch01)
	ch02 <- "hello"

	select {
	case a := <-ch01:
		fmt.Println("get data:", a)
	case b := <-ch02:
		fmt.Println("get data:", b)
	default:
		fmt.Println("no channel active")

	}

}
