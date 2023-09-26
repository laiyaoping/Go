package main

// 基础语法

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// 结构体
type user struct {
	name     string
	password string
}

type point struct {
	x, y int
}

// JOSN操作，需要保证每个字段的第一个字母是大写
type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

// 结构体方法
func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

// golang的函数有两个返回值，一个是要返回的值，另一个是返回的错误信息
func exite(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func add3(n int) {
	n += 2
}

func add3ptr(n *int) {
	*n += 2
}

// 错误处理
func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("Not Found")
}

func main() {
	// 变量和常量命名

	var a = "lyp"

	var b, c int = 1, 2

	var d = true

	var e float64 = 0.3

	f := float32(e)

	g := a + "xwq"
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g)

	//fmt.Scanf("%s", &a)
	//fmt.Println(a)

	fmt.Println("HelloWorld")

	const s string = "constant"
	const h = 500000
	const i = 3e20 / h
	fmt.Println(s, h, math.Sin(h), math.Sin(i))

	//if else语句
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 10 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//循环,go语言只有for循环
	ii := 1
	for j := 1; j < 9; j++ {
		fmt.Println(j)
	}

	for ii <= 3 {
		fmt.Println(ii)
		ii = ii + 1
	}

	//switch
	aa := 2
	switch aa {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5, 6, 7, 8:
		fmt.Println("between four to eight")
	default:
		fmt.Println("others")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//数组
	var num [5]int
	num[4] = 100
	fmt.Println(num[4], len(num))

	numb := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numb)

	var numc [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			numc[i][j] = i + j
			//fmt.Print(numc[i][j])
		}
	}
	fmt.Println("2d: ", numc)

	//切片(slice)
	ss := make([]string, 3)
	ss[0] = "a"
	ss[1] = "b"
	ss[2] = "c"
	fmt.Println("get ", s[2])
	fmt.Println("len:", len(ss))

	//往slice中增加元素，当容量不足时会扩容，扩容会返回新的slice,所以增加元素时要把增加元素后的slice返回回去。golang中的slice存储了长度+容量+指向数组的指针
	ss = append(ss, "d")
	ss = append(ss, "e", "f")
	fmt.Println(s)

	cc := make([]string, len(ss))
	copy(cc, ss)
	fmt.Println(cc)
	ss[0] = "hello_go"
	fmt.Println(ss)
	fmt.Println(cc)
	//可见go中的copy()函数时浅拷贝

	fmt.Println(ss[2:5])
	fmt.Println(ss[:5])
	fmt.Println(ss[2:])

	//map
	mmap := make(map[string]int)
	mmap["one"] = 1
	mmap["two"] = 2
	fmt.Println(mmap)
	fmt.Println(len(mmap))
	fmt.Println(mmap["one"])
	fmt.Println(mmap["un"])

	r, ok := mmap["nn"]
	fmt.Println(r, ok)

	delete(mmap, "one")
	fmt.Println(mmap)

	mmap2 := map[string]int{"one": 1, "two": 2}
	fmt.Println(mmap2)

	//range
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num)
		}
	}
	fmt.Println(sum)

	mmap3 := map[string]string{"a": "A", "b": "B"}
	for kkey, vvalue := range mmap3 {
		fmt.Println(kkey, vvalue)
	}
	for kkey := range mmap3 {
		fmt.Println("key", kkey)
	}

	//函数调用
	res := add(1, 2)
	fmt.Println(res)

	vv, ok := exite(map[string]string{"a": "A"}, "a")
	fmt.Println(vv, ok)

	//指针
	n := 5
	add3(n)
	fmt.Println(n)

	add3ptr(&n)
	fmt.Println(n)

	//结构体
	strua := user{name: "lyp", password: "123456"}
	strub := user{"wang", "1024"}
	struc := user{name: "wang"}
	var strud user
	strud.name = "xie"
	strud.password = "654321"

	fmt.Println(strua, strub, struc, strud)
	fmt.Println(strua.password == "123456")

	//结构体方法
	strua.resetPassword("2048")
	fmt.Println(strua.checkPassword("2048"))

	u, err := findUser([]user{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u.name)
	}

	if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u.name)
	}

	//字符串格式化
	strs := "hello"
	nn := 123
	poi := point{1, 2}
	fmt.Println(strs, nn) //hello 123
	fmt.Println(poi)

	fmt.Printf("s=%v\n", strs)
	fmt.Printf("s=%v\n", nn)
	fmt.Printf("s=%v\n", poi)
	fmt.Printf("p=%+v\n", poi)
	fmt.Printf("p=%#+v\n", poi)

	ff := 3.1415926535
	fmt.Println(ff)
	fmt.Printf("%.2f\n", ff)

	//JSON处理
	userinfo := userInfo{Name: "lai", Age: 18, Hobby: []string{"basketball", "Golang"}}
	//序列化userinfo
	buf, err := json.Marshal(userinfo)
	if err != nil {
		panic(err)
	}
	//不强制类型转换成string类型则会输出一串数字编码
	fmt.Println(buf)         //[123 34 78 97 ...]
	fmt.Println(string(buf)) //{"Name":"lai","age":18,"Hobby":["basketball","Golang"]}

	buf, err = json.MarshalIndent(userinfo, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var userinfo2 userInfo
	//反序列化
	err = json.Unmarshal(buf, &userinfo2)
	if err != nil {
		panic(err)
	}
	//格式化输出
	fmt.Printf("%#v\n", userinfo2) //main.userInfo{Name:"lai", Age:18, Hobby:[]string{"basketball", "Golang"}}

	//时间处理
	now := time.Now()
	fmt.Println(now)

	t1 := time.Date(2023, 9, 25, 21, 40, 36, 0, time.UTC)
	t2 := time.Date(2023, 9, 25, 23, 10, 36, 0, time.UTC)

	fmt.Println(t1)
	fmt.Println(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute())
	fmt.Println(t1.Format("2006-01-02 15:04:05"))
	diff := t2.Sub(t1)
	fmt.Println(diff)
	fmt.Println(diff.Minutes(), diff.Seconds())
	t3, err := time.Parse("2006-01-02 15:04:05", "2023-09-25 21:40:36")
	if err != nil {
		panic(err)
	}
	fmt.Println(t3 == t1)
	//转换成时间戳
	fmt.Println(now.Unix())

	//数字解析(字符串转数字)
	fff, _ := strconv.ParseFloat("1.234", 64) //1.234
	fmt.Println(fff)

	nnn, _ := strconv.ParseInt("111", 10, 64) //111
	fmt.Println(nnn)

	nnn, _ = strconv.ParseInt("0x1000", 0, 64) //4096
	fmt.Println(nnn)

	n2, _ := strconv.Atoi("123") //123
	fmt.Println(n2)

	n2, err = strconv.Atoi("AAA") //0 strconv.Atoi: parsing "AAA": invalid syntax
	fmt.Println(n2, err)

	//进程信息
	fmt.Println(os.Args)

	//获取环境变量
	fmt.Println(os.Getenv("PATH"))
	//写入环境变量
	//	fmt.Println(os.Setenv("AA", "BB"))

	//获取子进程
	buff, err := exec.Command("grep", "127.0.0.1", "/etc/hosts").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buff))
}
