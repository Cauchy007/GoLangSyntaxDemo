package basic_syntax

import (
	"fmt"
	"runtime"
)

func getCallerFunctionName() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return
	}

	fmt.Println(fn.Name())
}

var globalVariable string = "global variable"

// globalVariable2 := "global variable2"  error: non-declaration statement outside function body
func ExampleVariable() {
	getCallerFunctionName()
	// 1. 全局变量
	fmt.Print("打印全局变量：", globalVariable)
	//fmt.Print(globalVariable2)

	// 2. 局部变量申明和赋值
	localVariable1 := "local variable1"
	var localVariable2 = "local variable2"
	fmt.Print("打印局部变量：", localVariable1)
	fmt.Print("打印局部变量：", localVariable2)

	{
		//3. 变量作用域，此处的localVariable2与上面的localVariable2不是同一个变量
		localVariable2 := 4
		fmt.Print(localVariable2)
	}

	// 4. 常量申明
	const constant1 = "constant1"
	//constant1 = "constant1" // error: cannot assign to constant
	const constant2 string = "constant2"
	const constant3, constant4 = 3, 4
	fmt.Print(constant1)
	fmt.Print(constant2)
	fmt.Print(constant3)
	fmt.Print(constant4)
}

func ExampleExpression() {
	getCallerFunctionName()
	// 1. if表达式
	ExampleIf := func(x int) {
		getCallerFunctionName()
		if x > 0 {
			fmt.Println("1 > 0")
		} else if x < 0 {
			fmt.Println("1 < 0")
		} else {
			fmt.Println("1 == 0")
		}
	}
	ExampleIf(100)

	// 2. switch表达式
	ExampleSwitch := func(x int) {
		getCallerFunctionName()
		switch x {
		case 1:
			fmt.Println("x == 1")
		case 2:
			fmt.Println("x == 2")
		default:
			fmt.Println("x == default")
		}
	}
	ExampleSwitch(1)

	// 3. for表达式
	ExampleFor := func() {
		getCallerFunctionName()
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}

		i := 0
		for i < 5 {
			fmt.Println(i)
			i++
		}

		for {
			fmt.Println("无限循环")
			i--
			if i < 0 {
				break
			}
		}

		x := []int{1, 2, 3}
		for i, n := range x {
			fmt.Println(i, n)
		}
	}
	ExampleFor()
}

func ExampleFunction() {
	getCallerFunctionName()

	// 1. 函数声明
	ExampleFunction := func(x int, y int) int {
		return x + y
	}
	fmt.Println(ExampleFunction(1, 2))

	// 2. 函数闭包
	ExampleClosure := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}
	closure := ExampleClosure()
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
}

func ExampleDataStructure() {
	getCallerFunctionName()
	// 1. 数组
	ExampleArray := func() {
		getCallerFunctionName()
		var array1 [5]int
		array2 := [5]int{1, 2, 3, 4, 5}
		array3 := [...]int{1, 2, 3, 4, 5}
		array4 := [5]int{1: 10, 3: 30}
		fmt.Println(array1)
		fmt.Println(array2)
		fmt.Println(array3)
		fmt.Println(array4)
	}
	ExampleArray()

	// 2. 切片
	ExampleSlice := func() {
		getCallerFunctionName()
		var slice1 []int
		slice2 := []int{1, 2, 3, 4, 5}
		slice3 := make([]int, 5)
		slice4 := make([]int, 5, 10)
		fmt.Println(slice1)
		fmt.Println(slice2)
		fmt.Println(slice3)
		fmt.Println(slice4)
	}
	ExampleSlice()

	// 3. Map
	ExampleMap := func() {
		getCallerFunctionName()
		var map1 map[string]int
		map2 := map[string]int{"a": 1, "b": 2}
		map3 := make(map[string]int)
		fmt.Println(map1)
		fmt.Println(map2)
		fmt.Println(map3)
	}
	ExampleMap()

	// 4. 结构体
	ExampleStruct := func() {
		getCallerFunctionName()
		type Point struct {
			x int
			y int
		}
		point1 := Point{1, 2}
		point2 := Point{x: 1, y: 2}
		point3 := Point{x: 1}
		fmt.Println(point1)
		fmt.Println(point2)
		fmt.Println(point3)
	}
	ExampleStruct()
}

func ExamplePointer() {
	getCallerFunctionName()
	// 1. 指针
	ExamplePointer := func() {
		getCallerFunctionName()
		var x int = 1
		var y *int = &x
		fmt.Println(x)
		fmt.Println(y)

		*y = 2
		fmt.Println(*y)
	}
	ExamplePointer()
}

type ShapeArea interface {
	ShapeArea() float64
}
type Name interface {
	Name() string
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) ShapeArea() float64 {
	return r.width * r.height
}

type Circle struct {
	Radius float64
}

func (c Circle) ShapeArea() float64 {
	return 3.14159265359 * c.Radius * c.Radius
}
func (c Circle) Name() string {
	return "Circle"
}

func ExampleInterface() {
	getCallerFunctionName()
	// 1. 接口
	ExampleInterface := func() {
		var shape ShapeArea = Rect{2, 3}
		fmt.Println(shape.ShapeArea())
	}
	ExampleInterface()

	// 2. 接口类型断言
	ExampleInterfaceAssert := func() {
		var shape ShapeArea = Rect{2, 3}
		rect := shape.(Rect)
		fmt.Println(rect.width)
	}
	ExampleInterfaceAssert()

	// 3. 空接口
	ExampleEmptyInterface := func() {
		var x interface{}
		x = 1
		x = "hello"
		fmt.Println(x)
	}
	ExampleEmptyInterface()

	// 4. 类型选择
	ExampleTypeSwitch := func() {
		var x interface{}
		x = 1
		switch t := x.(type) {
		case int:
			fmt.Println("整型", t)
		case string:
			fmt.Println("字符串", t)
		}
	}
	ExampleTypeSwitch()

	// 5. 接口的嵌套
	ExampleInterfaceEmbed := func() {
		var shape ShapeArea = Circle{5}
		var name Name = shape.(Name)
		fmt.Println(name.Name())
	}

	ExampleInterfaceEmbed()
}
