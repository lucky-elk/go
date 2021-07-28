package main

import (
	"fmt"
	"strconv"
	"time"
)

const n = 1
const (
	price = n + 2
	_
	price2
	price3
	price4
)

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
	var v1 = "abc"
	resul, _ := strconv.Atoi(v1)
	fmt.Println(resul)

	//var v2 = 21000
	//var v3 = v1+strconv.Itoa(v2)
	//fmt.Println(v1)
	//fmt.Println(v3)
	//a,b := "张三","爸爸"
	//
	//result := fmt.Sprintf("我叫%s，是你%s",a,b)
	//fmt.Println(result)
	//var name string
	//fmt.Scanln(&name)
	//if name == "张三" {
	//	goto VVIP
	//}else if name == "李四"{
	//	goto VIP
	//}else{
	//
	//}
	//fmt.Println("A")
	//VIP:
	//fmt.Println("B")
	//VVIP:
	//fmt.Println("C")

	//data := 66
	//is := true
	//for is{
	//	var number int
	//	fmt.Println("请输入")
	//	fmt.Scanln(&number)
	//	if number > data {
	//		fmt.Println("大了")
	//	}else if number < data {
	//		fmt.Println("小了")
	//	}else{
	//		fmt.Println("成功")
	//		is = false
	//	}
	//}
	//fmt.Println(time.Second)
	//
	//for i:=1;i<=10;i++{
	//	fmt.Println("abc")
	//	time.Sleep(time.Second * 1)
	//	fmt.Println(i)
	//}
	//var n int
	//n=2
	//switch n{
	//case 1:
	//	fmt.Println(1)
	//case 2:
	//	fmt.Println(2)
	//case 3:
	//	fmt.Println(3)
	//default:
	//	fmt.Println("n")
	//}
	//var name string
	//var site string
	//var behavior string
	//fmt.Println("请输入姓名、位置、行为")
	//a,b := fmt.Scanf("%s 在%s 做%s",&name,&site,&behavior)
	//fmt.Println(a)
	//fmt.Println(b)
	//if name == "马化腾" {
	//	fmt.Println("yes")
	//}else {
	//	fmt.Println("no")
	//}
	//reader:= bufio.NewReader(os.Stdin)
	//line,_,_:= reader.ReadLine()
	//data := string(line)
	//fmt.Println(data)
	//var name string
	//var age int
	//fmt.Print("请输入用户名：")
	//_,err := fmt.Scanf("我叫%s 啊",&name)
	////_,_ = fmt.Scanf("我叫%s 啊",&name)
	//if err == nil {
	//	fmt.Println("用户名：",name,"\n年龄：",age)
	//}else{
	//	fmt.Println("用户信息输入错误",err)
	//}
	//fmt.Println(price)
	//fmt.Println(price2)
	//fmt.Println(price3)
	//fmt.Println(price4)
	//var sd string = "老男孩"
	//fmt.Println(sd)
	//fmt.Println(&sd)
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
	//api.Google()
	//api.Wiki()
	//fruits, num, price := "苹果", 2, 10.5
	//fmt.Printf("开车去卖%s,买%d个,一共%.2f元", fruits, num, price)
	//fmt.Println(&fruits)
}
