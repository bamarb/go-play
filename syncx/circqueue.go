package syncx

type CircQueue[T any] struct {
	elems     []T
	head, len int
}

func (q *CircQueue[T]) Cap() int {
	return len(q.elems)
}

func (q *CircQueue[T]) Enqueue(e T) {
	if q.len+1 > q.Cap() {
		q.grow()
	}
	i := (q.head + q.len) % q.Cap()
	q.elems[i] = e
	q.len++
}

func (q *CircQueue[T]) Dequeue() (T, bool) {
	if q.len == 0 {
		var zero T
		return zero, false
	}
	e := q.elems[q.head]
	// q.elems[q.head] = nil
	q.head = (q.head + 1) % q.Cap()
	q.len--
	return e, true
}

func (q *CircQueue[T]) peek() (any, bool) {
	if q.len == 0 {
		return nil, false
	}
	return q.elems[q.head], true
}

func (q *CircQueue[T]) clear() {
	*q = CircQueue[T]{}
}

func (q *CircQueue[T]) grow() {
	oldCap := q.Cap()
	newCap := oldCap * 2
	if newCap == 0 {
		newCap = 8
	}
	newElems := make([]T, newCap)
	oldLen := q.len
	for i := 0; i < oldLen; i++ {
		newElems[i] = q.elems[(q.head+i)%oldCap]
	}
	q.elems = newElems
	q.head = 0
}
