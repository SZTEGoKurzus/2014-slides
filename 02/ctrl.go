package main

func main() {
	x := 5
	// START1 OMIT
	if x == 2 {
		doSomething()
	}
	// END1 OMIT

	// START2 OMIT
	if y := 0; x == y {
		doSomething()
	}
	// END2 OMIT

	// START3 OMIT
	for x != 5 {
		doSomething()
	}
	// END3 OMIT

	// START4 OMIT
	for y := 0; y < x; y++ {
		doSomething()
	}
	// END4 OMIT

	arr := []int{}
	// START5 OMIT
	for i := range arr {
		doSomething2(i)
	}

	for i, v := range arr {
		doSomething2(i + v)
	}
	// END5 OMIT

	// START6 OMIT
	switch x {
	case 0:
		doSomething()
	case 1, 2:
		doSomething()
	case 3:
		doSomething()
		fallthrough
	default:
		doSomething()
	}
	// END6 OMIT

	var xy interface{} = 0
	// START7 OMIT
	switch xy.(type) {
	case nil:
		doSomething()
	case int:
		doSomething()
	case string:
		doSomething()
	default:
		doSomething()
	}
	// END7 OMIT
}

// START8 OMIT
func doSomething() {

}

// END8 OMIT

// START9 OMIT
func doSomething2(i int) {

}

// END9 OMIT

// START10 OMIT
func doSomething3(i, j int, s string, z ...[]string) {

}

// END10 OMIT

// START11 OMIT
func doSomething4() string {
	return ""
}

// END11 OMIT

// START12 OMIT
func doSomething5() (string, int) {
	return "", 0
}

// END12 OMIT

// START13 OMIT
func doSomething6() (s string, i int) {
	s = ""
	i = 0
	return
}

// END13 OMIT

func doSomething7() string {
	// START14 OMIT
	s, _ := doSomething6()
	// END14 OMIT
	return s
}
