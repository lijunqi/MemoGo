package main

import (
	"fmt"
	"my-pack/models"
	"my-pack/tempconv"
)

/*
 * Go 语言的代码通过包（package）组织，包类似于其它语言里的库（libraries）或者模块（modules）。
 * 一个包由位于单个目录下的一个或多个 .go 源代码文件组成，目录定义包的作用。
 * 每个源文件都以一条 package 声明语句开始
 *
 * main 包比较特殊。它定义了一个独立可执行的程序，而不是一个库。
 * 在 main 里的 main 函数也很特殊，它是整个程序执行时的入口
 *
 *
 * 每个包都对应一个独立的名字空间。
 * 例如，在image包中的Decode函数和在unicode/utf16包中的 Decode函数是不同的。
 * 要在外部引用该函数，必须显式使用image.Decode或utf16.Decode形式访问。
 *
 * 包还可以让我们通过控制哪些名字是外部可见的来隐藏内部实现信息。在Go语言中，一个简单的规则是：
 * 如果一个名字是大写字母开头的，那么该名字是导出的（汉字不区分大小写，因此汉字开头的名字是没有导出的）
 *
 * 如果一个名字是在函数内部定义，那么它就只在函数内部有效。
 * 如果是在函数外部定义，那么将在当前包的所有文件中都可以访问。
 */

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "Mcmillan",
	}
	fmt.Println(u)

	t := 32
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}
