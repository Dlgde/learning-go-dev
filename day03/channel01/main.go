package main

import (
	"fmt"
	"math/rand"
	"time"
)
//Produce is 
func Produce(header string, ch chan string) {
	for {
		ch <- fmt.Sprintf("%s:%v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}
//Consumer is 
func Consumer(ch chan string) {
	for {
		message := <-ch
		fmt.Println(message)
	}
}

func main() {
	ch := make(chan string)
	go Produce("pro01", ch)
	go Produce("pro02", ch)
	Consumer(ch)

}
