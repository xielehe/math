package main

import (
	"fmt"

	"../../22.5/jacobi"
)

func main() {
	i, e := jacobi.Jacobi(11, 1729)
	fmt.Print(i, e) //  -1
}
