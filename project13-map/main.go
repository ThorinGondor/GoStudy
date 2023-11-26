package main

import "fmt"

/*
*

	go 的 map
*/
func main() {
	// map 在使用前一定要 make
	// map 的 key 唯一，若出现重复，则以最后这个 key-value 为准
	// map 的 value 可相同
	// map 是引用类型，因此函数传参后如果出现修改，会影响原 map 的元素！
	var params = make(map[string]string)
	params["id"] = "100"
	params["name"] = "Kimi"
	params["district"] = "Finland"
	fmt.Println(params)

	// map 套娃练习
	var result = make(map[string]map[string]string)
	result["01"] = params
	result["02"] = params
	fmt.Println(result)

	// 遍历 map
	for k, v := range result {
		fmt.Printf("layer = %v\n", k)
		for k1, v1 := range v {
			fmt.Printf("\tkey = %v, val = %v\n", k1, v1)
		}
	}

	// map 操作元素
	delete(params, "district")
	fmt.Println(params)
	fmt.Println(result)

}
