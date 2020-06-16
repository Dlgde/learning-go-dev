package main

import (
	"fmt"
	"strconv"
	"sync"
)

var syncMap sync.Map
var wg sync.WaitGroup

func Addnumber(b int) {
	for i := 0; i < b; i++ {
		syncMap.Store(i, i)
	}
	wg.Done()
}

func main() {
	m := 5
	wg.Add(m)
	var size int

	for n := 0; n < m; n++ {
		go Addnumber(n)
	}

	syncMap.Range(func(key, value interface{}) bool {
		size++
		fmt.Println(key, value)
		return true
	})

	fmt.Println(strconv.Itoa(size))

	v, ok := syncMap.Load(0)
	if ok {
		fmt.Println("key 0 has value", v, "")
	}
	// fmt.Println("there is wrong")

	wg.Wait()

}
