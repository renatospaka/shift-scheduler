package event

type WorkerGottenEvent struct {
	data interface{}
}

func NewWorkerGottenEvent(data string) *WorkerGottenEvent {
	return &WorkerGottenEvent{
		data: data,
	}
}

func (e *WorkerGottenEvent) GetKey() string {
	return EVENT_WORKER_GOTTEN
}

func (e *WorkerGottenEvent) GetData() interface{} {
	return e.data
}
