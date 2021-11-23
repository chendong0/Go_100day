package main

import "fmt"

func main() {
	//1.创建一个普通的数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	//2.创建一个指针,存储该数组的地址--->数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)         //&[1,2,3,4]代表的是数组的指针
	fmt.Printf("%p\n", p1)  //数组arr1的地址
	fmt.Printf("%p\n", &p1) //p1指针自己的地址
	fmt.Printf("%p\n,%p\n,%p\n", &p1[0], &p1[1], p1[2])

	//3.根据数组指针,操作数组
	(*p1)[0] = 100
	//*p1[0]=789//invalid indirect of p1[0] (type int)要用括号
	fmt.Println(arr1)

	p1[0] = 200 //简化写法,数组指针
	fmt.Println(arr1)

	//4.指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2) //[1,2,3,4]
	fmt.Println(arr3)
	arr2[0] = 100
	fmt.Println(arr2)
	fmt.Println(a)
	*arr3[0] = 200
	fmt.Println(arr3)
	fmt.Println(a)

	b = 1000
	fmt.Println(arr2)
	fmt.Println(arr3)
	for i:=0;i<len(arr3);i++{
		fmt.Println(*arr3[i])
	}

}
