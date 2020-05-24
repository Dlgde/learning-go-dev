package main

import (
	"fmt"
	"go_dev02/day01/goroute_example/goroute"
)

func main() {
	var pipe chan int
	pipe = make(chan int, 1)
	go goroute.Add(100, 200, pipe)

	sum := <-pipe

	fmt.Println("sum=", sum)
}
