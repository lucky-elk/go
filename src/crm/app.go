package main

import (
	"fmt"
	"time"
)

func variable() {
	var a int = 1
	var c string = "ac"
	c="v"
	fmt.Printf("%d %d %q\n",a,a,c)
}

func main() {
	//variable()
	fmt.Println("hello world")
	//Vik()
	//api.Baidu()
	//api.BaiDu()
	str := "倒计时"
	for i:=1;i<=3;i++ {
		fmt.Printf("%s %d\n",str,i)
		time.Sleep(1000000000)
	}
	fmt.Println("Boom!")
	time.Sleep(1000000000)
}