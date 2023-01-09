package ods

import "math"

// Adheres to FIFO queue.
type ArrayQueue struct {
	a []interface{}
	j int
	n int
}

func (a *ArrayQueue) Size() int {
	return a.n
}

func (a *ArrayQueue) Init(i int) {
	a.a = make([]interface{}, i)
	a.j = 0
	a.n = 0
}

func (a *ArrayQueue) Add(i int, x interface{}) bool {
	if a.n+1 >= len(a.a) {
		a.resize()
	}
	a.a[(a.j+a.n)%len(a.a)] = x
	a.n++
	return true
}

func (a *ArrayQueue) Remove() interface{} {
	x := a.a[a.j]
	a.j = (a.j + 1) % len(a.a)
	a.n--
	if len(a.a) >= 3*a.n {
		a.resize()
	}
	return x
}

func (a *ArrayQueue) resize() {
	b := make([]interface{}, int(math.Max(1, float64(2*a.n))))
	for k := 0; k < a.n; k++ {
		b[k] = a.a[(a.j+k)%len(a.a)]
	}
	a.a = b
	a.j = 0
}
