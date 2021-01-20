package queue

import (
	"log"
	"testing"
)

func TestNewLKQueue(t *testing.T) {
	q := NewLKQueue()
	q.Enqueue(1)
	log.Println(q.Dequeue())
}
