package main

import (
	"fmt"
)

// panic()是一个内置的Go函数，它终止Go程序的当前流程并开始panicking！
// 另一方面，recover()函数也是一个内置的Go函数，允许你收回那些使用了panic()函数的goroutine的控制权
func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!") // NOT execute
	fmt.Println("Exiting a()") // NOT execute
}

// 另请注意，函数b对函数a一无所知。 但是，函数a包含处理b函数的panic情况的Go代码
func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	fmt.Println("Exiting b()") // NOT execute
}

func main() {
	a()
	fmt.Println("main() ended!")
}
