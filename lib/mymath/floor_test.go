package mymath

import (
	"testing"
)

// TestJacobi1 ...
func TestJacobi1(t *testing.T) {
	if onet(5.7, 5, t) &&
		onet(22/7, 3, t) &&
		onet(4, 4, t) &&
		onet(-1.3, -2, t) {
		t.Log("PASS")
	}
}

func onet(x float32, res int, t *testing.T) bool {
	if i := Floor(x); i != res {
		t.Errorf("Floor(%f), Expected %d,got %d", x, res, i)
		return false
	}
	return true
}
