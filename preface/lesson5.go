package main

import "fmt"

func main() {
	var x int8 = 16
	var y int8 = 32
	switch x {
	case y, 16: // case 可以带多个表达式
		fmt.Println("Yes")
	default:
		fmt.Println("Nope")
	}
}
