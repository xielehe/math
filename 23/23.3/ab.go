package main

import (
	"../../lib/mymath"
)

func g(x float64) int {
	return mymath.Floor(x) + mymath.Floor(x+0.5)
}

// (a)
func main() {
	g0 := g(0)        // 0
	g025 := g(0.25)   // 0
	g05 := g(0.5)     // 1
	g1 := g(1)        // 2
	g2 := g(2)        // 4
	g25 := g(2.5)     // 5
	g2499 := g(2.499) // 4
	println(
		g0,
		g025,
		g05,
		g1,
		g2,
		g25,
		g2499)
}

// (b)
// g(x) = [2x]
