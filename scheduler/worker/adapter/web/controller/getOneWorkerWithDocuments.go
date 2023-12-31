package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	pkgController "github.com/renatospaka/scheduler/adapter/web/controller"
	"github.com/renatospaka/scheduler/scheduler/worker/core/usecase"
)

func (c *WorkerController) getOneWorkerWithDocuments(w http.ResponseWriter, r *http.Request) {
	log.Println("scheduler.worker.adapter.web.controller.GetWithDocuments")

	var (
		in  usecase.GetWorkerWithDocumentsByIdInputDto
		out usecase.GetWorkerWithDocumentsByIdOutputDto
	)

	param := chi.URLParam(r, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		out.StandardStatusOutputDto = pkgController.FormatStatus(pkgController.REQUEST_FAILURE, http.StatusBadRequest, "worker id must be a number")
		w.WriteHeader(out.Error.Code)
		json.NewEncoder(w).Encode(out)
		return
	}

	if id <= 0 {
		out.StandardStatusOutputDto = pkgController.FormatStatus(pkgController.REQUEST_FAILURE, http.StatusBadRequest, "invalid worker id")
		out.ID = id
		w.WriteHeader(out.Error.Code)
		json.NewEncoder(w).Encode(out)
		return
	}

	// Finally gets the worker by id with all documents attached
	in = usecase.GetWorkerWithDocumentsByIdInputDto{ID: id}
	worker, err := c.usecases.GetWorkerWithDocumentsById(in)
	if err != nil {
		out.StandardStatusOutputDto = pkgController.FormatStatus(pkgController.REQUEST_FAILURE, http.StatusInternalServerError, fmt.Sprintf("error: %s", err.Error()))
		w.WriteHeader(out.Error.Code)
		json.NewEncoder(w).Encode(out)
		return
	}

	if worker.Status == pkgController.REQUEST_FAILURE {
		w.WriteHeader(worker.Error.Code)
		json.NewEncoder(w).Encode(worker.Error.Message)
		return
	}

	// Everything went well and the worker was found
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(worker)
}
