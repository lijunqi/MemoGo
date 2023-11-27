package main

import (
	"fmt"
	"my-scope/packageone"
)

/*
 * 任何在函数外部（也就是包级语法域）声明的名字可以在同一个包的任何源文件中访问的。
 *
 * 对于导入的包，例如tempconv导入的fmt包，则是对应源文件级的作用域，
 * 因此只能在当前的文件中访问导入的fmt包，当前包的其它源文件无法访问在当前源文件导入的包
 */
func f() int {
	return 123
}

func g(x int) int {
	return x + 1
}

var myVar = "This is a package level variable"

func main() {
	// 和for循环类似，if和switch语句也会在条件部分创建隐式词法域，
	// 还有它们对应的执行体词法域
	if x := f(); x == 0 {
		fmt.Println("Branch 1:")
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println("Branch 2:")
		fmt.Println(x, y)
	} else {
		fmt.Println("Branch 3:")
		fmt.Println(x, y)
	}

	// variables
	var blockVar = "This is the block level variable"
	packageone.PrintMe(myVar, blockVar)
}
