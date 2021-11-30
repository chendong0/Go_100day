package utils

import "fmt"

func Count() {
	fmt.Println("utils包下的Count()函数")

}
func int() {
	fmt.Println("utils包下的另一个init()函数")
}

func init() {
	fmt.Println("utils包的init()函数,用于初始化一些信息..")
}
