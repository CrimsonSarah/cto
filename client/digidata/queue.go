package digidata

// TODO: Is there a way to specify this as a type parameter alongside T?
const QUEUE_SIZE byte = 5

// Fixed-size using a circular buffer.
// Keeps track of size and drops Enqueues after it is full.
type Queue[T any] struct {
	Items     [QUEUE_SIZE]T
	HeadIndex byte
	TailIndex byte
	Size      byte
}

func MakeQueue[T any]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	if q.Size >= QUEUE_SIZE {
		return
	}

	q.Items[q.TailIndex] = value

	nextIndex := (q.TailIndex + 1) % QUEUE_SIZE
	q.TailIndex = nextIndex
	q.Size += 1
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.Size <= 0 {
		var zero T
		return zero, false
	}

	currentValue := q.Items[q.HeadIndex]

	nextIndex := (q.HeadIndex + 1) % QUEUE_SIZE
	q.HeadIndex = nextIndex
	q.Size -= 1

	return currentValue, true
}
