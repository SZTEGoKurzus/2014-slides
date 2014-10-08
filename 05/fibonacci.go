// START1 OMIT
package main

import "fmt"

func main() {
	ch := make(chan uint64)
	go fibonacci(ch)
	for i := range ch {
		if i > 1000 {
			return
		}
		fmt.Println(i)
	}
}

// START2 OMIT
func fibonacci(ch chan<- uint64) {
	// END1 OMIT
	a, b := uint64(0), uint64(1)
	for {
		ch <- b
		a, b = b, a+b
	}
}

// END2 OMIT
