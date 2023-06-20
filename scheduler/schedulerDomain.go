package scheduler

import (
	"context"
	"database/sql"
	"log"

	httpServer "github.com/renatospaka/scheduler/adapter/chi"
	workerDomain "github.com/renatospaka/scheduler/scheduler/worker"
)

type SchedulerDomain struct {
	Ctx        context.Context
	DB         *sql.DB
	WebServer *httpServer.HttpServer
	worker *workerDomain.WorkerDomain
}

// start the scheduler core
func NewSchedulerDomain(ctx context.Context, db *sql.DB) *SchedulerDomain {
	log.Println("initiating scheduler domain application")
	s := &SchedulerDomain{
		Ctx:       ctx,
		DB:        db,
	}

	s.InitiateHttp()
	return s
}


// Start web seerver adapter
func (s *SchedulerDomain) InitiateHttp() {
	s.WebServer = httpServer.NewHttpServer(s.Ctx)

	s.worker = workerDomain.NewWorker(s.Ctx, s.DB, s.WebServer)
	s.worker.StartWorkDomain()
}

