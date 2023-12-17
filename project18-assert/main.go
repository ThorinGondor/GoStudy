package main

import (
	"fmt"
	"time"
)

/*
 * 本章主要讲了如何使用断言判断数据类型或自定义结构体
 * 但 Go 的断言不能直接使用，需要使一个空接口通过多态指定你的变量，然后再判断该变量的数据类型
 */
func main() {
	var t interface{}
	var x float32 = 1.11
	t = x // 所有的结构体都实现了空接口，所以可以将空接口指向我们要判断的变量
	// 判断该接口的数据类型
	var y = t.(float32)
	fmt.Printf("%T %v \n", y, y)

	var slice = make([]interface{}, 10)
	slice[0] = 1.15
	slice[1] = 19304
	slice[2] = false
	slice[3] = time.Now()
	slice[4] = "test me"
	slice[5] = 1.5487234

	fmt.Println(slice)

	judgeDataType(slice[0], slice[1], slice[2], slice[3], slice[4])
}

func judgeDataType(items ...interface{}) {
	for _, item := range items {
		fmt.Printf("infer item: %v\n", item)
		switch item.(type) {
		case float32:
			fmt.Println("float32")
		case float64:
			fmt.Println("float64")
		case int32:
			fmt.Println("int32")
		case int64:
			fmt.Println("int64")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		case time.Time:
			fmt.Println("Time")
		default:
			fmt.Println("cannot infer type")
		}
	}

}
