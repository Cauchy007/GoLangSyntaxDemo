package main

import basic_syntax "github.com/Cauchy007/GoLangSyntaxDemo/BasicSyntax"

var name string

func init() {
	name = "cauchy"
}

// 调用BasicSyntax package中的FunctionExample函数
func BasicSyntaxExample() {
	basic_syntax.ExampleVariable()
	basic_syntax.ExampleExpression()
	basic_syntax.ExampleDataStructure()
	basic_syntax.ExamplePointer()
	basic_syntax.ExampleInterface()
	basic_syntax.ExampleFunction()
}

func main() {
	BasicSyntaxExample()
}
