package queue

import "testing"

func Test_IsEmpty(t *testing.T) {
	q := New()

	if !q.IsEmpty() {
		t.Errorf("Queue isn't empty")
	}
}

func Test_Enqueue(t *testing.T) {
	q := New()

	q.Enqueue("my element")

	if q.IsEmpty() {
		t.Errorf("Queue isn't empty")
	}
}

func Test_Dequeue(t *testing.T) {
	q := New()

	q.Enqueue("my element")
	q.Enqueue("other element")
	v1, ok1 := q.Dequeue()
	v2, ok2 := q.Dequeue()

	if !ok1 {
		t.Errorf("Unable to dequeue")
	}

	if !ok2 {
		t.Errorf("Unable to dequeue")
	}

	if v1 != "my element" {
		t.Errorf("value isn't equal to my element")
	}

	if v2 != "other element" {
		t.Errorf("value isn't equal to my element")
	}
}
