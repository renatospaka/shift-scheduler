package event

import "log"

type workerGottenListener struct {
	data interface{}
}

// Start the listener of the WORKER_GOTTEN event
func NewWorkerGottenListener() *workerGottenListener {
	return &workerGottenListener{}
}

// Update the data to be handled by the event
func (l *workerGottenListener) SetData(data interface{}) {
	l.data = data
}

// Trigger and process the WORKER_GOTTEN event
func (l *workerGottenListener) Handle() error {
	log.Println("...event related to worker fetching")
	return nil
}
