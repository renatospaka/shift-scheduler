package test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/renatospaka/scheduler/adapter/event"
)

func TestEventDispatcherAddListener(t *testing.T) {
	disp := event.NewEventDispatcher()
	testListener := &testListener{}
	disp.AddListener("event 1", testListener)

	assert.Equal(t, 1, len(disp.Listeners["event 1"]))
	assert.Equal(t, testListener, disp.Listeners["event 1"][0])
}

func TestEventDispatcherDispatch(t *testing.T) {
	disp := event.NewEventDispatcher()
	testListener := &testListener{}
	disp.AddListener("event 1", testListener)

	event := &testEvent{}
	disp.Dispatch(event)
	assert.True(t, testListener.called)
}
