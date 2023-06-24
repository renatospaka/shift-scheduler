package event

import "log"

type workerGottenListener struct {
	data interface{}
}

func NewWorkerGottenListener() *workerGottenListener {
	return &workerGottenListener{}
}

func (l *workerGottenListener) SetData(data interface{}) {
	l.data = data
}

func (l *workerGottenListener) Handle() error {
	log.Println("...event related to worker fetching")
	return nil
}
