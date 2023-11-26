package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * 本章主要讲 string 的多重方法
 */

func main() {
	// (1) 遍历 String
	var str string
	str = "Hello 北京"
	// 如果有中文，不能直接遍历，应先使用 rune 切片
	var r = []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}
	fmt.Println()
	// 或者我们可以使用 k, v键值对进行遍历
	for _, v := range str {
		fmt.Printf("%c", v)
	}
	fmt.Println()
	// 错误的遍历方式，这种方法仅适用于英文
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()

	// (2) String 与 整数 互转，使用 strconv 函数
	var origin string = "123"
	var number int
	number, _ = strconv.Atoi(origin)
	fmt.Println("number: ", number)
	origin = strconv.Itoa(number)
	fmt.Println("origin: ", origin)

	// (3) 查找Contains、统计Count、比较EqualFold、定位Index指定字符串，以及replace、split、ToUpper、ToLower函数
	var ifExist bool = strings.Contains("Hello Shanghai", "Shanghai")
	fmt.Printf("ifExist: %v\n", ifExist)

	var count int = strings.Count("Hello Shanghai", "h")
	fmt.Printf("count: %v\n", count)

	var ifSame bool = strings.EqualFold("abc", "ABC") // 不区分大小写的比较
	fmt.Printf("ifSame: %v\n", ifSame)

	var index int = strings.Index("Hello Shanghai", "h")
	fmt.Printf("index: %v\n", index)

	index = strings.LastIndex("Hello Shanghai", "h")
	fmt.Printf("lastIndex: %v\n", index)

	newStr := strings.Replace("Hello Shanghai", "Shanghai", "Peking", 2)
	fmt.Println(newStr)

	var sentence string = "Kimi,Alonso,Leclerc"
	arr := strings.Split(sentence, ",")
	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Println(strings.ToUpper(sentence))
	fmt.Println(strings.ToLower(sentence))
}
