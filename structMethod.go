package main

import (
	"fmt"

)
type Circle struct{
   r int

}



func main() {
	c:= Circle{3}
		fmt.Println(c.area())

}

func(c *Circle) area() int{
	area:= c.r*c.r
    return area
}