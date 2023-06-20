package worker

import (
	"context"
	"database/sql"
	"log"

	httpServer "github.com/renatospaka/scheduler/adapter/chi"
	repository "github.com/renatospaka/scheduler/scheduler/worker/adapter/postgres"
	"github.com/renatospaka/scheduler/scheduler/worker/adapter/web/controller"
	"github.com/renatospaka/scheduler/scheduler/worker/adapter/web/route"
	"github.com/renatospaka/scheduler/scheduler/worker/core/usecase"
)

type WorkerDomain struct {
	repo        *repository.WorkerRepository
	usecases    *usecase.WorkerUsecase
	controllers *controller.WorkerController
	routes      *route.WorkerRoute
	*domain
}

type domain struct {
	ctx       context.Context
	db        *sql.DB
	webServer *httpServer.HttpServer
}


// Initiate the Worker Domain
func NewWorker(ctx context.Context, db *sql.DB, WebServer *httpServer.HttpServer) *WorkerDomain {
	return &WorkerDomain{
		domain: &domain{
			ctx:       ctx,
			db:        db,
			webServer: WebServer,
		},
	}
}


// Set up all necessary structs to start the Worker Domain
func (w *WorkerDomain) StartWorkDomain() error {
	log.Println("initiating worker domain application")

	w.repo = repository.NewWorkerRepository(w.domain.db)
	w.usecases = usecase.NewWorkerUsecase(w.repo)
	w.controllers = controller.NewWorkerController(w.usecases)
	w.routes = route.NewWorkerRoute(w.domain.webServer, w.controllers)
	w.routes.Routes()

	return nil
}
