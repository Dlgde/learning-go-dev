package main

import "fmt"

func test01() {
	abc := make([]int, 2, 7)
	fmt.Printf("%p\n", abc)
	abc = append(abc, 1, 2, 3)
	fmt.Printf("%p\n", abc)
	fmt.Println(abc)
	fmt.Println(cap(abc))
	fmt.Println(len(abc))

}
func test02() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:6]
	s2 := s1[3:6]
	fmt.Println("s1==", s1)
	fmt.Println("s2==", s2)

}

func main() {
	test01()
	// test02()
}
