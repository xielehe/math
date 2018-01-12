package main

import (
	"fmt"
	"math"
)

func abcasyncMe(p int) bool {
	na := int(math.Sqrt(float64(p))) + 1
	channel := make(chan bool)
	enough := make(chan bool, na*na*na)
	for a := 0; a < na+1; a++ {
		for b := 0; b < na+1; b++ {
			go func() {
				for c := 0; c < na+1; c++ {
					if a*a+b*b+c*c == p {
						channel <- true
					} else {
						enough <- true
					}
				}
			}()
		}
	}
	for {
		select {
		case r := <-channel:
			return r
		default:
			if l := len(enough); l > na*na*na-1 {
				return false
			}
		}
	}
}

func main() {
	d := abcasync(41)
	// d := abcasync(23)
	fmt.Println(d)
}
