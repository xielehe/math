package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

type S struct {
	a int
	b int
	c int
	p int
}

func NewS(pa, pb, pc, pp int) *S {
	return &S{
		a: pa,
		b: pb,
		c: pc,
		p: pp,
	}
}

func (s *S) isMatch() bool {
	return s.a*s.a+s.b*s.b+s.c*s.c == s.p
}

func abcasync(ch chan *S, chResult chan *S, p int) *S {
	na := int(math.Sqrt(float64(p))) + 2
	for a := 0; a < na; a++ {
		for b := 0; b < na; b++ {
			for c := 0; c < na; c++ {
				select {
				case r := <-chResult:
					close(ch)
					return r
				default:
					ch <- (NewS(a, b, c, p))
				}
			}
		}
	}
	return nil
}

func main() {
	pNum := runtime.NumCPU()
	ch := make(chan *S, 100)
	chResult := make(chan *S)

	var wg sync.WaitGroup
	wg.Add(pNum)
	for i := 0; i < pNum; i++ {
		go func() {
			defer wg.Done()
			for s := range ch {
				if s.isMatch() {
					chResult <- s
				}
			}
		}()
	}
	d := abcasync(ch, chResult, 6)
	if d == nil {
		fmt.Println("not found")
	} else {
		fmt.Println(d)
	}
	wg.Wait()
}
