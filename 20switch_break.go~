package main

import "fmt"

func main() {
	/*
		switch中的break和fallthrough语句
		break:可以使用在switch中,也可以只用在for循环中
			强制结束case语句,从而结束switch分支
	fallthrough:用于穿透switch
			当switch中某个case匹配成功之后,就执行该case语句
			如果遇到fallthrough,那么后面紧邻的case,无需匹配,执行穿透执行.
			(只能放在case后面的一行)
	*/
	n := 2
	switch n {
	case 1:
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
	case 2:
		fmt.Println("我是熊2")

		 	fmt.Println("我是熊2")
		break //用于强制结束case,意味着switch被强制结束
		fmt.Println("我是熊2")

	case 3:
		fmt.Println("我是光头强")
		fmt.Println("我是光头强")
		fmt.Println("我是光头强")
	}
	fmt.Println("_________________")
	m :=2
	switch m {
	case 1:
		fmt.Println("第一季度..")
	case 2:
		fmt.Println("第2季度..")
		fallthrough
	case 3:
		fmt.Println("第3季度..")
	case 4:
		fmt.Println("第4季度..")
	}
	fmt.Println("main...over....")
}
