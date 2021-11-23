package main

import "fmt"

/*
赋值运算符:
	=,+=, -=,*=,/=, %=, <<=,>>=, &=, |=, ^=...
	=,把=右侧的数值,赋值给=左边的变量

	+=, a += b,相当于a = a + b
	a++, a +=1
	优先级运算符优先级
*/
func main() {

	var a int
	a = 3
	fmt.Println(a) //3

	a += 4 // a = a + 4
	fmt.Println(a)//7

	a -= 3
	fmt.Println(a) //7-3=4
	a *= 2
	fmt.Println(a)//4*2=8
	a /= 3
	fmt.Println(a)//8/3=2.7
	a %= 1
	fmt.Println(a)//任何数除以1,取模数都是0
}
