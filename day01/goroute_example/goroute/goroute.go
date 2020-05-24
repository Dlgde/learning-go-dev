package goroute

func Add(a, b int, c chan int) {
	c <- (a + b)
}
