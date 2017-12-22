package main

import (
	"fmt"
	"math"
	"time"

	"../../lib/mymath"
)

func a2abb2(p int) (int, int, bool) {
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

type a2ABB2 struct {
	a int
	b int
	p int
}

func a(N int) []a2ABB2 {
	result := []a2ABB2{}
	for n := 0; n < N+1; n++ {
		if mymath.IsPrime(n) {
			a, b, has := a2abb2(n)
			if has {
				result = append(result, a2ABB2{a: a, b: b, p: n})
			}
		}
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

func main() {
	// fmt.Println(a(50)) //[{1 1 3} {1 2 7} {1 3 13} {2 3 19} {1 5 31} {3 4 37} {1 6 43}]
	// 通过观察发现，满足要求的素数都满足模3余1(除了3),下面通过更多的素数测试,以便增加猜测正确的概率
	start := time.Now()
	limit := 1000000
	container := a(limit)
	isMod3r1 := testa(container)
	fmt.Println(isMod3r1)
	end := time.Now()
	fmt.Printf("longCalculation took this amount of time: %s\n", end.Sub(start))
	//通过测试，发现<1000000的结果都满足猜想.
	//TODO 实现异步代码

}
