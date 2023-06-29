package event

import "log"

type EventDispatcher struct {
	Listeners map[string][]Listener
}

func NewEventDispatcher() *EventDispatcher {
	log.Println("initiating events dispatcher")
	return &EventDispatcher{
		Listeners: make(map[string][]Listener),
	}
}

// Links an Event to a Listener
func (e *EventDispatcher) AddListener(event string, listener Listener) {
	if e.Listeners == nil {
		e.Listeners = make(map[string][]Listener)
	}
	e.Listeners[event] = append(e.Listeners[event], listener)
}

// Executes all triggered events at once
func (e *EventDispatcher) Dispatch(event Event) {
	if e.Listeners == nil {
		return
	}

	for _, listener := range e.Listeners[event.GetKey()] {
		listener.SetData(event.GetData())
		listener.Handle()
	}
}
