package main

import (
	"fmt"
	"math"
	"sync"
	"time"

	"../../lib/mymath"
)

func a2abb2(p int) (int, int, bool) {
	n := p - 1
	for i := -n; i < n; i++ {
		for j := -n; j < n; j++ {
			if i*i+i*j+j*j == p {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}

type a2ABB2 struct {
	a int
	b int
	p int
}

func a(N int) []a2ABB2 {
	result := []a2ABB2{}
	r := make(chan a2ABB2, N)
	wg := sync.WaitGroup{}
	for n := 0; n < N+1; n++ {
		wg.Add(1)
		go func(x int, r chan a2ABB2) {
			isHas := false
			if mymath.IsPrime(x) {
				a, b, has := a2abb2(x)
				if has {
					r <- a2ABB2{a: a, b: b, p: x}
					wg.Done()
					isHas = true
				}
			}
			if !isHas {
				wg.Done()
			}
		}(n, r)
	}
	wg.Wait()
	close(r)
	for i := range r {
		result = append(result, i)
	}
	return result
}

func testa(t []a2ABB2) bool {
	for _, i := range t {
		if i.p != 3 && i.p%3 != 1 {
			return false
		}
	}
	return true
}

//异步代码
func main() {
	// fmt.Println(a(50)) //[{1 1 3} {1 2 7} {1 3 13} {2 3 19} {1 5 31} {3 4 37} {1 6 43}]
	// 通过观察发现，满足要求的素数都满足模3余1(除了3),下面通过更多的素数测试,以便增加猜测正确的概率
	// start := time.Now()
	// limit := 10000
	// container := a(limit)
	// isMod3r1 := testa(container)
	// fmt.Println(isMod3r1)
	// end := time.Now()
	// fmt.Printf("LongCalculation took this amount of time: %s\n", end.Sub(start))
	// 通过测试，发现<1000000的结果都满足猜想.

	// 那么，模3余1的素数p,一定满足p=a^2+ab+b^2吗，虽然我暂时不打算证明，但是下面测试一下。
	start := time.Now()
	limit := 1000000
	container := inverseA(limit)
	isA2abb2 := inverseTesta(container)
	fmt.Println(isA2abb2)
	end := time.Now()
	fmt.Printf("LongCalculation took this amount of time: %s\n", end.Sub(start))
	// 通过测试，发现<1000000的结果都满足猜想.
}

func inverseA(N int) []int {
	result := []int{}
	r := make(chan int, N)
	var wg sync.WaitGroup
	for n := 0; n < N+1; n++ {
		wg.Add(1)
		go func(x int, r chan int) {
			defer wg.Done()
			if x%3 == 1 && mymath.IsPrime(x) {
				r <- x
			}
		}(n, r)
	}
	wg.Wait()
	close(r)
	for i := range r {
		result = append(result, i)
	}
	return result
}

func inverseTesta(t []int) bool {
	for _, i := range t {
		if _, _, b := simpleA2abb2(i); !b {
			// if _, _, b := a2abb2(i); !b {
			return false
		}
	}
	return true
}

func simpleA2abb2(p int) (int, int, bool) {
	n := int(math.Sqrt(float64(p))) + 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i*i+i*j+j*j == p {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}
