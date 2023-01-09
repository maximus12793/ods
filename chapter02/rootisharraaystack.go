package ods

import (
	"fmt"
	"math"
)

type RootishArrayStack struct {
	n      int
	blocks []ArrayStack
}

func (a *RootishArrayStack) Size() int {
	return a.n
}

func (a *RootishArrayStack) Init() {
	a.n = 0
}

func (a *RootishArrayStack) i2b(i int) int {
	return int(math.Ceil((-3.0 + math.Sqrt(9+8*float64(i)))) / 2)
}

func (a *RootishArrayStack) Get(i int) interface{} {
	b := a.i2b(i)
	j := i - b*(b+1)/2
	return a.blocks[b].Get(j)
}

func (a *RootishArrayStack) Set(i int, x interface{}) interface{} {
	b := a.i2b(i)
	j := i - b*(b+1)/2
	fmt.Printf("Block b=%d, j=%d\n", b, j)
	fmt.Printf("Block size is %d %d\n", a.n, len(a.blocks))
	fmt.Printf("Block j size=%d\n", a.blocks[b].Size())
	y := a.blocks[b].Get(j)

	fmt.Println("HERE :)")
	a.blocks[b].Set(j, x)
	return y
}

func (a *RootishArrayStack) Add(i int, x interface{}) {
	r := len(a.blocks)
	if r*(r+1)/2 < a.n+1 {
		a.grow()
	}
	a.n++
	for j := a.n; j > i+1; i-- {
		a.Set(j, a.Get(j-1))
	}
	a.Set(i, x)
}

func (a *RootishArrayStack) Remove(i int) interface{} {
	x := a.Get(i)
	for j := i; i < a.n-1; i++ {
		a.Set(j, a.Get(j+1))
	}
	a.n--
	r := len(a.blocks)
	if (r-2)*(r-1)/2 >= a.n {
		a.shrink()
	}
	return x
}

func (a *RootishArrayStack) grow() {
	new_array_stack := ArrayStack{}
	new_array_stack.Init(len(a.blocks) + 1)
	a.blocks = append(a.blocks, new_array_stack)
}

func (a *RootishArrayStack) shrink() {
	r := len(a.blocks)
	for r > 0 && (r-2)*(r-1)/2 >= a.n {
		a.blocks = a.blocks[:len(a.blocks)-1]
		r--
	}
}
