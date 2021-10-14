package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"os/exec"
)

//全局变量
//var aa string = "aa"

//func changeData(ptr * string){
//	*ptr = "2";
//}

func sendEmail(types int, email string, exec f1) (int, string) {
	v1, v2 := exec(types)
	fmt.Println("发送函数")
	fmt.Println(types)
	fmt.Println(email)
	fmt.Println(v1, v2)
	return types, email
}

type f1 func(arg int) (int, string)

func f2(arg int) (int, string) {
	return 1, "a"
}

func f3(num ...int) int {
	fmt.Println(num)
	return 123
}

func f4() func(num ...int) int {
	return f3
}

func do() int {
	fmt.Println("aaaa")
	//defer:函数结束时执行,多个倒叙执行
	defer fmt.Println("bbbb")
	defer f3(1)
	fmt.Println("cccc")
	return 1
}

type Person struct {
	name string
	age  int
}

var p Person = Person{
	name: "a",
	age:  1,
}

func doSome() *Person {
	return &p
}

//声明类型
type myInt int

//指针方法
func (i *myInt) doSomee(n1 int, n2 int) int {
	return n1 + n2 + int(*i)
}

/****7-17再谈结构体：结构体工厂start****/
//首字母开头小写表示私有，大写表示公有
type g1 struct {
	name string
	age  int
}

type g2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func newG1(name string, age int) *g1 {
	return &g1{name: name, age: age}
}

func push() {
	fmt.Print("hello world")
}

/****7-17再谈结构体：结构体工厂end****/
/****7-18接口：定义和作用start****/
//接口
type Base interface {
	ff1()
	ff2() int
	ff3(int, bool)
	ff4(n1 int, n2 bool) int
}

//空接口
type empty interface{}

func emptyFunc(n1 interface{}) {
	//接口转换类型
	tp, ok := n1.(g1)
	if ok {
		fmt.Println(tp.name, tp.age)
	} else {
		fmt.Println(tp, ok)
		fmt.Println("转换失败")
	}
	fmt.Println(n1)
}

/****7-18接口：定义和作用end****/

/****7-18接口：练习题start****/
type IMassage interface {
	send() bool
}

type Email struct {
	email   string
	content string
}

func (p *Email) send() bool {
	fmt.Println("发送邮箱提醒", p.email, p.content)
	return true
}

type Wechat struct {
	email   string
	content string
}

func (p *Wechat) send() bool {
	fmt.Println("发送微信提醒", p.email, p.content)
	return true
}

func something(objectList []IMassage) {
	for _, item := range objectList {
		result := item.send()
		fmt.Println(result)
	}
}
func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

