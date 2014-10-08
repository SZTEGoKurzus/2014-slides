package main

import "fmt"

// START1 OMIT
func generate(ch chan<- uint64) {
	for i := uint64(2); ; i++ {
		ch <- i
	}
}

func filter(src <-chan uint64, dst chan<- uint64, prime uint64) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
}

// END1 OMIT

// START2 OMIT
func main() {
	ch := make(chan uint64)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan uint64)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

// END2 OMIT
