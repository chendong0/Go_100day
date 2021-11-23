package main

import "fmt"

func main() {
	//便利,决定处理第几行
	for y := 1; y <= 9; y++ {
		//遍历,决定这一行有多少列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d \t", x, y, x*y) //不知道printf的语法和用法?
			//%d *空格隔开,就可以有空格和美观,\t表示一个tab制表符
			//fmt.Printf("\n")
			//fmt.Println()
		}
		//fmt.Println()
		//手动生成回车
		fmt.Printf("\n")
		fmt.Printf("\n")

	}
	////fmt.Printf("\n")

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v* %v=%v \t ",j,i ,j * i)
		}
		fmt.Println(" ")
	}

}
