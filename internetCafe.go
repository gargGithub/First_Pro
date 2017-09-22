package main

import (
	"fmt"
	"math/rand"
	"time"

)
var x [25]int
var y [25]int
var i int = 7
func online(n int,t int){

	fmt.Println("Tourist ",n," is online")
	amt:=time.Duration(t)
	time.Sleep(time.Millisecond*amt)
	leftstatus(n,t)
    i++
	if(i<25) {
		online(x[i], y[i])
	}


}
func leftstatus(n int,t int){
	fmt.Println("tourist ",n," is done, having spent ",t," minutes online")
}

func waiting(n int,t int){
	fmt.Println("Tourist ",n," is waiting for turn")

}
func waitingReturn(n int,t int)(int,int){

    return n,t

}

func main() {
	var min [25]int
	var arr [25]int

	for i := 0; i < 25; i++ {
		arr[i] = i + 1
		min[i] = rand.Intn(106) + 15
	}

	for i := 1; i <= 25; i++ {
		if (i <= 8) {
			go online(arr[i-1], min[i-1])
		} else {

			waiting(arr[i-1], min[i-1])
			x[i-1], y[i-1] = waitingReturn(arr[i-1], min[i-1])
		}

	}
      time.Sleep(time.Millisecond*450)
	fmt.Println("The place is empty, let's close up and go to the beach!")
	var input string
	fmt.Scanln(&input)

}



