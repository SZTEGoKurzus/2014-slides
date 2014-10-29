package evenfib

func EvenFib() int {
	sum := 0 // HL
	a, b := 0, 1
	for b < 4000000 {
		a, b = b, a+b
		if b&1 == 0 {
			sum += b
		}
	}

	return sum
}
