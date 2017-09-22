package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(N int){
	for i:=1;i<10;i++{
		fmt.Println(N,":",i)
		amt:=time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond*amt)
	}
}
func main() {
	for i:=0;i<10;i++{
		go f(i)
	}
	var input string
	fmt.Scanf("%s",&input)


}
