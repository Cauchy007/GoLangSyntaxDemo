package basic_syntax

import "fmt"

// 1. 条件语句 - if
func ifStatementExample() {
	age := 25

	if age < 18 {
		fmt.Println("未成年")
	} else if age >= 18 && age < 60 {
		fmt.Println("成年")
	} else {
		fmt.Println("老年")
	}
}

// 2. 循环 - for
func forLoopExample() {
	// 基本的 for 循环
	for i := 0; i < 5; i++ {
		fmt.Println("基本 for 循环迭代：", i)
	}

	// 使用 for 进行切片迭代
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("切片迭代，索引 %d，值 %d\n", index, value)
	}
}

// 3. 条件循环 - for
func conditionalLoopExample() {
	j := 0
	for j < 5 {
		fmt.Println("条件循环迭代：", j)
		j++
	}
}

// 4. switch 语句
func switchStatementExample() {
	day := "周一"

	switch day {
	case "周一":
		fmt.Println("今天是周一")
	case "周二":
		fmt.Println("今天是周二")
	default:
		fmt.Println("其他日期")
	}
}

func ControlExample() {
	ifStatementExample()
	forLoopExample()
	conditionalLoopExample()
	switchStatementExample()
}
