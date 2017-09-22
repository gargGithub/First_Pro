package main

import "fmt"

func f(i int){
	for n:=1;n<10;n++{
		fmt.Println(i,":",n)
	}
}


func main() {
	for i:=0;i<10;i++ {
		go f(i)
	}


	var input string
	fmt.Scanf("%s",&input)
}
