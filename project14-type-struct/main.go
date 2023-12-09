package main

import (
	"awesomeProject/project14-type-struct/pojo"
	"fmt"
	"time"
)

type A struct {
	id   int64 // 字段名称若小写，则该字段不允许被包外访问，
	name string
}

type B struct {
	id   int64
	name string
}

/**
 * 本章主要讲 结构体：type XXX struct {}
 * 结构体是用户自定义的一种类型
 * 看下如何：声明一个结构体，多种方法创建一个对象，以及两个不同结构体的对象强制转换
 */
func main() {
	// 工厂模式新建一个对象
	var employee pojo.Employee = pojo.Employee{Id: 398427427, Name: "Kimi"}
	(&employee).SetSex("MAN") // 方法的应用：使用引用传递，这样才能在方法里修改原对象的属性！
	fmt.Println(employee.GetSex())
	fmt.Println(employee)

	var employee2 = employee
	employee2.District = "Finland"
	// 顺便看下 go 如何创建一个时间：
	employee2.EnrollDate = time.Date(2001, time.Month(3), 12, 15, 30, 20, 0, time.Now().Location())
	fmt.Println(employee2)

	// 以指针的方式创建对象
	var employee3 *pojo.Employee = new(pojo.Employee)
	(*employee3).Id = 458937524 // 实际上去掉 * 也可以赋值！因为即使你不写 *，go在底层也会帮你加上 *，go允许我们这样写是为了提供方便
	(*employee3).Name = "Alonso"
	employee3.District = "Spain" // 去掉 * 之后正常赋值
	fmt.Println(*employee3)

	// 不同类型结构体，如果字段类型、名称、数量都完全相同，则可以强制互转：
	var a A = A{id: 1, name: "Durant"}
	a.name = "Curry"
	var b B = B(a)
	fmt.Println(b)

	// go 的结构体没有构造函数！那么如何构造对象呢？可以通过工厂模式，请看 project 15 factory mode
}
