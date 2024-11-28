package main

import (
	"fmt"
	"slices"
)

/*
 * 切片的底层是数组，这意味着Go为每一个切片创建一个底层数组
 *
 * 切片作为函数的形参时是传引用操作，传递的是指向切片的内存地址，
 * 这意味着在函数中对切片的任何操作都会在函数结束后体现出来。
 *
 * 另外，函数中传递切片要比传递同样元素数量的数组高效，
 * 因为Go只是传递指向切片的内存地址，而非拷贝整个切片。
 */
func main() {
	/* 1.
	 * 与定义数组相比，切片字面量只是没有指定元素数量。
	 * 如果你在[]中填入数字，你将得到的是数组。
	 * aSlice := []int{1, 2, 3}
	 */

	/* 2.
	 * 也可以使用make()创建一个空切片，并指定切片的长度和容量。
	 * 容量这个参数可以省略，在这种情况下容量等于长度
	 *
	 * Go自动将空切片的元素初始化为对应元素的零值，
	 * 意味着切片初始化时的值是由切片类型决定的
	 */
	integer := make([]int, 5)
	for idx, item := range integer {
		fmt.Printf("slice[%d] = %d\n", idx, item)
	}
	fmt.Println()

	// 追加元素到切片，此操作将触发切片自动扩容
	integer = append(integer, -5900)
	for idx, item := range integer {
		fmt.Printf("slice[%d] = %d\n", idx, item)
	}

	// * Contains
	strList := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slices.Contains(strList, "c"))
	fmt.Println(slices.Contains(strList, "z"))
}
