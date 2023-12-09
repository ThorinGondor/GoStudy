package pojo

import "time"

type Employee struct {
	Id         int64 // 若需允许被包外访问，字段名称必须大写开头，形成 field export，变成公开字段
	Name       string
	District   string
	sex        string // 小写开头字段，是私有，只有本包能访问，或通过方法对外提供修改方法（实现封装）
	EnrollDate time.Time
	IsActive   bool
	Info       map[string]string
}

/*
 * 方法，和函数不同，它和对象绑定，只能通过 对象.方法 的形式调用，而不能像函数那样直接被调用！
 * 这点和 java 类似，java 的类也有自己的方法，通过 对象.方法 可以调用类的方法
 */

// GetSex 如下这样声明，就是一个值传递
func (employee Employee) GetSex() string {
	return employee.sex
}

// SetSex 如下这样声明，就是一个引用传递
func (employee *Employee) SetSex(sex string) {
	employee.sex = sex
}
