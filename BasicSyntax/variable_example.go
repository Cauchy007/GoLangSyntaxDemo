package basic_syntax

import "fmt"

// 1. 整数类型
func integerExample() {
	var a int
	a = 10
	b := 5
	fmt.Printf("整数 a: %d, b: %d\n", a, b)
}

// 2. 浮点数类型
func floatExample() {
	var x float64
	x = 3.14
	y := 2.718
	fmt.Printf("浮点数 x: %f, y: %f\n", x, y)
}

// 3. 字符串类型
func stringExample() {
	var str string
	str = "Hello, Go!"
	anotherStr := "你好，世界！"
	fmt.Printf("字符串 str: %s, anotherStr: %s\n", str, anotherStr)
}

// 4. 布尔类型
func booleanExample() {
	var isTrue bool
	isTrue = true
	isFalse := false
	fmt.Printf("布尔值 isTrue: %t, isFalse: %t\n", isTrue, isFalse)
}

// 5. 复合数据类型 - 数组
func arrayExample() {
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	moreNumbers := [3]int{3, 4, 5}
	fmt.Printf("整数数组 numbers: %v, moreNumbers: %v\n", numbers, moreNumbers)
}

// 6. 复合数据类型 - 切片
func sliceExample() {
	var slice []int
	slice = []int{1, 2, 3}
	anotherSlice := make([]string, 3)
	fmt.Printf("整数切片 slice: %v, 字符串切片 anotherSlice: %v\n", slice, anotherSlice)
}

// 7. 复合数据类型 - 映射
func mapExample() {
	var person map[string]string
	person = make(map[string]string)
	person["name"] = "Alice"
	person["age"] = "30"
	anotherMap := map[string]int{"a": 1, "b": 2}
	fmt.Printf("映射 person: %v, anotherMap: %v\n", person, anotherMap)
}

// 8. 复合数据类型 - 结构体
func structExample() {
	type Point struct {
		X, Y int
	}
	p := Point{1, 2}
	fmt.Printf("结构体 p: %+v\n", p)
}

func VariableExample() {
	integerExample()
	floatExample()
	stringExample()
	booleanExample()
	arrayExample()
	sliceExample()
	mapExample()
	structExample()
}
