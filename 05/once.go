package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	var o sync.Once

	for i := 0; i < 100; i++ {
		o.Do(func() {
			fmt.Println("Csak egyszer")
		})
	}
	// END OMIT
}
