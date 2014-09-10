package main

import "fmt"

func main() {
	v := []int{0, 1, 2, 3, 5, 6, 2, -1}
	fmt.Println(haromosszege0(v))
}

func haromosszege0(input []int) bool {
	for i := range input {
		for j := range input {
			for k := range input {
				if i != j && i != k && j != k && input[i]+input[j]+input[k] == 0 {
					return true
				}
			}
		}
	}

	return false
}
