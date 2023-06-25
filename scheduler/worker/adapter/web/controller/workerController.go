package controller

import (
	"log"
	"net/http"

	"github.com/renatospaka/scheduler/scheduler/worker/core/usecase"
)

type WorkerController struct {
	usecases *usecase.WorkerUsecase
}

func NewWorkerController(usecases *usecase.WorkerUsecase) *WorkerController {
	log.Println("initiating worker controllers")

	return &WorkerController{
		usecases: usecases,
	}
}


// Gets only the Worker identified by the given id
func (c *WorkerController) Get(w http.ResponseWriter, r *http.Request) {	
	c.getOneWorker(w, r)
}

// Gets only the Worker identified by the given id with all related documents
func (c *WorkerController) GetWithDocuments(w http.ResponseWriter, r *http.Request) {	
	c.getOneWorkerWithDocuments(w, r)
}
