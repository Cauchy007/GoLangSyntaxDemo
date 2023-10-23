package basic_syntax

import "fmt"

// 1. 结构体指针
func structPointerExample() {
	type Point struct {
		X, Y int
	}
	p := Point{1, 2}
	ptr := &p // 结构体指针

	fmt.Printf("结构体 p 的X值：%d, Y值：%d\n", ptr.X, ptr.Y)
}

// 2. 切片和指针
func sliceAndPointerExample() {
	data := []int{1, 2, 3, 4, 5}
	ptr := &data[0] // 指向切片的第一个元素的指针

	fmt.Printf("切片的第一个元素：%d\n", *ptr)
}

// 3. 函数中的指针参数
func modifyValueWithPointer(x *int) {
	*x = *x * 2
}

func PointExample() {
	// 1. 结构体指针
	structPointerExample()

	// 2. 切片和指针
	sliceAndPointerExample()

	// 3. 函数中的指针参数
	value := 5
	modifyValueWithPointer(&value)
	fmt.Println("通过指针修改后的值:", value)
}
