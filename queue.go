package queue

type queue struct {
	elements []interface{}
}

// New create a new quey system
func New() *queue {
	return &queue{}
}

// IsEmpty check if queue is empty
func (q *queue) IsEmpty() bool {
	return len(q.elements) <= 0
}

// Enqueue put new element into queue
func (q *queue) Enqueue(e interface{}) {
	q.elements = append(q.elements, e)
}

// Dequeue remove element in from of queue
func (q *queue) Dequeue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	e := q.elements[0]
	q.elements = q.elements[1:]
	return e, true
}
