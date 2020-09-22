package event

import (
	"context"
	"time"
)

var EventStore = eventStore{store: make(map[string]interface{}), timeout: 5 * time.Second}

type (
	Event struct {
		Data interface{}
		*Observer
	}
	Observable func(*Observer) error

	//Operator func(Observable) Observable

	NextHandler interface {
		OnNext(*Event)
	}
	NextFunc func(*Event)

	//Operator   func(Observable) Observable
	//
	//NextCancel func()
	//NextChan   chan *Event

	Observer struct {
		context.Context
		ctxFunc context.CancelFunc
		next    NextHandler
	}

	eventStore struct {
		store   map[string]interface{}
		timeout time.Duration
	}
)

// NExFunc
func (next NextFunc) OnNext(event *Event) {
	next(event)
}
//func (next NextChan) OnNext(event *Event) {
//	next <- event
//}
//func (next NextCancel) OnNext(event *Event) {
//	next()
//}

// eventstore

func (set *eventStore) replayAtackCheck(eventId string) bool {
	_, ok := set.store[eventId]
	return ok
}

//------------
func (set *eventStore) add(eventId string) {
	set.store[eventId] = nil
}
func (set *eventStore) remove(eventId string) {
	delete(set.store, eventId)
}
func (set *eventStore) isEmpty() bool {
	return len(set.store) == 0
}

// observe

func (c *Observer) IsDisposed() bool {

	return c.Err() != nil
}

func (c *Observer) Disposed() {
	c.ctxFunc()
}

func (c *Observer) Next(data interface{}) {
	c.Push(&Event{Data: data, Observer: c})

}

func (c *Observer) Push(event *Event) {

	if !c.IsDisposed() {
		if event.Context != c {
			event = &Event{Data: event.Data, Observer: c}
		}

		c.next.OnNext(event)
	}
}

func (c *Observer) Create(next NextFunc) *Observer {
	return &Observer{c, c.ctxFunc, next}
}



