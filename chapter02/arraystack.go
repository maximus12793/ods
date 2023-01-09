package ods

import "math"

type ArrayStack struct {
	a []interface{}
	n int
}

func (a *ArrayStack) Size() int {
	return a.n
}

func (a *ArrayStack) Init(i int) {
	a.a = make([]interface{}, i)
	a.n = i
}

func (a *ArrayStack) Get(i int) interface{} {
	return a.a[i]
}

func (a *ArrayStack) Set(i int, x interface{}) interface{} {
	y := a.a[i]
	a.a[i] = x
	return y
}

func (a *ArrayStack) Add(i int, x interface{}) {
	if a.n == len(a.a) {
		a.resize()
	}
	copy(a.a[i+1:], a.a[i:a.n])
	a.a[i] = x
	a.n++
}

func (a *ArrayStack) Remove(i int) interface{} {
	x := a.a[i]
	copy(a.a[i:a.n-1], a.a[i+1:a.n])
	a.n--
	if len(a.a) >= 3*a.n {
		a.resize()
	}
	return x
}

func (a *ArrayStack) resize() {
	b := make([]interface{}, int(math.Max(1, float64(2*a.n))))
	copy(b, a.a)
	a.a = b
}
