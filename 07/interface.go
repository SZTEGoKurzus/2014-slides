package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)

	fmt.Println("real value:", v.Interface())

	fmt.Println("real value casted:", v.Interface().(float64))
}
