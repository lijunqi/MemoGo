package main

import (
	"fmt"
	"time"
)

func for_defer() {
	fmt.Println("=== Defer in for loop ===")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func block_defer() {
	fmt.Println("=== Defer in block ===")
	{
		defer fmt.Println("3")
		fmt.Println("1")
	}

	fmt.Println("2")
}

func pass_by_value1() {
	/*
	 * 调用 defer 关键字会立刻拷贝函数中引用的外部参数,
	 * 所以 time.Since(startedAt) 的结果不是在 main 函数退出之前计算的,
	 * 而是在 defer 关键字调用时计算的,最终导致上述代码输出 0s
	 */
	fmt.Println("=== Pass by value 1 ===")
	startAt := time.Now()
	defer fmt.Println(time.Since(startAt))

	time.Sleep(2 * time.Second)
}

func pass_by_value2() {
	/*
	 * 虽然调用 defer 关键字时也使用值传递, 但是因为拷贝的是函数指针,
	 * 所以 time.Since(startedAt) 会在 main 函数返回前调用并打印出符合预期的结果
	 */
	fmt.Println("=== Pass by value 2 ===")
	startAt := time.Now()
	defer func() {
		fmt.Println(time.Since(startAt))
	}()

	time.Sleep(2 * time.Second)
}

func retVal() int {
	/*
	 * defers run in the reverse order that you define, it means that the last written is the first executed.
	 * 多个defer的执行顺序为"后进先出".
	 * defer, return, 返回值, 三者的执行逻辑应该是:
	 * return最先执行, return负责将结果写入返回值中, 接着defer开始执行一些收尾工作, 最后函数携带返回值退出.
	 */
	fmt.Println("=== retVal ===")
	a := 1
	defer func() {
		fmt.Printf("a1 = %d\n", a)
		a = 4
		fmt.Printf("a2 = %d\n", a)
	}()
	return a
}

func main() {
	for_defer()
	block_defer()
	pass_by_value1()
	pass_by_value2()
	fmt.Printf("ret a = %d\n", retVal())
}
