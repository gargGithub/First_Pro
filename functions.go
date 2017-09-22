package main

import "fmt"

func fib(N int) int{
	if(N==0){
		return 0
	} else if(N==1){
	return 1
	} else{
	return fib(N-1)+fib(N-2)
	}


}
func main() {
  y:=fib(6)
	fmt.Println(y)

}
