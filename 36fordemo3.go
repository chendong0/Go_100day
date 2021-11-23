package main

import "fmt"

func main() {
	/*
		循环嵌套:多层循环嵌套在一起

		题目一:
		*****
		*****
		*****
		*****
		*****

		Print()
		Printf()
		Println()

		题目二:
		打印乘法表
	*/
	//第一行
	//1.打印*
	for j := 0; j < 5; j++ {
		fmt.Print("*")

	}


	//2.换行
	fmt.Println()



	//第二行
	//1.打印*
	for j := 0; j < 5; j++ {
		fmt.Print("*")
	}
	fmt.Println("------>")
	//2.换行
	fmt.Println()

	for i := 5; i <= 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*") //j:0,1,2,3,4
		}
		//2.换行
		fmt.Println()
	}
}
