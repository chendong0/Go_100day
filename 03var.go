package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := 11
	/*
	变量:variable
	概念:一小块内存,用于储存数据,在程序运行过程中数值可以改变

	使用:
		step1:变量的声明,也叫定义
		step2:变量的访问,赋值和取值
	go的特性:
		静态语言:强类型语言
			go, java,c++, c#
		动态语言:弱类语言
				JavaScript, php, python, ruby
	 */
	//声明并赋值 :=格式,只能用在函数体内.局部变量可以省var,全局变量需要var
	var (
		//var关键字声明变量格式:  var 变量名 变量类型
		//var identifier type
		//变量赋值格式: 变量名称 = 值, 表达式, 函数等
		//声明赋值同时进行: var 变量名称 变量类型 = 值, 表达式, 函数等
		/*分组变量声明格式:
		var(
			i int
			j float32
			name string
		)
		 */
		num1 int
		num2 int
		num3 int
	)
	num1, num2, num3 = 1, 2, 3
	//打印函数调用语句.用于打印上述三个变量的值.
	var a int//第一种:定义变量
	var b int = 456//第二种定义变量,并赋值
	var (
		c int
		//c为int的类型,没有赋值,默认为0
		d float32
		e string
		//string默认为空
		f int = 1999
		h string = "保持耐心,是进步的基础"
		//函数体内也可以声明和赋值

	)
	var i, j, k int = 7,8,9
	//单个声明,多个变量
	var i1,j2,k3 =7,8,9.9
	var o = 10

	//自动推断类型
	r, s := 11, 12

	var t, _, v   = 1, 2.1, 3
	//_下划线是丢到垃圾桶,抛弃了.

	var vv int =8
	vx := float64(vv)
	//20211027类型转换

	fmt.Println(num1, num2, num3)
	fmt.Print("\n")
	//末尾增加空格
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c, d, e)
	fmt.Println(f)
	fmt.Println(h)
	fmt.Println(i, j, k )
	fmt.Println(i1, j2, k3 )
	fmt.Print("k3的类型为:")
	fmt.Print(reflect.TypeOf(k3))
	fmt.Print("\n")
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(r, s)
	fmt.Println(t, v)
	fmt.Print(vx, "\n")
	fmt.Println(reflect.TypeOf(vx))
}




