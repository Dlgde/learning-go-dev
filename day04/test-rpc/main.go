package main

import (
	"errors"
	"fmt"
	"time"
)

//RPCclient is client
func RPCclient(ch chan string, req string) (string, error) {
	ch <- req
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("time out") //超时
	}
}

//RPCserver is server
func RPCserver(ch chan string) {
	for {
		data := <-ch
		fmt.Println("server recieved:", data)
		//加入延迟两秒，测试超时
		time.Sleep(time.Second * 2)
		ch <- "roger"
	}

}

func main() {
	ch := make(chan string)
	go RPCserver(ch) //并发执行服务器逻辑

	recv, err := RPCclient(ch, "hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client receieved", recv)
	}
}
