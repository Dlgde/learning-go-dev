package main

import "fmt"

func main() {
	ch01 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch01 <- i
	}

	close(ch01)

	fmt.Println("data length:", len(ch01))
	for v := range ch01 {
		fmt.Println(v)
	}

	fmt.Printf("data length: %d,", len(ch01))
}

/*
1. 使用range循环管道，如果管道未关闭会引发deadlocak错误（在此可通过注释打开close（ch01）来验证）
2. 如果采用for循环已经关闭的管道，当管道没有数据的时候，读取的数据会是管道的默认值，并且循环不会退出
*/
