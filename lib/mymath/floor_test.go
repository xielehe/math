package mymath

import (
	"testing"
)

// TestJacobi1 ...
func TestFloor(t *testing.T) {
	if f(5.7, 5, t) &&
		f(22/7, 3, t) &&
		f(4, 4, t) &&
		f(-1.3, -2, t) {
		t.Log("PASS")
	}
}

func f(x float64, res int, t *testing.T) bool {
	if i := Floor(x); i != res {
		t.Errorf("Floor(%f), Expected %d,got %d", x, res, i)
		return false
	}
	return true
}
