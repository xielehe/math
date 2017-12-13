package jacobi

import (
	"testing"
)

// TestJacobi1 ...
func TestJacobi1(t *testing.T) {
	if onet(5, 3593, -1, t) &&
		onet(5, 3889, 1, t) &&
		onet(14, 137, 1, t) &&
		onet(299, 397, -1, t) &&
		onet(37603, 48611, 1, t) &&
		onet(85, 101, 1, t) &&
		onet(101, 1987, 1, t) &&
		onet(29, 541, -1, t) &&
		onet(31706, 43789, -1, t) &&
		onet(55, 179, -1, t) {
		t.Log("PASS")
	}
}

func onet(up int64, down uint64, res int8, t *testing.T) bool {
	if i, e := Jacobi(up, down); i != res || e != nil {
		t.Errorf("Jacobi(%d %d), Expected %d,got %d", up, down, res, i)
		return false
	}
	return true
}
