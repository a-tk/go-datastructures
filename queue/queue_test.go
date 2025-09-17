package queue

import "testing"

func TestEmptyQueue(t *testing.T) {
	q := New[int](5)
	if !q.Empty() {
		t.Errorf("new queue should be empty")
	}

	_, ok := q.Dequeue()
	if ok {
		t.Errorf("dequeue from empty queue should fail")
	}
}

func TestEnqueueDequeueBasic(t *testing.T) {
	q := New[int](5)

	ok := q.Enqueue(1)
	if !ok {
		t.Fatalf("enqueue should succeed")
	}
	ok = q.Enqueue(2)
	if !ok {
		t.Fatalf("enqueue should succeed")
	}

	v, ok := q.Dequeue()
	if !ok || v != 1 {
		t.Errorf("expected 1, got %v", v)
	}

	v, ok = q.Dequeue()
	if !ok || v != 2 {
		t.Errorf("expected 2, got %v", v)
	}

	if !q.Empty() {
		t.Errorf("queue should be empty after removing all elements")
	}
}

func TestQueueFull(t *testing.T) {
	q := New[int](2)

	if !q.Enqueue(10) || !q.Enqueue(20) {
		t.Fatalf("should be able to enqueue two elements")
	}

	if q.Enqueue(30) {
		t.Errorf("enqueue should fail when queue is full")
	}

	v, _ := q.Dequeue()
	if v != 10 {
		t.Errorf("expected 10, got %v", v)
	}

	if !q.Enqueue(30) {
		t.Errorf("enqueue should succeed after dequeue")
	}
}

func TestWrapAround(t *testing.T) {
	q := New[int](2)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Dequeue()  // remove 1
	q.Enqueue(3) // should wrap around internally

	v, _ := q.Dequeue()
	if v != 2 {
		t.Errorf("expected 2, got %v", v)
	}

	v, _ = q.Dequeue()
	if v != 3 {
		t.Errorf("expected 3, got %v", v)
	}

	if !q.Empty() {
		t.Errorf("queue should be empty after wraparound dequeues")
	}
}

func TestGenericStrings(t *testing.T) {
	q := New[string](3)
	q.Enqueue("a")
	q.Enqueue("b")

	v, _ := q.Dequeue()
	if v != "a" {
		t.Errorf("expected 'a', got %v", v)
	}
	v, _ = q.Dequeue()
	if v != "b" {
		t.Errorf("expected 'b', got %v", v)
	}
}
