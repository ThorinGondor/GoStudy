package pojo

// 该结构体 student 开头为小写，不能在包外构建对象，那么可以通过工厂模式来解决问题
type student struct {
	id   int32
	name string
}

// StudentFactory 工厂模式：返回值类型
func StudentFactory(id int32, name string) student {
	return student{
		id:   id,
		name: name,
	}
}

// StudentAddressFactory 工厂模式：返回引用类型
func StudentAddressFactory(id int32, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}
