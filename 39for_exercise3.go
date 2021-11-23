package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 100; i < 1000; i+=1 { //i++ i+=1
		x := i / 100
		y := i / 10 % 10
		z := i % 10

		if math.Pow(float64(x), 3) + math.Pow(float64(y), 3)  + math.Pow(float64(z), 3) == float64(i) {
			fmt.Println(i)

		}
	}
	fmt.Println("---------")

	for a := 1; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				n := a*100 + b*10 + c*1
				if a*a*a+b*b*b+c*c*c == n {
					fmt.Println(n)
				}
			}
		}
	}
}
