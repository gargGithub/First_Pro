package main

import (
	"fmt"
	"math"
)

type Circle struct{
	r float64

}
type Rect struct{
	l,b float64
}

type shape interface {
	area() float64
	per() float64
}
func totalArea(shapes...shape) float64{
	var area float64
	for _,v:=range shapes{
		area+=v.area()
	}
	return area
}
func totalper(shapes...shape) float64{
	var perimeter float64
	for _,v:=range shapes{
		perimeter+=v.per()
	}
	return perimeter
}

func main() {
	c:= Circle{3}
	r:=Rect{2,3}
	fmt.Println(c.area())
	fmt.Println(r.area())
	fmt.Println(c.per())
	fmt.Println(r.per())
	fmt.Println(totalArea(&c,&r))
	fmt.Println(totalper(&c,&r))


}

func(r *Rect) area() float64{
	area:=r.l*r.b
	return area
}
func(c *Circle) area() float64{
	area:=math.Pi*c.r*c.r
	return area
}
func(r *Rect) per() float64 {
	perimeter := 2*(r.l + r.b)
	return perimeter
}
func(c *Circle) per() float64 {
	perimeter := 2*math.Pi*c.r
	return perimeter

}