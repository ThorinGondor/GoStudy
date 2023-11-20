package user

import "fmt"

func PrintUsername(param string) {
	fmt.Printf("User print: %v \n", param)
	// 同一个包下的函数可以被直接调用，无需 import
	PrintAge()
}
