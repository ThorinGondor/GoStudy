package main

import "fmt"

/**Go 中可以抛出一个 panic 的异常，然后在 defer 中通过 recover 捕获这个异常，然后正常处理
   因此panic 用来抛出异常，recover 用来恢复异常
   panic 能够改变程序的控制流，调用 panic 后会立刻停止执行当前函数的剩余代码，并在当前 Goroutine 中递归执行调用方的 defer；
   recover 可以中止 panic 造成的程序崩溃。它只能在 defer 中使用，在其他作用域无法调用；
**/

func main() {
	f()
	fmt.Println("end main")
}

// f方法中的 recover 捕获了 d 方法的 panic，打印了 panic 传递的参数，并且 main 方法是正常返回的。
// f 方法调用 d 抛出异常之后的代码没有执行! g 方法抛出 panic 之后的代码也没有执行!
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from panic")
		}
	}()
	fmt.Println("start f")
	d(0)
	fmt.Println("end f")
}

func d(param int) {
	fmt.Println("execute d")
	panic(param)
	fmt.Println("end d")
}
