package main

import (
	"../api"
	"fmt"
	"time"
)

const price int = 123

func variable() {
	var a int = 1
	var c string = "ac"
	c = "v"
	fmt.Printf("%d %d %q\n", a, a, c)
}

func sleepTime() {
	str := "倒计时"
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s %d\n", str, i)
		time.Sleep(1000000000)
	}
	fmt.Println("Boom!")
	time.Sleep(1000000000)
}

func main() {
	fmt.Println(price)
	var sd string = "老男孩"
	fmt.Println(sd)
	fmt.Println(&sd)
	//variable()
	//fmt.Println("hello world")
	//fmt.Println(1*1)
	//fmt.Println(1+1)
	//fmt.Println(1-1)
	//fmt.Println(1/1)
	//fmt.Println(1%1)
	//fmt.Println("abc"+"def")
	//fmt.Println(1>2)
	//fmt.Println(1<2)
	//if 1<2 {
	//	fmt.Println("真")
	//}else {
	//	fmt.Println("假")
	//}
	api.Google()
	api.Wiki()
	fruits, num, price := "苹果", 2, 10.5
	fmt.Printf("开车去卖%s,买%d个,一共%.2f元", fruits, num, price)
	fmt.Println(&fruits)
}
