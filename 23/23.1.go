package main

import (
	"fmt"
	"math"

	"../lib/mymath"
)

func main() {
	a := mymath.Floor(-7 / 3)                        // -3
	b := mymath.Floor(math.Sqrt(23))                 // 4
	c := mymath.Floor(math.Pi * math.Pi)             // 9
	d := mymath.Floor(math.Sqrt(73) / math.Cbrt(19)) // 3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
