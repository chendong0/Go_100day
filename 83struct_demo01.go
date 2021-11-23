package main

import "fmt"

func main() {
	/*
		结构体:是有一系列具有相同类型或不同类型的数据构成的数据集合
			结构体成员是由一系列的成员变量构成,这些成员变量也被称为"字段"
	*/
	//1.方法一
	var p1 Person
	fmt.Println(p1)
	p1.name = "王健林"
	p1.age = 40
	p1.sex = "男"
	p1.address = "深圳市"
	fmt.Printf("姓名:%s, 年龄: %d, 性别: %s,地址:%s\n", p1.name, p1.age, p1.sex, p1.address)

	//2.方法二
	p2 := Person{}
	p2.name = "guoguo"
	p2.age = 3
	p2.sex = "女"
	p2.address = "南宁市"
	fmt.Printf("姓名:%s, 年龄: %d, 性别: %s,地址:%s\n", p2.name, p2.age, p2.sex, p2.address)

	//3.方法三
	p3 := Person{name: "周星驰", age: 55, sex: "男", address: "香港"}
	fmt.Println(p3)
	p4 := Person{
		name:    "马云",
		age:     50,
		sex:     "男",
		address: "杭州市",
	}
	fmt.Println(p4)

	//4.方法四
	p5:=Person{"陈东",31,"男","南宁市"}
	fmt.Println(p5)
}

//1.定义结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}
