package route

import (
	"log"

	"github.com/go-chi/chi"
	httpServer "github.com/renatospaka/scheduler/adapter/chi"
	"github.com/renatospaka/scheduler/scheduler/document/adapter/web/controller"
)

type DocumentRoute struct {
	server      *httpServer.HttpServer
	controllers *controller.DocumentController
}

func NewDocumentRoute(server *httpServer.HttpServer, controllers *controller.DocumentController) *DocumentRoute {
	log.Println("initiating document routes")

	return &DocumentRoute{
		server:      server,
		controllers: controllers,
	}
}

func (d *DocumentRoute) Routes() {
	log.Println("scheduler.document.adapter.web.route.Routes")

	d.server.Server.Route("/documents", func(r chi.Router) {
		// r.Use(jwtauth.Verifier(configs.TokenAuth))
		// r.Use(jwtauth.Authenticator)

		r.Get("/{id}", d.controllers.Get)
	})
}
