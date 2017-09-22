package main

import "fmt"

func printSum(args...int) int{
	sum:=0
	for _, v:= range args{
		sum+=v
	}
	return sum
}
func main() {
	arr:=[]int{1,2,3,4}
	fmt.Println(printSum(arr...))
}
