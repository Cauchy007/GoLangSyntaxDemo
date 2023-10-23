package basic_syntax

import "fmt"

// 1. 基本函数
func add(a, b int) int {
	return a + b
}

// 2. 函数参数
func greet(name string) {
	fmt.Printf("你好，%s!\n", name)
}

// 3. 函数返回值
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	quotient := dividend / divisor
	return quotient, nil
}

// 4. 匿名函数
func operateOnNumbers(x, y int, operation func(int, int) int) int {
	result := operation(x, y)
	return result
}

// 5. 可变参数函数
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func FunctionExample() {
	// 1. 基本函数
	result := add(3, 5)
	fmt.Println("3 + 5 =", result)

	// 2. 函数参数
	greet("Alice")

	// 3. 函数返回值
	quotient, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("10.0 / 2.0 =", quotient)
	}

	// 4. 匿名函数
	double := func(a, b int) int {
		return 2 * (a + b)
	}
	result1 := operateOnNumbers(4, 3, double)
	fmt.Println("双倍结果:", result1)

	// 5. 可变参数函数
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("总和:", total)
}
