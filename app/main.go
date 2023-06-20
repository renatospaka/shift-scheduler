package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/renatospaka/scheduler/scheduler"
	"github.com/renatospaka/scheduler/utils/configs"
)

func main() {
	log.Println("initiating application")

	// start all configuration
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Panic(err)
	}

	// start database server
	log.Println("initiating database connection")
	conn := "postgresql://" + configs.DBUser + ":" + configs.DBPassword + "@" + configs.DBHost + ":" + configs.DBPort + "/" + configs.DBName + "?sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// start application server
	ctx := context.Background()
	schedulerApp := scheduler.NewSchedulerDomain(ctx, db)

	// start web server
	log.Printf("scheduler server is listening to port: %s\n", configs.WEBServerPort)
	http.ListenAndServe(":"+configs.WEBServerPort, schedulerApp.WebServer.Server)
}
