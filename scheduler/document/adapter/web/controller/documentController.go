package controller

import (
	"log"
	"net/http"

	"github.com/renatospaka/scheduler/scheduler/document/core/usecase"
)


type DocumentController struct {
	usecases *usecase.DocumentUsecase
}

func NewDocumentController(usecases *usecase.DocumentUsecase) *DocumentController {
	log.Println("initiating document controllers")

	return &DocumentController{
		usecases: usecases,
	}
}


// Gets only the Document identified by the given id
func (c *DocumentController) Get(w http.ResponseWriter, r *http.Request) {	
	c.getOneDocument(w, r)
}
