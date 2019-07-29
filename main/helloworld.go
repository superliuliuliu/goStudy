package main

import (
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

// 定义全局变量 一般采用这种因式分解的形式 除此之外全局变量允许声明但不使用 局部变量一旦声明必须使用
var (
	a int
	b bool
	c float64
)

func Hello(str string) {
	// 在函数体内快速定义变量一般采用:=快速定义的方式
	l := "liu"
	g := "gao"
	y := "yang"
	fmt.Println("测试函数体内快速定义变量")
	fmt.Println(l + g + y)
	// 定义多个变量
	fmt.Println("测试定义多个变量")
	var test1, test2, test3 int = 1, 2, 3
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	// 定义变量 若不赋值则默认取零值

	var test4 int
	var test5 bool
	var test6 string
	fmt.Println("测试定义变量但不执行赋值操作")
	fmt.Println(test4)
	fmt.Println(test5)
	// 不打印内容
	fmt.Println(test6)
	// 不同类型的变量可以在同一行内被赋值
	test1, test2, test5 = 6, 6, true

	// GO语言常量 常量的值在运行时不会被改变 常量中的数据类型只可能是布尔型、数字型和字符串型
	const a string = "a"
	const b = "a"
	const c, d = "abc", true
	e, f := "abcd", "abcd"
	fmt.Println(a + b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e + f)
	// 常量作为枚举使用
	const (
		MAN   = 0
		WOMAN = 1
		CHILD = 3
	)
	// 特殊常量 iota 一个可以被编译器修改的常量
	// iota在const关键字出现时被重置为0,const中每新增一行常量声明将对iota计数一次
	const (
		// 初始值为0
		num = iota
		// iota + 1
		num1
		// +1
		num2
		// 赋值
		num3 = "haha"
		// 赋值
		num4
		// num5 = 100
		num5 = 100
		// num6 = 100
		num6
		// 恢复计数 numb7 = 7
		num7 = iota
		// num8 = 8
		num8
	)
	fmt.Println(num, num1, num2, num3, num4, num5, num6, num7, num8)
	fmt.Println(num)

	fmt.Println("hello: ", str)
	// 算数运算符
}

// 运算符测试类
func testOperator() {
	// 算数运算符
	a := 1
	b := 2
	var c int

	c = a + b
	fmt.Println("加法测试: %d", c)
	c = b - a
	fmt.Println("加法测试: %d", c)

	// 关系运算符

	// 逻辑运算符
}
