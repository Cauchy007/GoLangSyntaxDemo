package basic_syntax

import "fmt"

// 1. 定义一个接口
type Shape interface {
	Area() float64
}

// 2. 创建两个类型，实现接口
type Circle2 struct {
	Radius float64
}

func (c Circle2) Area() float64 {
	return 3.14159265359 * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 3. 函数使用接口类型作为参数
func CalculateArea(shape Shape) {
	area := shape.Area()
	fmt.Printf("面积：%f\n", area)
}

func InterfaceExample() {
	// 1. 创建两个类型，实现接口
	circle := Circle2{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 3}

	// 2. 函数使用接口类型作为参数
	CalculateArea(circle)
	CalculateArea(rectangle)
}
