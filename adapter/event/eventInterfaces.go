package event

// Handles Listener data and execution
type Listener interface {
	SetData(data interface{})
	Handle() error
}


// Holds the information of the Event itself
type Event interface {
	GetKey() string
	GetData() interface{}
}
