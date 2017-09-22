package main

import (
	"fmt"
	"time"
)


func ping(channel chan string){
	for{
		channel <- "ping"
	}
}
func pong(channel chan string){
	for i:=0; ;i++{
		channel <- "pong"
	}
}
func printer(channel chan string){
	for{
		msg := <-channel
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}

}
func main() {

	var channel chan string = make(chan string)
	go ping(channel)
	go pong(channel)
	go printer(channel)
	var input string
	fmt.Scanln(&input)
}
