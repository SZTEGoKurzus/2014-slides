package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	fmt.Println("CanSet?:", v.CanSet())

	v.SetFloat(7.1)
}
