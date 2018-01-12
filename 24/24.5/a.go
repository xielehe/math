package main

import (
	"fmt"
	"math"
	"sync"
	"time"

	"../../lib/mymath"
)

func abc(p int) (int, int, int, bool) {
	na := int(math.Sqrt(float64(p))) + 2
	nb, nc := na, na
	for a := 0; a < na; a++ {
		for b := 0; b < nb; b++ {
			for c := 0; c < nc; c++ {
				if a*a+b*b+c*c == p {
					return a, b, c, true
				}
			}
		}
	}
	return 0, 0, 0, false
}

func abcAsync(p int) (int, int, int, bool) {
	var wg sync.WaitGroup
	na := int(math.Sqrt(float64(p))) + 2
	channel := make(chan bool)
	nb, nc := na, na
	for a := 0; a < na; a++ {
		for b := 0; b < nb; b++ {

			go func() {
				for c := 0; c < nc; c++ {
					if a*a+b*b+c*c == p {
						channel <- true
					}
				}
			}()

		}
	}
	return 0, 0, 0, false
}

//ABC ...
type ABC struct {
	a int
	b int
	c int
	p int
}

func f(N int) ([]ABC, []int) {
	result := []ABC{}
	non := []int{}
	r := make(chan ABC, N)
	rn := make(chan int, N)
	var wg sync.WaitGroup
	for n := 0; n < N+1; n++ {
		wg.Add(1)
		go func(x int, r chan ABC) {
			defer wg.Done()
			if mymath.IsPrime(x) {
				if a, b, c, has := abc(x); has {
					r <- ABC{a: a, b: b, c: c, p: x}
				} else {
					rn <- x
				}
			}
		}(n, r)
	}
	wg.Wait()
	close(r)
	close(rn)
	for i := range r {
		result = append(result, i)
	}
	for i := range rn {
		non = append(non, i)
	}
	return result, non
}

func testa(t []int) bool {
	for _, i := range t {
		if i%8 != 7 {
			return false
		}
	}
	return true
}

//异步代码
func main() {
	// fmt.Println(a(50)) //[{1 1 3} {1 2 7} {1 3 13} {2 3 19} {1 5 31} {3 4 37} {1 6 43}]
	// 通过观察发现，满足要求的素数都满足模3余1(除了3),下面通过更多的素数测试,以便增加猜测正确的概率
	start := time.Now()
	limit := 100000

	_, non := f(limit)
	// for _, ii := range container {
	// 	fmt.Println(ii)
	// }
	is := testa(non)
	fmt.Println(is)
	end := time.Now()
	fmt.Printf("LongCalculation took this amount of time: %s\n", end.Sub(start))
	// 通过测试，发现<1000000的结果都满足猜想.
}
