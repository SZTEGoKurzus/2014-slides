package main

import "fmt"

func main() {
	// START0 OMIT
	var a float64
	// END0 OMIT
	// START1 OMIT
	var x int = 0
	// END1 OMIT
	// START2 OMIT
	var y = 1
	// END2 OMIT
	// START3 OMIT
	z := 2
	// END3 OMIT
	// START4 OMIT
	e, f := 1, 2
	// END4 OMIT
	// START5 OMIT
	f, e = e, f
	// END5 OMIT

	fmt.Printf("A: %f X: %d Y: %d Z: %d E: %d F: %d\n", a, x, y, z, e, f)
}
