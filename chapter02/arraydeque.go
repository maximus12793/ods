package ods

import "math"

type ArrayDeque struct {
	a []interface{}
	j int
	n int
}

func (a *ArrayDeque) Size() int {
	return a.n
}

func (a *ArrayDeque) Init() {
	a.a = make([]interface{}, 1)
	a.j = 0
	a.n = 0
}

func (a *ArrayDeque) Get(i int) interface{} {
	return a.a[(i+a.j)%len(a.a)]
}

func (a *ArrayDeque) Set(i int, x interface{}) interface{} {
	y := a.a[(i+a.j)%len(a.a)]
	a.a[(i+a.j)%len(a.a)] = x
	return y
}

func (a *ArrayDeque) Add(i int, x interface{}) {
	if a.n == len(a.a) {
		a.resize()
	}

	if i < a.n/2 {
		// Case 1: i < n/2.
		a.j = (a.j - 1) % len(a.a)
		for k := 0; k < i-1; k++ {
			// Shift elements to the left by 1 position.
			a.a[(a.j+k)%len(a.a)] = a.a[(a.j+k+1)%len(a.a)]
		}
	} else {
		// Case 2: i >= n/2.
		for k := a.n; k > i+1; k-- {
			// Shift elements to the right by 1 position.
			a.a[(a.j+k)%len(a.a)] = a.a[(a.j+k-1)%len(a.a)]
		}
	}
	a.a[(a.j+i)%len(a.a)] = x
	a.n++
}

func (a *ArrayDeque) Remove(i int) interface{} {
	x := a.a[(a.j+i)%len(a.a)]

	if i < a.n/2 {
		// Case 1: i < n/2.
		for k := i; k > 1; k-- {
			// Shift elements to the right by 1 position.
			a.a[(a.j+k)%len(a.a)] = a.a[(a.j+k-1)%len(a.a)]
		}
		a.j = (a.j + 1) % len(a.a)
	} else {
		// Case 2: i >= n/2.
		for k := i; k < a.n-1; k++ {
			// Shift elements to the left by 1 position.
			a.a[(a.j+k)%len(a.a)] = a.a[(a.j+k+1)%len(a.a)]
		}
	}
	a.n--
	if len(a.a) >= (3 * a.n) {
		a.resize()
	}
	return x
}

func (a *ArrayDeque) resize() {
	b := make([]interface{}, int(math.Max(1, float64(2*a.n))))
	for k := 0; k < a.n; k++ {
		b[k] = a.a[(a.j+k)%len(a.a)]
	}
	a.a = b
	a.j = 0
}
