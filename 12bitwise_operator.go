package main

import "fmt"

func main() {
	/*
		位运算符:
			将数值,转为二进制后,按位操作
		按位与&:
			对应位的值如果都为1才为1,有一个为0就为0
		按位或|:
			对应位的值如果都是0才为0,有一个为1就为1
		按位异或^:
			二元:a^b
				对应位的值不同为1,相同为0
			一元:^a
				按位取反:
					1--->0
					1--->1
		位清空:&^
			对于 a &^ b
				对于b上的每个数值
				如果为0,则取a对应位上的数值
				如果为1,则结果位就取0
	*/

	a := 60                        //转为二进制60 0011 1100
	b := 13                        //转为     13 0000 1100自动补全前面四个0000
	fmt.Printf("a:%d, %b\n", a, a) //""表示内容是一个字符串,:是原样输出
	fmt.Printf("b:%d, %b\n", b, b)
	/*
	   a: 60 0011 1100
	   b: 13 0000 1101
	   &     0000 1100
	*/
	res1 := a & b
	fmt.Println(res1)

	c := byte(0)
	d := byte(1)
	fmt.Print("c的值为:")
	fmt.Print(c)
	fmt.Print("\n")
	fmt.Print("d的值为:")
	fmt.Print(d)
	fmt.Print("\n")
	fmt.Print("c&d 的值为:")
	fmt.Print("c&d")
	fmt.Print("\n")
	fmt.Print("d&d 的值为:")
	fmt.Print(b & b)
	fmt.Print("\n")
	fmt.Print("c|d 的值为: ")
	fmt.Print(c | d)
	fmt.Print("\n")
	fmt.Print("c^d 的值为:")
	fmt.Print(c^d)
	fmt.Print("\n")
	fmt.Print("c^c 的值为: ")
	fmt.Print(c^c)

	fmt.Print("\n")
	fmt.Print("d^d 的值为: ")
	fmt.Print(d^d)
	e := 0
	e = a<<2 // 240 = 0011 0001
	fmt.Printf("e的值为 %d\n", e)

	e = a >>2 //15 = 0000 1111
	fmt.Printf("e的值为 %d\n", e)
}
