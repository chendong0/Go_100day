package main

import (
	"fmt"
)
//常量定义格式: const 标识符 类型 = 值
const c1 string = "2018年4月23日" //显式类型定义
const c2 = "2019年4月23日" //隐式类型定义
//和变量一样,常量也可以不使用数据类型直接赋值,其类型与值一致
const c3, c4 = "20200423", "20210423"//常量声明简写
	/*
	常量:
	1.概念:同变量类似,程序执行过程中数值不能改变
	2.语法:显示型定义,隐式类型定义
	3.常数:固定的值:100, "abc"
	*/
const (
	c5 = "20220423"
	c6 = "20230413"
)
//推荐的常量使用方式
//修改常量数值会报错
//4.一族常量中,如果某个常量没有初始值,默认和上一行一致
const (
	c7 int = 100
	c8
	c9 string ="ruby"
	c10
	c11
)

func main()  {
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println(c5)
	fmt.Println(c7)
	fmt.Println(c8)
	fmt.Println(c9)
	fmt.Println(c10)
	fmt.Println(c11)
	enums() //要定义这个枚举函数,才显示出iota

	const LENGTH int = 10
	const WIDTH int =5
	var area int
	const a, b, c = 1, false, "str" //多重赋值
	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
}
func enums() {
	const (
		cpp = iota
		java
		python
		javascript

		unknown = 0
		female = 1
		male = 2
	)
//常量枚举方式
//6.枚举类型:使用常量组作为枚举类型.一组相关数值的数据
	fmt.Println(cpp,java, python, javascript)
	fmt.Println()
	fmt.Println(unknown, female, male)
}