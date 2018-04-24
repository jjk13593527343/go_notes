package main

import (
	"time"
	"fmt"
)

func time1(){
	start := time.Now()
	fmt.Println(start)
	s := <- time.After( 2 * time.Second)
	fmt.Println(s)

}
func time2(){
	start := time.Now()
	timer := time.AfterFunc(2 * time.Second,func() {
		fmt.Println("asdfsadfasfd",time.Now().Sub(start))
	})

	time.Sleep(3 * time.Second)

	fmt.Println(timer)

}
func main() {
	time2()





}






