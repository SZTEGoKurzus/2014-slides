package main

import "fmt"

// START OMIT
func CastToInt(i interface{}) *int {
	if v, ok := i.(int); ok { // HL
		return &v
	} else {
		return nil
	}
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
