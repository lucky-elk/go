package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

//全局变量
//var aa string = "aa"

func main() {
	/****5-12字符串常见功能****/
	name := "张三三"
	name2 := "zhangsansan"
	name3 := "ZHANGSANSAN"
	//获取字节长度
	fmt.Println(len(name)) //输出：9
	//获取字符长度
	fmt.Println(utf8.RuneCountInString(name)) //输出：3

	//判断是否以“x”开头
	res1 := strings.HasPrefix(name, "张")
	fmt.Println(res1) // 输出：true
	//判断是否以“x”结尾
	res2 := strings.HasSuffix(name, "三")
	fmt.Println(res2) // 输出：true
	//判断是否包含“x”
	res3 := strings.Contains(name, "三")
	fmt.Println(res3) // 输出：true

	//转大写
	res4 := strings.ToUpper(name2)
	fmt.Println(res4) // 输出：ZHANGSANSAN
	//转小写
	res5 := strings.ToLower(name3)
	fmt.Println(res5) // 输出：zhangsansan

	//去两边
	result1 := strings.TrimRight(name2, "an") // 去除右边的an
	result2 := strings.TrimLeft(name2, "zh")  // 去除左边的zh
	result3 := strings.Trim(name2, "n")       // 去除两边的n
	fmt.Println(result1, result2, result3)    // 输出：zhangsans angsansan zhangsansa

	//替换
	match1 := strings.Replace(name2, "san", "si", 1)  // 找到pei替换为PE，从左到右找第一个替换
	match2 := strings.Replace(name2, "san", "SI", 2)  // 找到pei替换为PE，从左到右找前两个替换
	match3 := strings.Replace(name2, "san", "si", -1) // 找到pei替换为PE，替换所有
	fmt.Println(match1, match2, match3)               //输出：zhangsisan zhangSISI zhangsisi

	//分割
	res6 := strings.Split(name2, "s")
	// 根据`s`进行切割，获取一个切片（类似于一个数组）
	fmt.Println(res6) // [zhang,an,an]

	//拼接-低效率
	message := "hello" + "world"
	fmt.Println(message)
	//拼接-高效率
	dataList := []string{"hello", "world"}
	result := strings.Join(dataList, "")
	fmt.Println(result)
	//拼接-高高效率
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString("world")
	data := buffer.String()
	fmt.Println(data)
	//拼接-高高高效率
	var builder strings.Builder
	builder.WriteString("hello")
	builder.WriteString("world")
	value := builder.String()
	fmt.Println(value)

	//sring转int

	/****5-11布尔类型****/
	////字符串转布尔
	//res,err := strconv.ParseBool("1")
	//fmt.Println(res,err)
	////布尔转字符串
	//v2 := strconv.FormatBool(false)
	//fmt.Println(v2,reflect.TypeOf(v2))

	/****5-10decimal精确的小数计算****/
	////引入 go get github.com/shopspring/decimal
	//v1 := decimal.NewFromFloat(0.0000000000019)
	//v2 := decimal.NewFromFloat(0.02)
	////加
	//v3 := v1.Add(v2)
	//fmt.Println(v3)
	////减
	//v4 := v1.Sub(v2)
	//fmt.Println(v4)
	////乘
	//v5 := v1.Mul(v2)
	//fmt.Println(v5)
	////除
	//v6 := v1.Div(v2)
	//fmt.Println(v6)
	//
	//var price = decimal.NewFromFloat(3.4626)
	//// 保留小数点后1位（四舍五入）
	//var data1 = price.Round(1)
	//// 保留小数点后1位
	//var data2 = price.Truncate(1)
	//fmt.Println(data1, data2) // 输出：3.5  3.4

	/****5-9浮点型****/
	//var v1 float32
	//v1 = 3.14
	////默认float64
	//v2 := 1.15
	//v3 := float64(v1)+v2
	//fmt.Println(v3)

	/****5-8超大整型****/
	////定义
	//var v1 big.Int
	////指针方式-推荐
	//v2,v3:= new(big.Int),new(big.Int)
	////赋值
	//v1.SetString("123",10)
	//v2.SetInt64(123456789)
	//v3.SetInt64(987654321)
	////计算
	//res := new(big.Int)
	////整型和指针类型相加
	//res.Add(&v1,v2)
	//fmt.Println(res)
	////加
	//res.Add(v2,v3)
	//fmt.Println(res)
	////减
	//res.Sub(v2,v3)
	//fmt.Println(res)
	//// 乘
	//res.Mul(v2,v3)
	//fmt.Println(res)
	////除（地板除，只能得到商）
	//res.Div(v2,v3)
	//fmt.Println(res)
	////除，得商和余数
	//minder := new(big.Int)
	//res.DivMod(v2,v3, minder)
	//fmt.Println(res, minder)

	/****5-7初始new-指针-nil****/
	//声明指针
	//var v1 *int
	//v2 := new(int)
	//fmt.Println(v1,v2)

	/****5-6math常见函数****/
	//fmt.Println(math.Abs(-19))                // 取绝对值
	//fmt.Println(math.Floor(3.14))             // 向下取整
	//fmt.Println(math.Ceil(3.14))              // 向上取整
	//fmt.Println(math.Round(3.3478))           // 就近取整
	//fmt.Println(math.Round(3.5478*100) / 100) // 保留小数点后两位
	//fmt.Println(math.Mod(11, 3))              // 取余数，同11 % 3
	//fmt.Println(math.Pow(2, 5))               // 计算次方，如：2的5次方
	//fmt.Println(math.Pow10(2))                // 计算10次方，如：2的10次方
	//fmt.Println(math.Max(1, 2))               // 两个值，取较大值
	//fmt.Println(math.Min(1, 2))               // 两个值，取较小值

	/****5-4整型与字符串转换****/
	//v1 := 1
	////整型转字符串
	//res := strconv.Itoa(v1)
	//fmt.Println(res,reflect.TypeOf(res))
	////int8转字符串
	//var v2 int8 = 2
	//res2 := strconv.Itoa(int(v2))
	//fmt.Println(res2,reflect.TypeOf(res2))
	////字符串转整型
	//v3 := "3"
	//res3, err := strconv.Atoi(v3)
	//if err == nil {
	//	fmt.Println(res3,reflect.TypeOf(res3))
	//}else{
	//	fmt.Println("转换失败")
	//}

	/****5-3整型****/
	//var v1 int8 = 1
	//var v2 int16 = 2
	//v3 := v1+int8(v2)
	//fmt.Println(v3)
	////显示变量类型
	//fmt.Println(reflect.TypeOf(v3))

	/****5-2整型****/
	//int8(-128->127)
	//int16(-32,768->32,767)
	//int32(-2,147,483,648->2,147,483,647)
	//int64(-9,223,372,036,854,775,808->9,223,372,036,854,775,807)
	//uint8(0->255)
	//uint8(0->65,535)
	//uint8(0->4,294,967,295)
	//uint8(0->18,446,744,073,709,551,615)

	/****3-11字符串格式化****/
	//var name string
	//var age int
	//fmt.Print("请输入姓名：")
	//fmt.Scanln(&name)
	//fmt.Print("请输入年龄：")
	//fmt.Scanln(&age)
	//
	//res := fmt.Sprintf("我叫%s,今年%d岁",name,age)
	//fmt.Println(res)

	/****3-10goto掉到代码指定位置****/
	//var name string
	//fmt.Print("请输入姓名：")
	//fmt.Scanln(&name)
	//if name == "张三" {
	//	goto SVIP
	//}else if name == "李四" {
	//	goto VIP
	//}
	//
	//fmt.Println("第一步")
	//VIP:
	//fmt.Println("第二步")
	//SVIP:
	//fmt.Println("第三步")

	/****3-4for循环****/
	//for {
	//	//死循环
	//}
	//for i:=1;i<10;i++ {
	//	//结束当前循环，进行下一轮循环
	//	continue
	//	//跳出循环
	//	break
	//	fmt.Println(i)
	//}

	/****3-3switch****/
	//var v1 string
	//fmt.Print("请输入：")
	//fmt.Scanln(&v1)
	//switch v1 {
	//case "1":
	//	fmt.Println("A")
	//case "2":
	//	fmt.Println("B")
	//case "3":
	//	fmt.Println("C")
	//default:
	//	fmt.Println("其他")
	//}

	/****2-12输入****/
	//var name string
	//var age int
	////fmt.Scan实例1 需要输完所有参数，否则一直等待
	//fmt.Println("请输入用户名：")
	//fmt.Scan(&name)
	//fmt.Println(name)
	//
	////fmt.Scan实例2
	//fmt.Println("请输入用户名：")
	//res,err := fmt.Scan(&name,&age)
	////_,err := fmt.Scan(&name,&age) 省略
	//if err == nil {
	//	fmt.Println("用户输入:",name,age)
	//}else{
	//	fmt.Println("成功数量：",res,",原因:",err)
	//}
	//
	////fmt.Scanln 回车即执行
	//fmt.Println("请输入用户名：")
	//fmt.Scanln(&name,&age)
	//fmt.Println(name,age)

	//fmt.Scanf 格式化
	//fmt.Print("请输入用户名：")
	////需要输入相同内容才会显示：我叫xx,今年xx岁
	//fmt.Scanf("我叫%s 今年%d岁",&name,&age)
	//fmt.Println(name,age)

	/****2-11常量****/
	//const h = 1
	//const i int = 2
	//const (
	//	j = 3
	//)
	////iota
	//const (
	//	k1 = iota
	//	k2 = iota + 2
	//	_
	//	k3
	//	_
	//	k4
	//)
	//fmt.Println(h,i,j)
	//fmt.Println(k1,k2,k3,k4)

	/****2-9赋值和内存相关****/
	//n1 := 1
	//n2 := n1
	//var(
	//	n3 = 3
	//)
	//fmt.Println(&n1,&n2,n3)
	//
	/****2-7变量简写****/
	//d := "a"
	//var(
	//	e = "d"
	//	f = 2
	//)
	//fmt.Println(d,e,f)

	/****2-6变量****/
	//var a string =  "aaa"
	//b := "bbb"
	//fmt.Println(a)
	//fmt.Println(b)
	//var c string
	//fmt.Scanf("%s",&c)
	//fmt.Println("输入：",c)

	/****2-5数据类型****/
	//fmt.Println(1+1-1*1/1%1)
	//fmt.Println("a"+"b"+"c")
	//fmt.Println(1>2)

	/****2-3输出****/
	//带换行输出
	//fmt.Println("hello world")
	//fmt.Println("hello world2")
	////中间自动加空格
	//fmt.Println("hello","world","!")
	////格式化输出
	//fmt.Printf("%s开车去%s","老张","东北，撞啦\n")
	//fmt.Printf("整数：%d\n小数：%f\n两位小数：%.2f",100,100.0012,100.123)
}
