package test

type testListener struct {
	data   interface{}
	called bool
}

func (l *testListener) Handle() error {
	l.called = true
	return nil
}

func (l *testListener) SetData(data interface{}) {
	l.data = data
}
