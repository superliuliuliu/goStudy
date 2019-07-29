package main

import (
	lib5 "awesomeProject/main/lib"
	"flag"
	"fmt"
)

/**
 * 测试命令源码文件接收参数
 */
var name string

func init() {
	// 使用flag.StringVar函数来接收参数 第一个参数指待是存储该命令参数的地址，第二个参数是为了指定该命令参数的名称 第三个参数指的是在未指定命令参数时的默认值
	// 第四个参数则是在使用该命令参数的简单说明
	flag.StringVar(&name, "name", "everyone", "the name you want to welcome")
}

func main() {
	flag.Parse()
	lib5.Test()
	fmt.Printf("hello, %s!\n", name)
}
