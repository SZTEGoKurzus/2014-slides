package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)

	fmt.Println("CanSet?:", p.CanSet())
	fmt.Println("type:", p.Type())

	v := p.Elem()
	v.SetFloat(7.1)

	fmt.Println("new val:", v.Interface())
}
