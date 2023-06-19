package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"context"
	"net/http"


	httpServer "github.com/renatospaka/scheduler/adapter/chi"
	"github.com/renatospaka/scheduler/utils/configs"
)

func main() {
	log.Println("initiating application")
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Panic(err)
	}

	//open connection to the database
	log.Println("initiating database connection")
	conn := "postgresql://" + configs.DBUser + ":" + configs.DBPassword + "@" + configs.DBHost + ":" + configs.DBPort + "/" + configs.DBName + "?sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	log.Println("initiating shifts scheduler")
	// web server
	ctx := context.Background()
	webServer := httpServer.NewHttpServer(ctx)

	// start web server
	log.Printf("gerador de transações escutando porta: %s\n", configs.WEBServerPort)
	http.ListenAndServe(":"+configs.WEBServerPort, webServer.Server)
}
