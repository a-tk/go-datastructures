package queue

import (
	"fmt"
	"strings"
)

// Queue is a fast but size restricted queue using a fixed array
type Queue[E any] struct {
	a               []E
	head, tail, cap int
}

func New[E any](cap int) *Queue[E] {
	return &Queue[E]{
		a:    make([]E, cap+1),
		head: 0,
		tail: 0,
		cap:  cap + 1, // CLRS provides algos for n-1 elements, so make the extra space
	}
}

func (q *Queue[E]) Empty() bool {
	return q.head == q.tail
}

func (q *Queue[E]) Enqueue(e E) (ok bool) {
	if q.head == ((q.tail + 1) % q.cap) {
		return false
	} else {
		q.a[q.tail] = e
		if q.tail == q.cap-1 {
			q.tail = 0
		} else {
			q.tail++
		}
		return true
	}
}
func (q *Queue[E]) Dequeue() (e E, ok bool) {
	if q.Empty() {
		return e, false
	} else {
		e = q.a[q.head]
		if q.head == q.cap-1 {
			q.head = 0
		} else {
			q.head++
		}
		return e, true
	}
}

func (q *Queue[E]) String() string {
	var b strings.Builder

	for i := 0; i < q.cap; i++ {
		var err error
		if i == q.head {
			// green
			_, err = fmt.Fprintf(&b, "\x1b[42m\u001B[4m[ %v ]\x1b[0m ", q.a[i])
		} else if i == q.tail {
			// red
			_, err = fmt.Fprintf(&b, "\x1b[41m\u001B[4m[ %v ]\x1b[0m ", q.a[i])
		} else {
			_, err = fmt.Fprintf(&b, "[ %v ] ", q.a[i])
		}
		if err != nil {
			return "error building string"
		}
	}

	return b.String()
}
