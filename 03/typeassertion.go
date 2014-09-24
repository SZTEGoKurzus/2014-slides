package main

import "fmt"

// START OMIT
func CastToInt(i interface{}) *int {
	v := i.(int) // HL
	return &v
}

func main() {
	i := struct{}{}
	if casted := CastToInt(i); casted != nil {
		fmt.Println(*CastToInt(i))
	} else {
		fmt.Println(0)
	}
}

// END OMIT
