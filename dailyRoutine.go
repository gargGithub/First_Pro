package main

import (
	"fmt"
	"math/rand"
	"time"
)

func johnReady(n int){
	amt:=time.Duration(n)
	time.Sleep(time.Second*amt)
	fmt.Println("John spent ",n," seconds getting ready!")

}
func jackReady(n int) {
	amt:=time.Duration(n)
	time.Sleep(time.Second*amt)
	fmt.Println("Jack spent ", n, " seconds getting ready!")

}
func johnShoes(n int) {
	amt:=time.Duration(n)
	fmt.Println("John started putting on shoes")
	time.Sleep(time.Second*amt)
	fmt.Println("John spent ", n, " seconds putting on shoes")

}
func jackShoes(n int) {
	amt:=time.Duration(n)
	fmt.Println("Jack started putting on shoes")
	time.Sleep(time.Second*amt)
	fmt.Println("Jack spent ", n, " seconds putting on shoes")

}
func armAlarm(){


	fmt.Println("Alarm is couting down")
	time.Sleep(time.Second*60)
	fmt.Println("Alarm is armed!")

}
func main() {

	jR:=rand.Intn(31)+60
	jaR:=rand.Intn(31)+60
	js:=rand.Intn(11)+35
	jas:=rand.Intn(11)+35


	fmt.Println("Let's go for a walk!")
	fmt.Println("John started getting ready")
	fmt.Println("Jack started getting ready")
	if(jR<jaR){
         go johnReady(jR)
		jackReady(jaR)
	}else {
	go jackReady(jaR)
	johnReady(jR)
	}
	fmt.Println("Arming Alarm!")
	go armAlarm()


	if(js<jas){
		go johnShoes(js)

		jackShoes(jas)
	}else {
		go jackShoes(jas)
		johnShoes(js)
	}
	fmt.Println("Exiting & Locking the door")
	var input string
	fmt.Scanln(&input)
}
