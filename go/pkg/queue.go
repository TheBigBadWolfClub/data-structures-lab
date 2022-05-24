package pkg

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() (T, bool)

	Peek() (T, bool)
	IsFull() bool
	IsEmpty() bool
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{}
}

type queue[T any] struct {
	mem []T
}

//Enqueue add (store) an item to the queue.
func (q *queue[T]) Enqueue(elem T) {
	q.mem = append(q.mem, elem)
}

//Dequeue remove (access) an item from the queue.
func (q *queue[T]) Dequeue() (elem T, ok bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	tmp := q.mem[0]
	q.mem = q.mem[1:]
	return tmp, true
}

//Peek Gets the element at the front of the queue without removing it.
func (q *queue[T]) Peek() (elem T, ok bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	return q.mem[0], true
}

//IsFull Checks if the queue is full.
//no queue limit yet ... always returns false
func (q *queue[T]) IsFull() bool {
	return false
}

//IsEmpty Checks if the queue is empty.
func (q *queue[T]) IsEmpty() bool {
	return len(q.mem) == 0
}
