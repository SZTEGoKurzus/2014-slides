package main

import "fmt"

func main() {
	// START OMIT
	ch := make(chan bool)

	go func() {
		select {
		case ch <- true:
		case ch <- false:
		}
	}()

	fmt.Println(<-ch)
	// END OMIT
}
