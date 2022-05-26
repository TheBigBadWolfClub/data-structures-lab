package pkg

import "sync"

// Queue is an abstract data structure. A queue is open at both its ends.
// One end is always used to insert data (enqueue)
// and the other is used to remove data (dequeue).
//
// Queue follows First-In-First-Out methodology,
// i.e., the data item stored first will be accessed first.
type Queue[T any] interface {
	Enqueue(T)
	Dequeue() (T, bool)

	Peek() (T, bool)
	IsFull() bool
	IsEmpty() bool
	Size() int
}

// NewQueue create a concurrent safe instance of Queue
// whose elements are of type T
func NewQueue[T any]() Queue[T] {
	return &queue[T]{}
}

type queue[T any] struct {
	mutex sync.RWMutex
	mem   []T
}

// Enqueue add (store) an item to the queue.
func (q *queue[T]) Enqueue(elem T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.mem = append(q.mem, elem)
}

// Dequeue remove (access) an item from the queue.
func (q *queue[T]) Dequeue() (elem T, ok bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	tmp := q.mem[0]
	q.mem = q.mem[1:]
	return tmp, true
}

// Peek Gets the element at the front of the queue without removing it.
func (q *queue[T]) Peek() (elem T, ok bool) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	return q.mem[0], true
}

// IsFull Checks if the queue is full.
// no queue limit yet ... always returns false
func (q *queue[T]) IsFull() bool {
	return false
}

// IsEmpty Checks if the queue is empty.
func (q *queue[T]) IsEmpty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return len(q.mem) == 0
}

// Size return the current number of elements
func (q *queue[T]) Size() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return len(q.mem)
}
