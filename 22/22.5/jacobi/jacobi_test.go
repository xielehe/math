package jacobi

import (
	"testing"
)

// TestJacobi1 ...
func TestJacobi1(t *testing.T) {
	if i, e := Jacobi(5, 3593); i != -1 || e != nil { //try a unit test on function
		t.Error("测试没通过")
	} else {
		t.Log("通过了") //记录一些你期望记录的信息
	}
}

// Jacobi(5, 3889) // 1
// Jacobi(14, 137) // 1
// Jacobi(55, 179) // -1
// Jacobi(299, 397) // -1
// Jacobi(37603, 48611) // 1
// Jacobi(85, 101) // 1
// Jacobi(29, 541) // -1
// Jacobi(101, 1987) // 1
// Jacobi(31706, 43789) // -1
