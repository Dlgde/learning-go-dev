package main

import (
	"fmt"
	"time"
)

func testch(ch chan int) {
	for i := 0; i < 3; i++ {
		a := <-ch

		fmt.Println(a)
	}
}

func main() {
	ch := make(chan int, 3)

	ch <- 10
	close(ch)

	go testch(ch)
	time.Sleep(time.Second * 1)

}

// 如果采用for循环已经关闭的管道，当管道没有数据的时候，读取的数据会是管道的默认值，并且循环不会退出
