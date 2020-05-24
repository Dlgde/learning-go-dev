package main

import (
	"fmt"
	"runtime"
	"time"
)

func test(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
		runtime.Gosched()

	}
}

func main() {
	// runtime.GOMAXPROCS(2)
	go test("world")
	test("hello")
	time.Sleep(time.Second)
}

// runtime.Gosched()用于让出CPU时间片。这就像跑接力赛，A跑了一会碰到代码runtime.Gosched()就把接力棒交给B了，A歇着了，B继续跑
// 当一个goroutine发生阻塞，Go会自动地把与该goroutine处于同一系统线程的其他goroutines转移到另一个系统线程上去，以使这些goroutines不阻塞
