package evenfib

import "testing"

func TestEvenFib(t *testing.T) {
	ef := EvenFib()
	if ef != 4613732 {
		t.Fatalf("Expected 4613732, got: %d\n", ef)
	}
}
