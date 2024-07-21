// package queue implements an regular queue data structure with safe concurrency access
package queue

import (
	"errors"
	"sync"
)

var (
    ErrEmptyQueue = errors.New("queue is empty!")
)

// Queue data structure
type Queue[T any] struct {
    mutex sync.Mutex
    data []T
}

// Return if the queue is empty
func (q *Queue[T]) Empty() bool {
    return len(q.data) == 0
}

// Put a new element in the back of the queue
func (q *Queue[T]) Enqueue(value T) {
    q.mutex.Lock()
    defer q.mutex.Unlock()

    q.data = append(q.data, value)
}

// Remove and return the element in the front or error if is empty
func (q *Queue[T]) Dequeue() (T, error) {
    q.mutex.Lock()
    defer q.mutex.Unlock()
    var v T
    if q.Empty() {
        return v, ErrEmptyQueue
    }
    v = q.data[0]
    q.data = q.data[1:]
    return v, nil
}

// Return the element in front of the queue but doenst remove
func (q *Queue[T]) Peek() (T, error) {
    q.mutex.Lock()
    defer q.mutex.Unlock()
    var v T
    if q.Empty() {
        return v, ErrEmptyQueue
    }
    v = q.data[0]
    return v, nil
}
