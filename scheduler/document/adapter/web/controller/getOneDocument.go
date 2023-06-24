package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	pkgController "github.com/renatospaka/scheduler/adapter/web/controller"
	"github.com/renatospaka/scheduler/scheduler/document/core/usecase"
)

var (
	in  usecase.GetDocumentByIdInputDto
	out usecase.GetDocumentByIdOutputDto
)

func (c *DocumentController) getOneDocument(w http.ResponseWriter, r *http.Request) {
	log.Println("scheduler.document.adapter.web.controller.Get")

	param := chi.URLParam(r, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		out.StandardStatusOutputDto = pkgController.FormatStatus(pkgController.REQUEST_FAILURE, http.StatusBadRequest, "document id must be a number")
		w.WriteHeader(out.Error.Code)
		json.NewEncoder(w).Encode(out)
		return
	}

	if id <= 0 {
		out.StandardStatusOutputDto = pkgController.FormatStatus(pkgController.REQUEST_FAILURE, http.StatusBadRequest, "invalid document id")
		out.ID = id
		w.WriteHeader(out.Error.Code)
		json.NewEncoder(w).Encode(out)
		return
	}

	// Finally gets the document by id
	in = usecase.GetDocumentByIdInputDto{ID: id}
	docto, err := c.usecases.GetDocumentById(in)
	if err != nil {
		out.StandardStatusOutputDto = pkgController.FormatStatus(pkgController.REQUEST_FAILURE, http.StatusInternalServerError, fmt.Sprintf("error: %s", err.Error()))
		w.WriteHeader(out.Error.Code)
		json.NewEncoder(w).Encode(out)
		return
	}

	if docto.Status == pkgController.REQUEST_FAILURE {
		w.WriteHeader(docto.Error.Code)
		json.NewEncoder(w).Encode(docto.Error.Message)
		return
	}

	// Everything went well and the document was found
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(docto)
}
