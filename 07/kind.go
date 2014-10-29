package main

import (
	"fmt"
	"reflect"
)

type lolfloat float64

func main() {
	var x lolfloat = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("kind:", reflect.ValueOf(x).Kind())
	fmt.Println("value:", reflect.ValueOf(x).Float())
}
