package event

import "container/list"

type EventQueue struct {
	queue *list.List
}

func NewEventQueue() *EventQueue {
	return &EventQueue{list.New()}
}

func (q *EventQueue) Length() int {
	return q.queue.Len()
}

func (q *EventQueue) Front() *Event {
	e := q.queue.Front()
	if e != nil {
		return q.queue.Front().Value.(*Event)
	} else {
		return nil
	}
}

func (q *EventQueue) Pop() {
	q.queue.Remove(q.queue.Front())
}

func (q *EventQueue) PushBack(e *Event) {
	q.queue.PushBack(e)
}
