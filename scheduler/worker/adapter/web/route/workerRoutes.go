package route

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	httpServer "github.com/renatospaka/scheduler/adapter/chi"
	"github.com/renatospaka/scheduler/scheduler/worker/adapter/web/controller"
)

type WorkerRoute struct {
	server      *httpServer.HttpServer
	controllers *controller.WorkerController
}

func NewWorkerRoute(server *httpServer.HttpServer, controllers *controller.WorkerController) *WorkerRoute {
	log.Println("initiating worker routes")

	return &WorkerRoute{
		server:      server,
		controllers: controllers,
	}
}

func (w *WorkerRoute) Routes() {
	log.Println("scheduler.worker.adapter.web.route.Routes")

	w.server.Server.Route("/workers", func(r chi.Router) {
		// r.Use(jwtauth.Verifier(configs.TokenAuth))
		// r.Use(jwtauth.Authenticator)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Workers"))
		})
		r.Get("/{id}", w.controllers.Get)
	})
}
