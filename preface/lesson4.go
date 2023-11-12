package main

import "fmt"

func main() {
	var number int32 = 8
	fmt.Println(number >> 3) // 8 ÷ (2^3)
	fmt.Println(number << 2) // 8 x (2^2)
}
