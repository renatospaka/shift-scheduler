package event

type WorkerGottenEvent struct {
	data interface{}
}

// Start an instance of the event
func NewWorkerGottenEvent(data string) *WorkerGottenEvent {
	return &WorkerGottenEvent{
		data: data,
	}
}

// Register the Key
func (e *WorkerGottenEvent) GetKey() string {
	return EVENT_WORKER_GOTTEN
}

// Get the data of the raised event
func (e *WorkerGottenEvent) GetData() interface{} {
	return e.data
}
