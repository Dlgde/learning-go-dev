package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func test(num int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("number %v goroutine execute:%v\n", num, i)
	}

}

func main() {

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	fmt.Println("the main goroutine is out")
}

/*
printf 函数常见的转义字符
%d 十进制整数
%x %o %b 十六进制 八进制 二进制
%f %g %e 浮点数
%t  布尔型：true 或者false
%c  字符
%s  字符串
%q  带引号的字符串如：“abc”或者字符“c”
%v  内置格式的任何值
%T  任何值的类型
%%  百分号本身
*/
