// START1 OMIT
package main

import "fmt"

func main() {
	// func fibonacci() func() int {}
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// END1 OMIT

// START2 OMIT
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// END2 OMIT
