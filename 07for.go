package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		//for条件里不需要括号
		sum += i
	}
	fmt.Print(sum)
	/*
		for 循环：某些代码会被多次的执行
		语法：
			for 表达式1；表达式2；表达式3｛
			｝
	*/

	for j := 1; j <= 10; j++ {
		fmt.Println("hello world!!")
	}
}
