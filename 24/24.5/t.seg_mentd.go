package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

func abcasync(p int) bool {
	na := int(math.Sqrt(float64(p))) + 2
	var waiteRet sync.WaitGroup
	var find int32 = 0
loop:
	for a := 0; a < na; a++ {
		for b := 0; b < na; b++ {
			if find == 1 {
				break loop
			}
			waiteRet.Add(1)
			go func() {
				defer waiteRet.Done()
				for c := 0; c < na; c++ {
					if find == 1 {
						break
					}
					if a*a+b*b+c*c == p {
						if atomic.CompareAndSwapInt32(&find, 0, 1) {
							fmt.Println(a, b, c)
							break
						}
					}
				}
			}()

		}
	}
	waiteRet.Wait()
	return atomic.LoadInt32(&find) == 1
}

func main() {
	d := abcasync(41)
	// d := abcasync(23)
	fmt.Println(d)
}
