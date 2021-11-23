package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		if i == 3{
			break
			//continue
		}

		fmt.Println(i)
	}
	for k := 1; k <= 5; k++ {
		if k == 3{
			//break
			continue
		}

		fmt.Println(k)
	}

}
