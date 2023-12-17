package main

import "fmt"

/*
 * 本章主要讲如何使用 接口
 * 实际上，Go 的多态就是通过接口实现的，和 Java 不同，Java 通过父类的通过对多个子类向下转型实现多态
 */

// interface 类型可以定义一组方法，但是这些方法在interface内不能实现。并且 interface 不能包含任何变量。只有某个自定义类型(比如结构体 Mouse)要使用的时候,再根据具体情况实现接口提供的这些方法。
// 声明一个 usb 设备接口
type useDevice interface {
	start()
	running()
}

// 声明鼠标，并应当实现 usb 设备的全部方法！
type mouse struct {
}

func (m *mouse) start() {
	fmt.Println("mouse starting ...")
}

func (m *mouse) running() {
	fmt.Println("mouse running")
}

// 声明键盘，并应实现 usb 设备的全部方法！
type keyboard struct {
}

func (k *keyboard) start() {
	fmt.Println("keyboard starting ...")
}

func (k *keyboard) running() {
	fmt.Println("keyboard running")
}

// PC 声明一台电脑，并使用这些 usbDevice的方法
type PC struct {
}

// Run 该函数的传参是一个接口！因此任何实现该接口的实现类都可以作为参数！
func (p *PC) Run(device useDevice) {
	device.start()
	device.running()
}

// 注：空接口 interface{} 没有任何方法，所以所有类型都实现了空接口, 即我们可以把任何一个变量赋给空接口

// T 定义一个空接口
type T interface {
}

func main() {
	var myPC = &PC{}

	var newMouse = &mouse{}
	var newKeyboard = &keyboard{}

	myPC.Run(newKeyboard)
	myPC.Run(newMouse)

	// 任何类型的变量都可以赋值给空接口类型的变量
	var oneMouse T = &mouse{}
	fmt.Println(oneMouse)
}

// 说明：接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低偶合的思想。
// Golang 中的接口，不需要显式的实现。只要一个结构体，含有接口类型中的所有方法，那么这个结构体就实现这个接口。因此，Golang 中没有 implement 这样的关键字。
// 一个自定义结构体可以实现多个接口
// 一个接口(比如 A 接口)可以继承多个别的接口(比如 B,C 接口)，如下所示，这时如果要实现 A 接口，则也必须将 B,C 接口的方法也全部实现。接口继承和结构体继承一样，都是通过匿名接口实现。

type C interface {
	CMethod()
}

type B interface {
	BMethod()
}

type A interface {
	B
	C
	AMethod()
}

// 接口和继承解决的解决的问题不同
// 继承的价值主要在于：解决代码的复用性和可维护性。
// 接口的价值主要在于：功能的设计，设计好各种规范(方法)，让其它自定义类型去实现这些方法。

// 在 Go 语言，oop的多态特征是通过接口实现的。可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态。
// 比如 Usb 接口，Usb usb ，既可以接收手机变量，又可以接收相机变量，就体现了 Usb 接口 多态特性
