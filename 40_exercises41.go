package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 2; i <= 100; i ++  {
		flag := true //记录i是不是素数
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				flag = false //不是素数了.
				break
			}
		}
		if flag { //==true
			fmt.Println(i)
		}
	}
}