/****7-18接口：练习题end****/
func main() {
	/****8-2内置包：执行命令****/
	//var out bytes.Buffer
	//cmdPrt := exec.Command("xxx/go","build","-h","127.0.0.1")
	//cmdPrt.Dir = "../../" //进入目录
	//cmdPrt.Stdout = &out
	//cmdPrt.Run()
	//实时输出
	cmd := exec.Command("ping", "baidu.com")
	fmt.Println(cmd.Args)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		//转码
		response := ConvertToByte(string(line), "gbk", "utf8")
		fmt.Print(string(response))
	}
	cmd.Wait()
	/****8-2内置包：文件和路径相关****/
	//创建单级目录，已存在会报错
	//err := os.Mkdir("x2",0755)
	//fmt.Println(err)
	////创建多级目录，已存在不报错
	//errAll := os.MkdirAll("x3/z1/c2",0755)
	//fmt.Println(errAll)
	////方法二
	//if err2 := os.Mkdir("x2",0755); err2 == nil{
	//	fmt.Println("创建成功")
	//}else{
	//	fmt.Println("创建失败")
	//}

	//if os.Mkdir("x2",0755) == nil{
	//	fmt.Println("创建成功")
	//}else{
	//	fmt.Println("创建失败")
	//}
	//判断路径是否存在
	//_,err := os.Stat("x3/z1/c2")
	//if err != nil {
	//	if os.IsNotExist(err){
	//		fmt.Println("文件夹不存在")
	//	}
	//}

	/****8-2内置包****/
	//json
	//v1 := []interface{}{
	//	"张三",
	//	1,
	//	true,
	//	4.13,
	//	map[string]interface{}{
	//		"name":"李四",
	//		"age":5,
	//	},
	//	g2{ //参数名需要是大写，才能转换成json
	//		Name: "zz",
	//		Age:  1,
	//	},
	//}
	//fmt.Println(v1)
	//res,_ := json.Marshal(v1)
	//jsonData := string(res)
	//fmt.Println(jsonData)
	//var value []interface{}
	//json.Unmarshal([]byte(jsonData),&value)
	//fmt.Println(value)
	////time
	////获取当前时间
	//t1 := time.Now()
	//fmt.Println(t1)
	////格式化
	//dateString := t1.Format("2006/01/02 15:04:05")
	//fmt.Println(dateString)
	////创建文件夹
	//os.Mkdir(dateString,0755)
	////获取当前utc时间
	//t2 := time.Now().UTC()
	//fmt.Println(t2)
	////转换成time类型
	//t3,_ := time.Parse("2006/01/02 15:04:05","2021/10/11 10:11:25")
	//fmt.Println(t3)
	////时间增加
	//t6 := t1.Add(time.Hour * 1) //当前时间+1小时
	//fmt.Println(t6)
	////时间减少
	//t7 := t1.Add(-time.Minute * 1) //当前时间-1分钟
	//fmt.Println(t7)
	//t8 := t1.Sub(t6)
	//fmt.Println("时间间隔:",t8)
	//fmt.Println("时间间隔小时:",t8.Hours())
	//fmt.Println("时间间隔分钟:",t8.Minutes())
	//fmt.Println("时间间隔秒:",t8.Seconds())
	////获取时间戳
	//fmt.Printf("时间戳（秒）：%v\n", time.Now().Unix())
	//fmt.Printf("时间戳（纳秒）：%v\n",time.Now().UnixNano())
	//fmt.Printf("时间戳（毫秒）：%v\n",time.Now().UnixNano() / 1e6)
	//fmt.Printf("时间戳（纳秒转换为秒）：%v\n",time.Now().UnixNano() / 1e9)

	////flag 往脚本里传值
	//host := flag.String("h","127.0.0.1","主机名")
	//port := flag.Int("p",8080,"端口")
	//
	//var email string
	//flag.StringVar(&email,"e","","邮箱")
	//flag.Parse()
	//fmt.Println(*host,*port,email)

	//正则表达式
	//匹配
	//m2,_ := regexp.MatchString("foo.*","seafoo2")
	//fmt.Println(m2)
	////查找
	//reg1 := regexp.MustCompile(`\d{4}`)
	//ret1 := reg1.FindString("515152as12")
	//ret2 := reg1.FindAllString("12312as214214",-1)
	//fmt.Println(ret1,ret2)

	/****7-18接口：练习题****/
	//v1 := &Email{
	//	email:   "111",
	//	content: "2222",
	//}
	//v1.send()
	//fmt.Println(*v1)

	//messageObjList := []IMassage{
	//	&Email{
	//		email:   "a",
	//		content: "1",
	//	},
	//	&Wechat{
	//		email:   "b",
	//		content: "2",
	//	},
	//}
	//something(messageObjList)

	/****7-18接口：定义和作用****/
	//v1 := make([]empty,0)
	//v1 = append(v1,"张三")
	//v1 = append(v1,2)
	//v1 = append(v1,true)
	////定义空接口-简写
	//v2 := make([]interface{},0)
	//fmt.Println(v1,v2)
	//
	//emptyFunc(153)
	//emptyFunc(g1{name: "123",age:  0})

	/****7-17再谈结构体：结构体工厂****/
	//v1 := g1{name: "a",age:  2,}
	//fmt.Println(v1)
	//v2 := newG1("b",3)
	//fmt.Println(*v2)
	/****7-15再谈结构体：类型方式****/
	//var v1 myInt = 123
	//v2:= v1.doSomee(1,2)
	//fmt.Println(v2)
	/****7-14再谈结构体：返回值拷贝问题****/
	//data := doSome()
	//p.name = "b"
	//fmt.Println(*data)
	//fmt.Println(p)

	/****7-13函数：defer和自执行函数****/
	//延迟执行
	//ret := do()
	//fmt.Println(ret)
	////自执行函数
	//func(n1 int) int {
	//	fmt.Println(n1)
	//	return n1
	//}(123)

	/****7-12函数：闭包****/
	//fun := func(arg int) func() {
	//	return func() {
	//		fmt.Println(arg)
	//	}
	//}
	//var funList []func()
	//for i := 0; i < 5; i++ {
	//	v2 := fun(i)
	//	funList = append(funList,v2)
	//}
	//funList[0]()
	//funList[1]()
	//funList[4]()

	/****7-9函数****/
	//v1,v2 := sendEmail(1,"1.1@1",f2)
	//fmt.Println(v1,v2)
	////变长参数
	//f3(1,2,3)
	//返回函数
	//v3 := f4()
	//v4 := v3(1)
	//fmt.Println(v4)
	////匿名函数
	//v5 := func(n1 int,n2 int) int {
	//	return 1
	//}
	//fmt.Println(v5(1,2))
	////匿名函数-直接执行
	//v6 := func(n1 int,n2 int) int {
	//	return 123
	//}(11,22)
	//fmt.Println(v6)

	/****7-8结构体:练习题****/
	//type v1 struct {
	//	name string "姓名"
	//	age int "年龄"
	//}
	//type v2 struct {
	//	title string "班级"
	//	//count int "人数"
	//	v1
	//}
	//
	//var list []v2
	//for {
	//	var name,title string
	//	var age int
	//	fmt.Print("请输入班级：")
	//	fmt.Scan(&title)
	//	if title == "Q" {
	//		break
	//	}
	//	fmt.Print("请输入姓名：")
	//	fmt.Scan(&name)
	//	fmt.Print("请输入性别：")
	//	fmt.Scan(&age)
	//	//count := v2.count
	//	list = append(list,v2{title: title,v1:v1{name:name,age:age}})
	//
	//}
	//fmt.Println(list)

	/****7-6结构体:标签****/
	//type v1 struct {
	//	name string "姓名"
	//	age int "年龄"
	//}
	//p1 := v1{name: "李四",age:15}
	////获取标签名，方法一
	//plType := reflect.TypeOf(p1)
	//fmt.Println(p1)
	//f1 := plType.Field(0)
	//fmt.Println(f1)
	//fmt.Println(f1.Tag)
	//fmt.Println(f1.Type)
	////获取标签名，方法二
	//f2,_ := plType.FieldByName("name")
	//fmt.Println(f2.Tag)
	////获取标签名，循环获取
	//fieldNum := plType.NumField() //总共多少个字段
	//for i := 0; i < fieldNum; i++ {
	//	fmt.Println(plType.Field(i).Tag,plType.Field(i).Name)
	//}

	/****7-6结构体:嵌套赋值拷贝****/
	//type v1 struct {
	//	name string
	//	age int
	//}
	//type v2 struct {
	//	name string
	//	v1
	//}
	//p1 := v2{name: "李四",v1:v1{"张三1",15}}
	//p2 := p1
	//fmt.Println(p1,p2)
	//p2.v1.name = "张三2"
	//fmt.Println(p1,p2)
	////拷贝
	//type v3 struct {
	//	name string
	//	age int
	//	hobby [2]string
	//	num *[]int
	//	parent map[string]int
	//	v1
	//}
	//p3 := v3{
	//	name:   "张三",
	//	age:    12,
	//	hobby:  [2]string{"a","b"},
	//	num:    &[]int{1,2}, //拷贝的是内存地址
	//	parent: map[string]int{"a":1,"b":2}, //拷贝的是内存地址
	//	v1:     v1{"李四",15},
	//}
	//p4 := p3
	//fmt.Println(p3)
	//fmt.Println(p4)
	//p3.hobby[0] = "c"
	//p3.num = &[]int{2,3}
	//p3.parent["a"] = 3
	//fmt.Println(p3)
	//fmt.Println(p4)
	//fmt.Println(*p3.num)
	//fmt.Println(*p4.num)

	/****7-5结构体:赋值拷贝****/
	////结构体
	//type v1 struct {
	//	name string
	//	age int
	//}
	//p1 := v1{name: "张三",age: 18}
	//p2 := p1 //重新拷贝一份
	//fmt.Println(p1,p2) //{张三 18} {张三 18}
	//p2.name = "李四"
	//fmt.Println(p1,p2) //{张三 18} {李四 18}
	////结构体指针
	//p3 := &v1{name: "张三",age: 18}
	//p4 := p3 //指向同一地址
	//fmt.Println(p3,p4)
	//p4.name = "李四"
	//fmt.Println(p3,p4)

	/****7-3结构体:指针****/
	//type v1 struct {
	//	name string
	//	age int
	//}
	////初始化结构体指针
	//p2 := &v1{name: "张三",age: 18}
	//fmt.Println(p2.name,p2.age)

	/****7-3结构体:初始化****/
	//type v1 struct {
	//	name string
	//	age int
	//	hobby []string
	//}
	////方法一
	//p1 := v1{"张三1",15,[]string{"a","b"}}
	////方法二
	//p2 := v1{name: "张三2",age: 16,hobby: []string{"a","b"}}
	////方法三
	//var p3 v1
	//p3.name = "张三3"
	//p3.age = 17
	//p3.hobby = []string{"a","b"}
	//fmt.Println(p1,p2,p3)
	////结构体嵌套
	//type v2 struct {
	//	name string
	//	v1
	//}
	//p4 := v2{name: "李四",v1:v1{"张三1",15,[]string{"a","b"}}}
	//fmt.Println(p4)

	/****7-2结构体:定义****/
	//type p1 struct {
	//	name,email string
	//	age int
	//}
	//type p2 struct {
	//	name,email string
	//	age int
	//	p1 p1 //嵌套结构体
	//}
	//type p3 struct {
	//	name,email string
	//	age int
	//	p1 //匿名字段（和名字一直）
	//}

	/****7-1结构体****/
	////定义
	//type Person struct {
	//	name string
	//	age int
	//	email string
	//}
	////结构体
	//v1 := Person{"张三",18,"1@1.com"}
	////取值
	//fmt.Println(v1.name,v1.age,v1.email)
	//v1.age = 19
	//fmt.Println(v1.name,v1.age,v1.email)

	/****6-24指针：指针的指针****/
	//v1 := 1
	//var v2 *int = &v1
	//var v3 **int = &v2
	//var v4 ***int = &v3
	//*v2 = 2
	//**v3 = 3
	//***v4 = 4
	//fmt.Println(v1,v2,v3,v4)

	/****6-22指针****/
	////场景一：需要同步修改时就用指针
	//var v1 *int
	//v2 := 12
	//v1 = &v2
	//fmt.Println(v1,*v1)
	//v2 = 13
	//fmt.Println(v1,*v1)
	////场景2：函数同步修改
	//v3 := "1"
	//changeData(&v3)
	//fmt.Println(v3)
	////场景3
	//var v4 string
	//fmt.Print("请输入用户名：")
	//fmt.Scanf("%s",&v4)
	//if v4 == "张三" {
	//	fmt.Println("Yes")
	//}else{
	//	fmt.Println("No")
	//}

	/****6-15map：变量赋值****/
	//v1 := map[string]string{"a":"b","c":"d"}
	//v2 := v1 //会指向同一地址
	//v1["a"] = "abc"
	//fmt.Println(v2) //输出：map[a:abc c:d]
	//fmt.Println(v2) //输出：map[a:abc c:d]

	/****6-13map：增删改查****/
	////增加
	//v1 := map[int]string{1:"a",2:"b"}
	//v1[3] = "c"
	//fmt.Println(v1)
	////修改
	//v1[2] = "d"
	//fmt.Println(v1)
	////删除
	//delete(v1,1)
	//fmt.Println(v1)
	////查看
	//for i,v := range v1 {
	//	fmt.Println(i,v)
	//}
	////嵌套
	//q1 := make(map[string]int)
	//q1["a"] = 1
	//q2 := make(map[string]string)
	//q2["a"] = "b"
	//q4 := make(map[string][2]int)
	//q4["a"] = [2]int{1,2}
	//q5 := make(map[string][]int)
	//q5["a"] = []int{1,2,3}
	//q6 := make(map[string]map[string]int)
	//q6["a"] = map[string]int{"a":1,"b":2}
	//q7 := make(map[string][2]map[string]string)
	//q7["a"] = [2]map[string]string{{"a":"b","c":"d"},{"b":"c"}}
	//fmt.Println(q1,q2,q4,q5,q6,q7)
	//q8 := make(map[int]int)
	//q8[1] = 2
	//q9 := make(map[string]int)
	//q9["a"] = 2
	//q10 := make(map[float32]int)
	//q10[3.2] = 2
	//q11 := make(map[bool]int)
	//q11[true] = 2
	//q12 := make(map[[2]int]int)
	//q12[[2]int{1,2}] = 2
	//q13 := make(map[[]int]int) //不可以是切片
	//q14 := make(map[map[int]int]int) //不可以是map

	/****6-10map：常见操作****/
	////获取map长度
	//v1 := map[string]string{"a":"11","b":"22"}
	//fmt.Println(v1,len(v1))
	//v2 := make(map[string]string,10)
	//v2["a"] = "11"
	//v2["b"] = "22"
	//fmt.Println(v2,len(v2))

	/****6-10map：初始化****/
	//v1 := map[string]string{"a":"1","b":"2"}
	//v2 := map[int]string{1:"a",2:"b"}
	//v3 := make(map[int]int)
	//v4 := make(map[string]string,10)
	//v3[1] = 2
	//v4["a"] = "b"
	//fmt.Println(v1,v2,v3,v4)
	//
	////声明
	//v5 := new(map[string]string)
	////指针不能直接赋值
	//v5 = &v4
	//fmt.Println(v5)
	//
	////键是数组
	//v7 := make(map[[2]int]float32)
	//v7[[2]int{1,1}] = 1.5
	//v7[[2]int{1,2}] = 1.4
	//v7[[2]int{1,2}] = 1.8
	//fmt.Println(v7)
	//
	//v8 := make(map[[2]int][3]string)
	//v8[[2]int{1,1}] = [3]string{"a","b","c"}

	/****6-6切片：嵌套****/
	//v1 := []int{1,2,3}
	//v2 := [][]int{[]int{1,2},[]int{3,4}}
	//v3 := [][2]int{[2]int{1,2},[2]int{3,4}}
	////fmt.Println(v1,v2,v3)
	//v1 := []int{1,2,3}
	//v2 := [3]int{1,2,3}
	//v3 := [...]int{1,2,3}
	//fmt.Println(v1,v2,v3)
	//fmt.Println(reflect.TypeOf(v1),reflect.TypeOf(v2),reflect.TypeOf(v3))
	/****6-5切片：常见操作****/
	////切片
	//v2 := make([]int,3,3)
	//v2[0] = 1
	//v2[1] = 2
	//v2[2] = 3
	//fmt.Println(v2[0:3]) //从索引0开始到3（不包括） 1,2,3
	//fmt.Println(v2[1:]) //索引1之后的 1,2
	//fmt.Println(v2[:2]) //索引2之前的 1,2

	////追加
	//a1 := []int{1,2,3}
	//a2 := append(a1,4)
	//a3 := append(a1,4,5)
	//a4 := append(a1,[]int{4,5,6}...) //加三个点，拿到新切片中的元素
	//fmt.Println(a1,a2,a3,a4)

	////删除元素“3”
	//d1 := []int{1,2,3,4,5,6,7}
	////“3”的索引2
	//index := 2
	////分割追加
	//d2 := append(d1[0:index],d1[index+1:]...) //执行完v1会被覆盖
	//fmt.Println(d2)
	//fmt.Println(d1)

	////插入
	//s1 := []int{1,2,3,4,5}
	//index := 2
	//s2 := make([]int,0,len(s1)+1)
	//s2 = append(s2,s1[:index]...)
	//s2 = append(s2,100)
	//s2 = append(s2,s1[index:]...)
	//fmt.Println(s2)

	////循环
	//f1 := []int{1,2,3,4,5}
	//for i:=0;i<len(f1);i++ {
	//	fmt.Println(i,f1[i])
	//}
	//for i,v := range f1 {
	//	fmt.Println(i,v)
	//}

	/****6-4切片：自动扩容****/
	//v1 := make([]int,1,3)
	////len获取长度，cap获取容量
	//fmt.Println(len(v1),cap(v1))
	//
	////追加
	//v1 = append(v1,66)
	//v2 := append(v1,99)
	//fmt.Println(v1)
	//fmt.Println(v2)
	//
	////扩容
	//v3 := []int{11,22,33}
	//v4 := append(v3,44)
	//fmt.Println(len(v3),cap(v3),v3)
	//fmt.Println(len(v4),cap(v4),v4)

	/****6-3切片：初始化和声明****/
	////创建切片1
	//var nums []int
	////创建切片2
	//var data = []int{1,2,3}
	////创建切片3,make只用于切片、字典、channel
	//var users = make([]int,2,3) //长度2，容量3
	//fmt.Println(nums)
	//fmt.Println(data)
	//fmt.Println(users)

	/****5-19数组：多层嵌套****/
	//var arr [2][3][4]int
	//arr[1][2][3] = 12
	//arr[0][1] = [4]int{1,2,3}
	//fmt.Println(arr) //[[[0 0 0 0] [1 2 3 0] [0 0 0 0]] [[0 0 0 0] [0 0 0 0] [0 0 0 12]]]
	//
	////声明并赋值
	//arr2 := [2][3][2]int{{{1,2},{1,2},{1,2}},{{3,4},{3,4},{3,4}}}
	//arr2 := [2][3]int{[3]int{1,2,3},[3]int{4,5,6}}
	//fmt.Println(arr2)

	/****5-18数组：长度、索引、切片和循环****/
	////长度
	//name := [2]string{"张三","李四"}
	//fmt.Println(len(name))
	////索引
	//data := name[0]
	//fmt.Println(data)
	//name[0] = "王五"
	//fmt.Println(name)
	////切片
	//nums := [3]int32{11,22,33}
	//v1 := nums[1:3]
	//fmt.Println(v1)
	////循环
	//for i:=0;i<len(name);i++ {
	//	fmt.Println(i,name[i])
	//}
	//for index,item := range name {
	//	fmt.Println(index,item)
	//}

	/****5-17数组可变拷贝的特性****/
	////可变
	//names := [2]string{"张三","李四"}
	//names[1] = "王五"
	//fmt.Println(names)
	////拷贝
	//name1 := [2]string{"张三","李四"}
	//name2 := name1
	//name1[1] = "王五"
	//fmt.Println(name2)

	/****5-15数组****/
	////先声明再赋值，方式一
	//var num1 [3]int
	//num1[0] = 123
	//num1[1] = 456
	//num1[2] = 789
	//fmt.Println(num1[1])
	////先声明再赋值，方式二
	//var num2 = [2]string{"张三","李四"}
	//fmt.Println(num2)
	////声明加赋值,3不写或者改成...则省略个数
	//var num3 = []int{0:123,1:456,2:789}
	//fmt.Println(num3)
	//var num4 = [...]string{0:"张三",2:"李四"}
	//fmt.Println(num4)

	////声明指针类型的数组,不会开辟内存
	//var ber1 *[3]int
	//ber1[0] = 123
	//fmt.Println(ber1)
	//
	//ber2 := new([3]int)
	//ber2[2] = 789
	//fmt.Println(ber2)

	/****5-14索引切片和循环****/
	////索引获取字节
	//name := "张三"
	//v1 := name[0]
	//fmt.Println(v1)
	//
	////切片获取字节区间 0:3张 3:6三
	//v2 := name[3:6]
	//fmt.Println(v2)
	//
	////for range
	//for index,item := range name{
	//	fmt.Println(index,item,string(item))
	//}
	//
	////rune集合
	//dataList := []rune(name)
	//fmt.Println(dataList)
	//fmt.Println(dataList[0],string(dataList[0]))
	//fmt.Println(dataList[1],string(dataList[1]))

	/****5-13字符串常见功能****/
	//name := "张三三"
	//name2 := "zhangsansan"
	//name3 := "ZHANGSANSAN"
	////获取字节长度
	//fmt.Println(len(name)) //输出：9
	////获取字符长度
	//fmt.Println(utf8.RuneCountInString(name)) //输出：3
	//
	////判断是否以“x”开头
	//res1 := strings.HasPrefix(name, "张")
	//fmt.Println(res1) // 输出：true
	////判断是否以“x”结尾
	//res2 := strings.HasSuffix(name, "三")
	//fmt.Println(res2) // 输出：true
	////判断是否包含“x”
	//res3 := strings.Contains(name, "三")
	//fmt.Println(res3) // 输出：true
	//
	////转大写
	//res4 := strings.ToUpper(name2)
	//fmt.Println(res4) // 输出：ZHANGSANSAN
	////转小写
	//res5 := strings.ToLower(name3)
	//fmt.Println(res5) // 输出：zhangsansan
	//
	////去两边
	//result1 := strings.TrimRight(name2, "an") // 去除右边的an
	//result2 := strings.TrimLeft(name2, "zh")  // 去除左边的zh
	//result3 := strings.Trim(name2, "n")       // 去除两边的n
	//fmt.Println(result1, result2, result3)    // 输出：zhangsans angsansan zhangsansa
	//
	////替换
	//match1 := strings.Replace(name2, "san", "si", 1)  // 找到pei替换为PE，从左到右找第一个替换
	//match2 := strings.Replace(name2, "san", "SI", 2)  // 找到pei替换为PE，从左到右找前两个替换
	//match3 := strings.Replace(name2, "san", "si", -1) // 找到pei替换为PE，替换所有
	//fmt.Println(match1, match2, match3)               //输出：zhangsisan zhangSISI zhangsisi
	//
	////分割
	//res6 := strings.Split(name2, "s")
	//// 根据`s`进行切割，获取一个切片（类似于一个数组）
	//fmt.Println(res6) // [zhang,an,an]
	//
	////拼接-低效率
	//message := "hello" + "world"
	//fmt.Println(message)
	////拼接-高效率
	//dataList := []string{"hello", "world"}
	//result := strings.Join(dataList, "")
	//fmt.Println(result)
	////拼接-高高效率
	//var buffer bytes.Buffer
	//buffer.WriteString("hello")
	//buffer.WriteString("world")
	//data := buffer.String()
	//fmt.Println(data)
	////拼接-高高高效率
	//var builder strings.Builder
	//builder.WriteString("hello")
	//builder.WriteString("world")
	//value := builder.String()
	//fmt.Println(value)

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
