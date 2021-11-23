package main

import (
	"fmt"
)

func main() {
	/*
		for循环练习题:
		练习1:打印58-2数字
		练习2:求1-100的和
		练习3:打印1-100内,能够被3整除,但是不能被5整除的数字,统计被打印的个数,每行打印5个
	*/
	for i := 58; i >= 2 && i <= 58; i-- {
		fmt.Println(i)
	}
	sum := 0
	for j := 1; j <= 100; j++ {
		sum += j
		//j += sum

	}
	fmt.Println("1-100的和为", sum)

	count := 0  //计数器
	for k := 1; k <= 100; k++ {
		//fmt.Println(k)
		if k%3 == 0 && k%5 != 0 { //嵌套循环语句
			fmt.Print(k,"\t")
			count++ //5,10,15,20...
			if count % 5 == 0 {
				fmt.Print("\n")
			}
		}

	}
	fmt.Println()
	fmt.Println("count-->", count)

	fmt.Println(`我是熊大
我是熊二 (波浪线下的左上角反引号,可以自动理解为换行.)
`)
}
