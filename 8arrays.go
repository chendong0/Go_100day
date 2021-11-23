package main

import "fmt"

func main() {
	//定义数组的五种方法
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int //定义grid为4行,5列的,4个长度为5的int数
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	fmt.Println(len(arr1),cap(arr1))

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for i := range arr3 { //range关键字可以获得下标
		fmt.Println(arr3[i])
	}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	for _, v := range arr3 { //不要下标,只要数值
		fmt.Println(v)
	}
	//数组的遍历
	numbers := [6]int{1, 3, 2, 5, 8, 4}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

}