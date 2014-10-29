package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{47, "Foo"}
	s := reflect.ValueOf(&t).Elem()
	tt := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, tt.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(42)
	s.Field(1).SetString("Bar")
	fmt.Println("t: ", t)
}
