package main

import "fmt"

func main() {
	/*
			循环结束:
				循环条件不满足,循环自动结束了
				但是可以通过break和continue来强制的结束循环

			循环控制语句
			break:策底的结束循环..终止`
	break:跳出循环体.break语句用于在结束其正常执行之前突然终止for循环
			continue:结束了某一次循环,下次继续..中止
		注意点:多层循环嵌套,break和continue,默认结束的是里层循环.
		如果想结束指定的某个循环,可以给循环贴标签(起名).
			break 循环标签名
			continue 循环标签名
	*/
	//死循环
	//for {
	//	fmt.Println("hello")
	//}
	//i := 1
	for i := 10; i >= 1; i-- {
		fmt.Print(i)
	}
	//fmt.Println("------------------")
	out:for j := 1; j <= 10; j++ {
		if j == 5 {
			break out//用out结束外循环
			//终止
		}
		fmt.Println(j)
		fmt.Println("main...over")

	in:for w := 1; w <= 10; w++ {
		if w == 5 {
			continue in //中止f
		}
		fmt.Println(w)
	}
	fmt.Println("------------------")
	for k := 1; k <= 3; k++ {
		for b := 1; b <= 3; b++ {
			if b == 2 {
				break
				}
			fmt.Printf("k:%d,b:%d\n", k, b)
			}
			fmt.Println("------------------")

	for d := 1; d <= 2; d++ {
		for a := 1; a <= 3; a++ {
			if a == 2 {
				continue
				}
				fmt.Printf("d:%d,a:%d\n", d, a)
					//fmt.Println("------------------")
				}
			}
		}
	}
}