package document

import (
	"context"
	"database/sql"
	"log"

	httpServer "github.com/renatospaka/scheduler/adapter/chi"
	repository "github.com/renatospaka/scheduler/scheduler/document/adapter/postgres"
	"github.com/renatospaka/scheduler/scheduler/document/adapter/web/controller"
	"github.com/renatospaka/scheduler/scheduler/document/adapter/web/route"
	"github.com/renatospaka/scheduler/scheduler/document/core/usecase"
)

type DocumentDomain struct {
	repo        *repository.DocumentRepository
	usecases    *usecase.DocumentUsecase
	controllers *controller.DocumentController
	routes      *route.DocumentRoute
	*domain
}

type domain struct {
	ctx        context.Context
	db         *sql.DB
	webServer  *httpServer.HttpServer
	// dispatcher *event.EventDispatcher
}

// Initiate the Document Domain
func NewDocument(ctx context.Context, db *sql.DB, webServer *httpServer.HttpServer) *DocumentDomain {
	return &DocumentDomain{
		domain: &domain{
			ctx:        ctx,
			db:         db,
			webServer:  webServer,
			// dispatcher: dispatcher,
		},
	}
}

// Set up all necessary structs to start the Document Domain
func (w *DocumentDomain) StartDocumentDomain() error {
	log.Println("--------------------------------------")
	log.Println("initiating document domain application")

	w.repo = repository.NewDocumentRepository(w.domain.db)
	w.usecases = usecase.NewDocumentUsecase(w.repo)
	w.usecases.SetContext(w.ctx)
	w.controllers = controller.NewDocumentController(w.usecases)
	w.routes = route.NewDocumentRoute(w.domain.webServer, w.controllers)
	w.routes.Routes()

	return nil
}
