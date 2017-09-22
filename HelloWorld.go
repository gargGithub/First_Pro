package main

import (
	"fmt"
	"unsafe"
)
func sum(arg []int,arr []int) []int{
	sum:=make([]int,3)
    final:=make([]int,1)
	//for _,i := range args{
	//	for _,j:=range args {

	//		copy(sum, args)
	//	}
	//}
//
	copy(sum, arg)
//	for _,v:=range arr{

//		final=append(sum, v)
//	}
	//fmt.Println(sum)
	fmt.Println(len(final))
	fmt.Println(unsafe.Sizeof(final))
	final = append(sum,arr[0],arr[1],arr[2])
	fmt.Println(len(final))
	return final
}
func main() {
    arr:= []int{2,3,5}
    arr1:=[]int{5,4,3}
	fmt.Println(sum(arr,arr1))
}

