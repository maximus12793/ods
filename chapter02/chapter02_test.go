package ods

import (
	"math"
	"testing"
)

func TestArrayStack(t *testing.T) {
	stack := ArrayStack{}
	stack.Init(1)

	if stack.Size() != 0 {
		t.Errorf("stack.Size(): expected %d, got %d", 0, stack.Size())
	}

	stack.Add(0, "hello")
	stack.Add(1, true)
	stack.Add(2, 123)

	if stack.Size() != 3 {
		t.Errorf("stack.Size(): expected %d, got %d", 3, stack.Size())
	}

	if v := stack.Get(0); v != "hello" {
		t.Errorf("stack.Get(0): expected %q, got %v", "hello", v)
	}

	if v := stack.Get(1); v != true {
		t.Errorf("stack.Get(1): expected %t, got %v", true, v)
	}

	if v := stack.Get(2); v != 123 {
		t.Errorf("stack.Get(2): expected %d, got %v", 123, v)
	}

	if v := stack.Set(1, false); v != true {
		t.Errorf("stack.Set(1, false): expected %t, got %v", true, v)
	}

	if v := stack.Get(1); v != false {
		t.Errorf("stack.Get(1): expected %t, got %v", false, v)
	}

	if v := stack.Remove(0); v != "hello" {
		t.Errorf("stack.Remove(0): expected %q, got %v", "hello", v)
	}

	if stack.Size() != 2 {
		t.Errorf("stack.Size(): expected %d, got %d", 2, stack.Size())
	}

	if v := stack.Remove(1); v != 123 {
		t.Errorf("stack.Remove(1): expected %d, got %v", 123, v)
	}

	if stack.Size() != 1 {
		t.Errorf("stack.Size(): expected %d, got %d", 1, stack.Size())
	}

	if v := stack.Remove(0); v != false {
		t.Errorf("stack.Remove(0): expected %t, got %v", false, v)
	}

	if stack.Size() != 0 {
		t.Errorf("stack.Size(): expected %d, got %d", 0, stack.Size())
	}
}

func TestArrayQueue(t *testing.T) {
	queue := ArrayQueue{}
	queue.Init(1)

	// Test size
	if queue.Size() != 0 {
		t.Error("Expected size 0, got", queue.Size())
	}

	// Test add
	queue.Add(0, 123)
	if queue.Size() != 1 {
		t.Error("Expected size 1, got", queue.Size())
	}

	// Test remove
	if v := queue.Remove(); v != 123 {
		t.Error("Expected 123, got", v)
	}
	if queue.Size() != 0 {
		t.Error("Expected size 0, got", queue.Size())
	}
}

func TestArrayDeque(t *testing.T) {
	d := ArrayDeque{}
	d.Init()

	// Test Size
	if d.Size() != 0 {
		t.Errorf("Expected deque to be empty, but got size %d", d.Size())
	}

	// Test Add and Get
	d.Add(0, 1)
	d.Add(1, 2)
	d.Add(2, 3)
	if d.Size() != 3 {
		t.Errorf("Expected deque to have size 3, but got %d", d.Size())
	}
	if d.Get(0) != 1 {
		t.Errorf("Expected element at index 0 to be 1, but got %v", d.Get(0))
	}
	if d.Get(1) != 2 {
		t.Errorf("Expected element at index 1 to be 2, but got %v", d.Get(1))
	}
	if d.Get(2) != 3 {
		t.Errorf("Expected element at index 2 to be 3, but got %v", d.Get(2))
	}

	// Test Set
	d.Set(1, 4)
	if d.Get(1) != 4 {
		t.Errorf("Expected element at index 1 to be 4, but got %v", d.Get(1))
	}

	// Test Remove
	d.Remove(1)
	if d.Size() != 2 {
		t.Errorf("Expected deque to have size 2, but got %d", d.Size())
	}
	if d.Get(1) != 3 {
		t.Errorf("Expected element at index 1 to be 3, but got %v", d.Get(1))
	}
}

func TestArrayDequeResize(t *testing.T) {
	d := ArrayDeque{}
	d.Init()

	// Add elements until the deque becomes full
	for i := 0; i < 10; i++ {
		d.Add(i, i)
	}

	// Check that the deque is full
	if d.Size() != 10 {
		t.Errorf("Expected deque to have size 10, but got %d", d.Size())
	}

	// Add one more element to trigger a resize
	d.Add(10, 10)

	// Check that the deque was resized
	if d.Size() != 11 {
		t.Errorf("Expected deque to have size 11, but got %d", d.Size())
	}
}

