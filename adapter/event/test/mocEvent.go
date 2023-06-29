package test

// Mocked Event adapter
// NOT TO BE CALLED under no circumstances out of the testing process
type testEvent struct {
	data any
}

// Mocked method GetKey of an event
// NOT TO BE CALLED under no circumstances out of the testing process
func (e *testEvent) GetKey() string {
	return "event 1"
}

// Mocked method GetData of an event
// NOT TO BE CALLED under no circumstances out of the testing process
func (e *testEvent) GetData() interface{} {
	e.data = "event 1"
	return e.data
}
