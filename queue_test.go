package queue

import (
	"math/rand"
	"testing"
)

func TestEnqueueDequeue(t *testing.T) {
	for tries := 0; tries < 1000; tries++ {
		times := 1000
		var q Queue[int]
		valuesToPut := make([]int, times)
		for i := 0; i < times; i++ {
			valuesToPut[i] = rand.Intn(1000)
			q.Enqueue(valuesToPut[i])
		}
		for i := 0; i < times; i++ {
			if v, err := q.Dequeue(); err == nil {
				if v != valuesToPut[i] {
					t.Fatalf("expected %d got %d!", valuesToPut[i], v)
				}
			} else {
				t.Fatal("queue was not expected to be empty!")
			}
		}
	}
}

func TestQueueEmpty(t *testing.T) {
	for tries := 0; tries < 1000; tries++ {
		q := Queue[int]{}
		times := rand.Intn(10000)
		for i := 0; i < times; i++ {
			q.Enqueue(rand.Int())
			if q.Empty() {
				t.Fatal("queue was expected to be not empty here!")
			}
		}
		for i := 0; i < times; i++ {
			if q.Empty() {
				t.Fatal("queue was expected to be not empty here!")
			}
			if _, err := q.Dequeue(); err != nil {
				t.Fatal("error value different of what expected!")
			}
		}
		if !q.Empty() {
			t.Fatal("queue was expected to be empty here!")
		}
	}
}

func TestPeek(t *testing.T) {
	for tries := 0; tries < 1000; tries++ {
        var q Queue[int]
        value := rand.Int()
        q.Enqueue(value)
        if peekedValue, err := q.Peek(); err == nil {
            if peekedValue != value {
                t.Fatalf("expected %d got %d!", value, peekedValue)
            }
        } else {
            t.Fatal("queue was not expected to be empty!")
        }
	}
}
