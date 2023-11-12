package main

import "fmt"

func main() {
	var x int8 = 16
	var y int8 = 32
	switch x {
	case y, 16: // case 可以带多个表达式
		fmt.Println("Yes")
	default: // default 代码块非必须
		fmt.Println("Nope")
	}

	switch x {
	case 16:
		fmt.Println("case 1")
		fallthrough // attention: 可以向下面的case穿透，即使不符合 case 里的表达式判断结果，默认只能穿透一层case
	case 32:
		fmt.Println("case 2")
		fallthrough
	case 64:
		fmt.Println("case 3")
	case 108:
		fmt.Println("case 3")
	}

	var itf interface{}
	switch variable := itf.(type) {
	case int8:
		fmt.Printf("type int8: %T", variable)
	case nil:
		fmt.Printf("type nil %T", variable)
	}
}
