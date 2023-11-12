package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 一系列变量赋值方法
	// Java 的字符串由字符 char 组成，但 Go 的字符串由字节 byte 组成
	var head string
	head = "Micheal"
	var name = "Kimi"
	// 整形的默认值为 0
	var i int
	i = 10
	var j float64
	j = 12.33556
	fmt.Println("Hello Go! ", name, "; head is ", head)
	fmt.Println("i = ", i, "; j = ", j)

	// 一次赋值多个变量
	var a, b, c int
	a = 10
	b = 15
	c = 20
	fmt.Println("a = ", a, "; b = ", b, "; c = ", c)

	fmt.Println("url: ", url, "; username: ", username, "; password: ", password)

	// 布尔类型变量，注意默认值为 false！
	var flag bool
	fmt.Println("default flag value: ", flag)
	flag = true
	fmt.Println("new flag value: ", flag)

	var data int32 = 9999
	// 基本数据类型之间的互相转化
	transferDataType(data)
	//基本数据类型 <--> string
	formatString()
	// 指针
	ptr()
	/** Go根据值和引用可分为两种类型：
	值类型：基本数据类型;
	引用类型：指针 ptr，map，管道 chan，interface
	*/

}

// 全局变量
var url, username, password = "https://www.baidu.com", "Jason", 12345

func transferDataType(i int32) {
	var n1 = float64(i) // 低精度转高精度
	var n2 = int8(i)    // 大转小时，容易出现溢出！
	var n3 = int64(i)   // 小转大，不会溢出
	fmt.Printf("n1: %v; n2: %v; n3: %v\n", n1, n2, n3)
}

func formatString() {
	// 可以使用 Sprintf 基本类型转 string
	var n1 = 99
	var str string
	str = fmt.Sprintf("%d", n1)
	fmt.Println("Sprintf formatString", str)

	var n2 = 17.483563
	str = fmt.Sprintf("%f", n2)
	fmt.Println("Sprintf formatString", str)

	var flag bool
	flag = true
	str = fmt.Sprintf("%t", flag)
	fmt.Println("Sprintf formatString", str)

	// 或者使用 strconv 函数，使基本数据类型转string
	str = strconv.FormatInt(int64(n1), 10) // 10进制
	fmt.Println("strconv FormatInt", str)

	str = strconv.FormatInt(int64(n1), 2) // 2进制
	fmt.Println("strconv FormatInt", str)

	str = strconv.FormatFloat(n2, 'f', 10, 64)
	fmt.Println("strconv FormatFloat", str)

	// strconv 函数也可以使 string 转基本数据类型
	str = "true"
	flag, _ = strconv.ParseBool(str)
	fmt.Println("strconv ParseBool", flag)

	str = "12345"
	var n3 int64
	n3, _ = strconv.ParseInt(str, 10, 64)
	fmt.Println("strconv ParseInt", n3)

	str = "12.45"
	var n4 float64
	n4, _ = strconv.ParseFloat(str, 64)
	fmt.Println("strconv ParseFloat", n4)
}

func ptr() {
	// 指针
	var num = 10
	var ptr *int = &num
	fmt.Println("ptr 地址：", ptr, "值：", *ptr)

	// 通过指针修改指向地址的值
	*ptr = 99
	fmt.Println("ptr 地址：", ptr, "值：", *ptr)

	var str = "Jason"
	var ptr2 *string = &str
	fmt.Println("ptr2 地址：", ptr2, "值：", *ptr2)
}
