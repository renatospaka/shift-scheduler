package scheduler

import (
	"context"
	"database/sql"
	"log"

	httpServer "github.com/renatospaka/scheduler/adapter/chi"
	workerDomain "github.com/renatospaka/scheduler/scheduler/worker"
	"github.com/renatospaka/scheduler/adapter/event"
)

type SchedulerDomain struct {
	Ctx        context.Context
	DB         *sql.DB
	WebServer  *httpServer.HttpServer
	Dispatcher *event.EventDispatcher
	worker     *workerDomain.WorkerDomain
}

// start the scheduler core
func NewSchedulerDomain(ctx context.Context, db *sql.DB) *SchedulerDomain {
	log.Println("initiating scheduler domain application")
	s := &SchedulerDomain{
		Ctx: ctx,
		DB:  db,
	}

	// start event dispatcher
	s.Dispatcher = event.NewEventDispatcher()

	// start web server
	s.InitiateWebServer()
	return s
}

// Start web server adapter
func (s *SchedulerDomain) InitiateWebServer() {
	s.WebServer = httpServer.NewHttpServer(s.Ctx)

	s.worker = workerDomain.NewWorker(s.Ctx, s.DB, s.WebServer, s.Dispatcher)
	s.worker.StartWorkDomain()
}
