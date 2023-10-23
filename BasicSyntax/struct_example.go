package basic_syntax

import "fmt"

// 1. 结构体的定义
type Point struct {
	X, Y int
}

// 2. 为结构体定义方法
func (p Point) DistanceToOrigin() float64 {
	return float64(p.X*p.X + p.Y*p.Y)
}

// 3. 为结构体定义带指针接收者的方法
func (p *Point) Scale(factor int) {
	p.X *= factor
	p.Y *= factor
}

// 4. 结构体嵌套
type Circle struct {
	Center Point
	Radius float64
}

// 5. 为结构体定义方法
func (c Circle) Area() float64 {
	return 3.14159265359 * c.Radius * c.Radius
}

func StructExample() {
	// 1. 结构体的定义
	p1 := Point{3, 4}
	p2 := Point{X: 5, Y: 12}

	// 2. 为结构体定义方法
	distance1 := p1.DistanceToOrigin()
	distance2 := p2.DistanceToOrigin()
	fmt.Printf("p1到原点的距离：%f, p2到原点的距离：%f\n", distance1, distance2)

	// 3. 为结构体定义带指针接收者的方法
	p1.Scale(2)
	fmt.Printf("p1的坐标：%v\n", p1)

	// 4. 结构体嵌套
	c := Circle{Point{0, 0}, 5.0}

	// 5. 为结构体定义方法
	area := c.Area()
	fmt.Printf("圆的面积：%f\n", area)
}
