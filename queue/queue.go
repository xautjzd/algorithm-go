package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	mu   sync.Mutex
	data []interface{}
}

func New(args ...int) *Queue {
	defaultSize := 1 << 4
	if len(args) > 0 {
		defaultSize = args[0]
	}

	return &Queue{
		data: make([]interface{}, 0, defaultSize),
	}
}

func (q *Queue) EnQueue(item interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.data = append(q.data, item)
}

func (q *Queue) DeQueue() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.IsEmpty() {
		return nil
	}

	item := q.data[0]
	q.data = q.data[1:]

	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Size() int {
	return len(q.data)
}

func main() {
	queue := New()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)

	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
}
