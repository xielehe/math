package main

import (
	"fmt"
	"math"
	"sync"
	"time"

	"../../../lib/mymath"
)

type A2B struct {
	a int
	b int
	p int
}

func a(N int) []A2B {
	result := []A2B{}
	r := make(chan A2B, N)
	var wg sync.WaitGroup
	for n := 0; n < N+1; n++ {
		wg.Add(1)
		go func(x int, r chan A2B) {
			defer wg.Done()
			if mymath.IsPrime(x) {
				if a, b, has := a2b(x); has {
					r <- A2B{a: a, b: b, p: x}
				}
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

func a2b(p int) (int, int, bool) {
	SUPa := int(math.Sqrt(float64(p))) + 2
	SUPb := int(math.Sqrt(float64(p/2))) + 2
	for a := 0; a < SUPa; a++ {
		for b := 0; b < SUPb; b++ {
			if a*a+2*b*b == p {
				return a, b, true
			}
		}
	}
	return 0, 0, false
}

func testa(t []A2B) bool {
	for _, i := range t {
		if p := i.p; p != 3 && p != 2 && p%6 != 1 && p%6 != 5 {
			return false
		}
	}
	return true
}

func main() {
	// fmt.Println(a(100)) //{1 1 3} {0 1 2} {3 1 11} {3 2 17} {1 3 19} {3 4 41} {5 3 43} {3 5 59} {7 3 67} {1 6 73} {9 1 83} {9 2 89} {5 6 97}
	// 通过观察发现，满足要求的素数都满足模6余|1|(除了3和2),下面通过更多的素数测试,以便增加猜测正确的概率
	start := time.Now()
	limit := 1000000
	container := a(limit)
	isMod6r1 := testa(container)
	fmt.Println(isMod6r1)
	end := time.Now()
	fmt.Printf("LongCalculation took this amount of time: %s\n", end.Sub(start))
	// 通过测试，发现<1000000的结果都满足猜想.

	// 那么，模6余|1|的素数p,一定满足p=a^2+2b^2吗，虽然我暂时不打算证明，但是下面测试一下。
	inverselimit := 100
	inversecontainer := inverseA(inverselimit)
	fmt.Print(inversecontainer)
	// [1 5 7 11 13 17 19 23 25 29 31 37 41 43 47 49 53 55 59 61 67 71 73 79 83 85 89 91 97]
	// 通过测试，发现不满足猜想.
}

func inverseA(N int) []int {
	result := []int{}
	r := make(chan int, N)
	var wg sync.WaitGroup
	for n := 0; n < N+1; n++ {
		wg.Add(1)
		go func(x int, r chan int) {
			defer wg.Done()
			if m := x % 6; m == 1 || m == 5 && mymath.IsPrime(x) {
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