func TestDualArrayDeque(t *testing.T) {
	d := DualArrayDeque{}
	d.Init()

	// Test Size method
	if d.Size() != 0 {
		t.Errorf("Expected size 0, got %d", d.Size())
	}

	// Test Add and Get methods
	d.Add(0, "a")
	d.Add(1, "b")
	d.Add(2, "c")
	if d.Get(0) != "a" {
		t.Errorf("Expected element at index 0 to be 'a', got %v", d.Get(0))
	}
	if d.Get(1) != "b" {
		t.Errorf("Expected element at index 1 to be 'b', got %v", d.Get(1))
	}
	if d.Get(2) != "c" {
		t.Errorf("Expected element at index 2 to be 'c', got %v", d.Get(2))
	}

	// Test Set method
	d.Set(1, "d")
	if d.Get(1) != "d" {
		t.Errorf("Expected element at index 1 to be 'd', got %v", d.Get(1))
	}

	// Test Remove method
	d.Remove(1)
	if d.Get(1) != "c" {
		t.Errorf("Expected element at index 1 to be 'c', got %v", d.Get(1))
	}
}

func TestDualArrayDequeAdvanced(t *testing.T) {
	d := DualArrayDeque{}
	d.Init()

	// Test balance method
	for i := 0; i < 1500; i++ {
		d.Add(i, i)
	}
	if math.Abs(float64(d.front.Size()/d.back.Size())) > 3 {
		t.Errorf("Expected front and back sizes to be roughly equal, got front size %d and back size %d", d.front.Size(), d.back.Size())
	}

	// Test edge cases of Add, Get, and Set methods
	d.Add(0, "a")
	d.Add(d.Size(), "b")
	if d.Get(0) != "a" {
		t.Errorf("Expected element at index 0 to be 'a', got %v", d.Get(0))
	}
	if d.Get(d.Size()-1) != "b" {
		t.Errorf("Expected element at index %d to be 'b', got %v", d.Size()-1, d.Get(d.Size()-1))
	}
	d.Set(0, "c")
	d.Set(d.Size()-1, "d")
	if d.Get(0) != "c" {
		t.Errorf("Expected element at index 0 to be 'c', got %v", d.Get(0))
	}
	if d.Get(d.Size()-1) != "d" {
		t.Errorf("Expected element at index %d to be 'd', got %v", d.Size()-1, d.Get(d.Size()-1))
	}
}

func TestRootishArrayStack(t *testing.T) {
	// Test empty RootishArrayStack
	ras := RootishArrayStack{}
	ras.Init()
	if ras.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ras.Size())
	}

	// Test adding and getting elements
	ras.Add(0, "a")
	if ras.Size() != 1 {
		t.Errorf("Expected size 1, got %d", ras.Size())
	}
	if ras.Get(0) != "a" {
		t.Errorf("Expected element 'a', got %v", ras.Get(0))
	}

	// Test setting elements
	ras.Set(0, "b")
	if ras.Get(0) != "b" {
		t.Errorf("Expected element 'b', got %v", ras.Get(0))
	}

	// Test removing elements
	if ras.Remove(0) != "b" {
		t.Errorf("Expected element 'b', got %v", ras.Remove(0))
	}
	if ras.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ras.Size())
	}
}

func TestRootishArrayStackVariant(t *testing.T) {
	// Test empty RootishArrayStack
	ras := RootishArrayStack{}
	ras.Init()
	if ras.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ras.Size())
	}
	str := "abcdefgh"
	for i, elem := range str {
		ras.Add(i, elem)
	}

	ras.Add(2, 'x')
	ras.Remove(1)
	ras.Remove(7)
	ras.Remove(6)

	str = "axcdef"
	for i, elem := range str {
		if ras.Get(i) != elem {
			t.Errorf("Expected element '%c', got %c", elem, ras.Get(i))
		}
	}

	// Test setting elements
	ras.Set(0, 'b')
	if ras.Get(0) != 'b' {
		t.Errorf("Expected element 'b', got %v", ras.Get(0))
	}
	// Test removing elements
	if ras.Remove(0) != 'b' {
		t.Errorf("Expected element 'b', got %v", ras.Remove(0))
	}
}
