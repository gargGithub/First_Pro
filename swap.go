package main

import "fmt"

func swap(a *int,b *int){
	var temp int
	temp=*a
	*a=*b
	*b=temp
	fmt.Printf("a: %d b: %d",*a,*b)
}
func main() {
	x:=1
	y:=2
	swap(&x,&y)
}
