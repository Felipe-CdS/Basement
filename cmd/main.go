package main

import (
	"fmt"
	"log"
	"net/http"

	"nugu.dev/basement/pkg/models/postgres"
)

type application struct {
	activities *postgres.ActivityModel
	dayStats   *postgres.DayStatsModel
}

func main() {

	app := &application{
		activities: &postgres.ActivityModel{Db: postgres.NewPostgresDB()},
		dayStats:   &postgres.DayStatsModel{Db: postgres.NewPostgresDB()},
	}

	srv := &http.Server{
		Addr:    ":3000",
		Handler: app.routes(),
	}

	fmt.Println("Server Working...")
	err := srv.ListenAndServe()
	log.Fatalln(err)
}
