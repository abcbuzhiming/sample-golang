package main

import (
	// 需在此处添加代码。[1]
	"flag"
	"fmt"
)

var name string

func init() {
	// 需在此处添加代码。[2]
	//第 1 个参数是用于存储该命令参数值的地址;第 2 个参数是为了指定该命令参数的名称;第 3 个参数是为了指定在未追加该命令参数时的默认值，这里是everyone;第 4 个函数参数，即是该命令参数的简短说明了
	flag.StringVar(&name, "name", "everyone", "The greeting object.")	
}

func main() {
	// 需在此处添加代码。[3]
	flag.Parse()		//用于真正解析命令参数，并把它们的值赋给相应的变量
	fmt.Printf("Hello, %s!\n", name)
}


