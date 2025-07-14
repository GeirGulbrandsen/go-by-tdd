package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	if sum != 4 {
		t.Errorf("Expected 4 but got %d", sum)
	}
}
