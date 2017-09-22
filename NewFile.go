package main

import "fmt"

type hyundai struct{
	distance int
	fuel int
	name string
}
type marutiSuzuki struct{
	distance int
	fuel int
	name string
}
type ford struct{
	distance int
	fuel int
	name string
}
type honda struct{
	distance int
	fuel int
	name string
}
type Cars interface {
	carName() string
	mileage() int
}

func main() {
	hynd:=hyundai{20,1,"Xcent"}
	mZ:=marutiSuzuki{13,1,"Swift"}
	frd:=ford{22,2,"Figo"}
	hnda:=honda{15,1,"Amaze"}

	bestMileage(&hynd,&frd,&mZ,&hnda)
}


func(hynd *hyundai) mileage()int{
	mileage:=hynd.distance/hynd.fuel
	return mileage
}
func(mz *marutiSuzuki) mileage()int{
	mileage:=mz.distance/mz.fuel
	return mileage
}
func(frd *ford) mileage()int{
	mileage:=frd.distance/frd.fuel
	return mileage
}
func(hnda *honda) mileage()int{
	mileage:=hnda.distance/hnda.fuel
	return mileage
}
func(hynd *hyundai) carName()string{
	name:=hynd.name
	return name
}
func(hnda *honda) carName()string{
	name:=hnda.name
	return name
}
func(frd *ford) carName()string{
	name:=frd.name
	return name
}
func(mZ *marutiSuzuki) carName()string{
	name:=mZ.name
	return name
}
func bestMileage(car...Cars)  {
	bmileage:=18
	var name string
	for _,v:=range car {
         if bmileage<v.mileage() {
			 bmileage = v.mileage()
			 name = v.carName()
		 }
	}

	fmt.Printf("%s is economical with mileage of %d",name,bmileage)
}
