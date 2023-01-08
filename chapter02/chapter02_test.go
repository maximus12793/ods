package ods

import "testing"

func TestArrayStack(t *testing.T) {
	stack := ArrayStack{}
	stack.Init()

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
	queue.Init()

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
