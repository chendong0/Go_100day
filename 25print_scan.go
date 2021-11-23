package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		输入和输出:
			fmt包:输入,输出

				输出:
		   			Print() //打印
					Printf() //格式化打印
					Println() //打印后换行

				格式化打印占位符:
					%v,原样输出
					%T,打印类型
					%t,bool类型
					%s,字符串
					%f,浮点
					%d,10进制的整数
					%b,2进制的整数
					%o,8进制的整数
					%x,%V,16进制
						%v: 0-9, a-f
						%V: 0-9, A-F
					%c,打印字符
					%p,打印地址
					%%,打印百分号%
					......
	*/
	a := 100           //int
	b := 3.14          //float64
	c := true          //bool
	d := "Hello World" //srting
	e := "Ruby"        //srting
	f := 'A'
	fmt.Printf("%T,%b\n", a, a)
	fmt.Printf("%T,%F\n", b, b)
	fmt.Printf("%T,%t\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%s\n", e, e)
	fmt.Printf("%T,%d,%c\n", f, f, f)
	fmt.Println("-----------------------")
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Println("-----------------------")

	/*
		输入;
			Scanln()

	*/
	var x int
	var y float64
	fmt.Println("请输入一个整数,一个浮点数类型:")
	fmt.Scanln(&x, &y) //读取键盘的输入,通过操作地址,赋值给x和y  阻塞式的,需要输入程序才继续
	fmt.Printf("a的数值: %d, b的数值: %f\n", a, b)

	fmt.Scanf("%d,%f", &x, &y)
	fmt.Printf("x:%d,y:%f\n", x, y)

	fmt.Println("请输入一个字符串: ")
	reader:=bufio.NewReader(os.Stdin) //os包下的标准输入
	s1,_:=reader.ReadString('\n')
	fmt.Println("读到的数据:",s1)
}
