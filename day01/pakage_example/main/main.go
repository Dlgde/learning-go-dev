package main

import (
	"fmt"
	"go_dev/day02/package_demo/calc"
)

func main() {
	sum := calc.Add(100, 300)
	sub := calc.Sub(100, 250)

	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)

}
