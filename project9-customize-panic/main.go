package main

import (
	"errors"
	"fmt"
)

// 我们还可以自定义错误：使用 errors.New 和 panic 内置函数

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from the panic")
		}
	}()

	err := sth(5)
	if err != nil {
		panic(err)
	}
}

func sth(param int) error {
	if param == 10 {
		fmt.Println("param is: ", param)
		return nil
	} else {
		fmt.Println("param must be 10")
		return errors.New("param must be 10")
	}
}
