package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test()", i)
		time.Sleep(time.Microsecond * 20)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main(),", i)
		time.Sleep((time.Millisecond * 20))
	}
	// time.Sleep(time.Second)

}
