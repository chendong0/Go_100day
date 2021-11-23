package main

import "fmt"

func main() {
	/*
		算术运算符: +, -, *,/,%, ++,--
	++: 给自己加1
	--: 给自己减1
	*/
	a := 10
	b := 3
	sum := a + b
	fmt.Printf("%d + %d =%d\n", a, b, sum)

	sub := a - b
	fmt.Printf("%d + %d =%d\n", a, b, sub)

	mul := a * b
	fmt.Printf("%d + %d =%d\n", a, b, mul)

	div := a / b //取商
	mod := a % b //取余,取模
	fmt.Printf("%d / %d = %d\n", a, b, div)
	fmt.Printf("%d / %d = %d\n", a, b, mod)

	c := 3
	d := 2
	c++
	d ++
	//++和--不能放在一起,只能单独运算
	//错误用法sum2 := c + c++ + c--
	fmt.Println(c,d)

	c-- //取了上一行输出的结果,再运算
	d--
	fmt.Println(c,d)

}
