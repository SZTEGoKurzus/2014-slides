package main

import "fmt"

func main() {
	sum := 0
	a, b := 0, 1
	for b < 4000000 {
		a, b = b, a+b
		if b&1 == 0 {
			sum += b
		}
	}

	fmt.Println(sum)
}
