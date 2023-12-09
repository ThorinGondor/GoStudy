package main

import "fmt"

/**
 * 本章主要讲如何实现继承
 * 可通过嵌套匿名结构体 实现继承
 * 和 java 不同，java 的继承下，父类字段和子类字段平级，但是 Go 的继承，父类字段嵌套在内部：{{Nvidia Graphic Vision Card 4999.5 } 12 230}
 * java 只支持单继承，而 Go 支持多继承（不推荐）
 */

type Goods struct {
	name           string
	price          float32
	originalRegion string
}

func (g *Goods) getName() string {
	return g.name
}

func (g *Goods) getPrice() float32 {
	return g.price
}

type visionCard struct {
	Goods                // 嵌套匿名结构体，可直接访问它的字段与方法！
	manufacturingProcess int8
	powerDissipation     int32
}

func main() {
	var myVisionCard = visionCard{
		manufacturingProcess: 12,
		powerDissipation:     230,
	}
	myVisionCard.Goods.name = "Nvidia Graphic Vision Card"
	myVisionCard.Goods.price = 4999.5

	fmt.Println(myVisionCard.getName())
	fmt.Println(myVisionCard.getPrice())
	fmt.Println(myVisionCard)
}
