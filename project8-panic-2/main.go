package main

import "fmt"

// 使用 panic 和 recover 可以提高代码的健壮性！也可以进行发生异常后续的完善处理！

func main() {
	fmt.Println("start main")
	f()
	fmt.Println("end main")
}

func f() {
	fmt.Println("start f")
	g() // 此函数内部可能出现异常
	fmt.Println("end f")
}

// 为了让别的函数执行该函数时保证能安全返回并继续，需要使用 defer + recover 函数捕获异常
func g() {
	defer func() {
		if err := recover(); err != nil {
			// 此处捕获 panic 并恢复
			fmt.Println("err happened: ", err)
			// 此处完成后续处理
			fmt.Println("Send mail ====> roderick.li@ottoint.com")
		}
	}()
	fmt.Println("start g")
	var num1, num2 = 1, 0
	num3 := num1 / num2        // 此处触发 panic
	fmt.Println("end g", num3) // 若上一段抛出异常，则该段代码不会被执行
}
