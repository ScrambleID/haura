package server

import "testing"

func TestQueue(t *testing.T) {
	q := newQueue()
	q.Enqueue("hello")
	if q.Dequeue() != "hello" {
		t.Errorf("expected %v", "hello")
	}
}
