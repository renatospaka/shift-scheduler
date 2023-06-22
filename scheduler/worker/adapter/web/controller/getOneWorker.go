package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/renatospaka/scheduler/core/dto"
	"github.com/renatospaka/scheduler/core/errorList"
)

type getWorker struct {
	Procedure string `json:"procedure`
	Id        int    `json:"worker_id"`
}

func (c *WorkerController) getOneWorker(w http.ResponseWriter, r *http.Request) {
	log.Println("scheduler.worker.adapter.web.controller.Get")

	param := chi.URLParam(r, "id")
	log.Printf("param: %v\n", param)
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println("error converting to int")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	} else if id <= 0 {
		log.Println(errorList.ErrWorkerIdInvalid.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	worker, err := c.usecases.GetWorker(id)
	if err != nil {
		failures := dto.ErrorTraillerDto{}
		failure := dto.ErrorDto{
			Error: err.Error(),
		}
		failures.Errors = append(failures.Errors, failure)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(failures.Errors)
		return
	}

	if worker == nil {
		failures := dto.ErrorTraillerDto{}
		failure := dto.ErrorDto{
			Error: errorList.ErrWorkerIdNotFound.Error(),
		}
		failures.Errors = append(failures.Errors, failure)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(failures.Errors)
		return
	}
}
