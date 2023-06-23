package test

type testEvent struct {
	data any
}

func (e *testEvent) GetKey() string {
	return "event 1"
}

func (e *testEvent) GetData() interface{} {
	e.data = "event 1"
	return e.data
}
