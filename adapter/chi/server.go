package chi

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	middlewares "github.com/renatospaka/scheduler/adapter/web/middleware"
)

type HttpServer struct {
	ctx         context.Context
	// controllers *controller.TransactionController
	Server      *chi.Mux
}

func NewHttpServer(ctx context.Context) *HttpServer {
	log.Println("initiating web seerver")
	httpServer := &HttpServer{
		ctx:         ctx,
		// controllers: controller,
	}
	httpServer.connect()

	return httpServer
}

func (s *HttpServer) connect() {
	s.Server = chi.NewRouter()
	s.Server.Use(middleware.Logger)
	s.Server.Use(middleware.Recoverer)
	s.Server.Use(middlewares.Cors)
	// s.Server.Use(middleware.WithValue("jwt", configs.TokenAuth))
	// s.Server.Use(middleware.WithValue("JWTExpiresIn", configs.JWTExpiresIn))

	s.Server.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode("Healthy")
		})
	})
}
