package ods

type DualArrayDeque struct {
	// Note: front stores elements in reverse order.
	front ArrayStack
	// Note: back stores elements in standard order.
	back ArrayStack
}

func (a *DualArrayDeque) Size() int {
	return a.front.Size() + a.back.Size()
}

func (a *DualArrayDeque) Init() {
	a.front = ArrayStack{}
	a.front.Init()

	a.back = ArrayStack{}
	a.back.Init()
}

func (a *DualArrayDeque) Get(i int) interface{} {
	if i < a.front.Size() {
		return a.front.Get(a.front.Size() - i - 1)
	} else {
		return a.back.Get(i - a.front.Size())
	}
}

func (a *DualArrayDeque) Set(i int, x interface{}) interface{} {
	if i < a.front.Size() {
		return a.front.Set(a.front.Size()-i-1, x)
	} else {
		return a.back.Set(i-a.front.Size(), x)
	}
}

func (a *DualArrayDeque) Add(i int, x interface{}) {
	if i < a.front.Size() {
		a.front.Add(a.front.Size()-i, x)
	} else {
		a.back.Add(i-a.front.Size(), x)
	}
	a.balance()
}

func (a *DualArrayDeque) Remove(i int) interface{} {
	var x interface{}

	if i < a.front.Size() {
		x = a.front.Remove(a.front.Size() - i - 1)
	} else {
		x = a.back.Remove(i - a.front.Size())
	}
	a.balance()
	return x
}

func (a *DualArrayDeque) balance() {
	if (3*a.front.Size() < a.back.Size()) || (3*a.back.Size() < a.front.Size()) {
		n := a.Size()
		mid := n / 2
		f := ArrayStack{}
		f.Init()
		for i := 0; i < mid; i++ {
			f.Add(i, a.Get(mid-i-1))
		}
		b := ArrayStack{}
		b.Init()
		for i := 0; i < n-mid; i++ {
			b.Add(i, a.Get(mid+i))
		}
		a.front = f
		a.back = b
	}
}
