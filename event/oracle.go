package event

import (
	"math/rand"
	"time"
)

type Oracle struct {
	eq *EventQueue
}

func NewOracle() *Oracle {
	return &Oracle{NewEventQueue()}
}

func (o *Oracle) EventGenerator(nodeSize int, eventSize int) {
	// 500 millisecond triggers an event
	for {
		nodeID := rand.Int() % nodeSize
		eventType := rand.Int()%eventSize + 1
		o.eq.PushBack(&Event{nodeID, 1 << uint(eventType)})
		time.Sleep(100 * time.Millisecond)
	}
}

func (o *Oracle) FetchEvent() *Event {
	e := o.eq.Front()
	if e != nil {
		o.eq.Pop()
	}
	return e
}
