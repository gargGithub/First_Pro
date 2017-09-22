package main

import (
	"time"
	"fmt"
	"math"
)

func main() {

	c1:=make(chan string)
	c2:=make(chan string)
	c3:=make(chan string)
	c4:=make(chan string)
	c5:=make(chan string)
	c6:=make(chan string)
	c7:=make(chan string)
	c8:=make(chan string)



	go func(){
		for{

	         c1<-"from1"
			time.Sleep(time.Second*2)
		}
	}()
	go func(){
		for{

			c2<-"from2"
			time.Sleep(time.Second*3)
		}
	}()
	go func(){
		for{

			select{
			case msg1:= <-c1: fmt.Println(msg1)
			case msg2:= <-c2: fmt.Println(msg2)
			case <- time.After(time.Second): fmt.Println("timeout")
			//default: fmt.Println("Nothing Ready")
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}
