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

func (c *WorkerController) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("scheduler.worker.adapter.web.controller.Get")
	
	panic("implement me")
}
