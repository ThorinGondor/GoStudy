package main

import "fmt"

func main() {
	// 数组和切片很类似，但也不同，数组的使用场景相对有限，切片才更加常用。
	// 切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
	// 当切片作为函数参数时，和数组是不同的，如果一个函数接受一个切片参数，它对切片元素所做的更改将对调用者可见，类似于将指针传递给了底层数组。
	// 所以数组是值传递，相当于拷贝一份，函数对它的修改不影响原数组，而切片是引用传递，函数对它的修改影响原切片！
	// 当然，在方法的使用上，遍历、访问、len 等他们都相同
	var slice = make([]int, 5, 100)
	fmt.Println("len", len(slice), "cap: ", cap(slice))

	slice[3] = 70 // 查找并赋值元素
	fmt.Println(slice)

	// append 方法：切片 append 操作的本质就是对数组扩容
	// go 底层会创建一下新的数组 newArr(安装扩容后大小)
	slice = append(slice, 100, 45, 71) // 向后追加元素
	fmt.Println(slice)

	slice = append(slice, slice...) // 向后追加某切片的所有元素
	fmt.Println(slice)

	// copy 方法，dst <- src，若目标 dst 切片容量不够，则只拷贝 dst 切片的容量允许的数量
	// 注意该方法的后一个参数才是拷贝源！
	var targetSlice = make([]int, 5)
	copy(targetSlice, slice) // targetSlice容量只有5，那就只拷贝 slice 的前5个
	fmt.Println(targetSlice)

	// 注意：string 其实也是 slice 的一种，因此它也可以进行切片处理
	var str string = "ABC123456789"
	var strSlice = str[2:] // 截取下标为2及其以后的字符
	fmt.Println(strSlice)

	// 但也有区别 string 不可变，因此不能使用形如 str[3] = 'X' 这种方式去赋值修改字符串
}
