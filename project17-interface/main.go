package main

import "fmt"

/*
 * 本章主要讲如何使用 接口
 * 实际上，Go 的多态就是通过接口实现的，和 Java 不同，Java 通过类的向下转型实现多态
 */

// interface 类型可以定义一组方法，但是这些方法在interface内不能实现。并且 interface 不能包含任何变量。只有某个自定义类型(比如结构体 Mouse)要使用的时候,再根据具体情况实现接口提供的这些方法。
// 声明一个 usb 设备接口
type useDevice interface {
	start()
	running()
}

// 声明鼠标，并实现 usb 设备的方法
type mouse struct {
}

func (m *mouse) start() {
	fmt.Println("mouse starting ...")
}

func (m *mouse) running() {
	fmt.Println("mouse running")
}

// 声明键盘，并实现 usb 设备的方法
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

// 空接口 interface{} 没有任何方法，所以所有类型都实现了空接口, 即我们可以把任何一个变量赋给空接口

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
