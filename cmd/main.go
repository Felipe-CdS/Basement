package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"nugu.dev/basement/pkg/models/postgres"
)

type application struct {
	AuthToken     string
	ReadBucketURL string
	RWBucketURL   string
	activities    *postgres.ActivityModel
	dayStats      *postgres.DayStatsModel
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}

	app := &application{
		AuthToken:     "123",
		ReadBucketURL: os.Getenv("PUBLIC_SINGLE_READ_BUCKET"),
		RWBucketURL:   os.Getenv("PRIVATE_LIST_RW_BUCKET"),
		activities:    &postgres.ActivityModel{Db: postgres.NewPostgresDB()},
		dayStats:      &postgres.DayStatsModel{Db: postgres.NewPostgresDB()},
	}

	srv := &http.Server{
		Addr:    ":3000",
		Handler: app.routes(),
	}

	fmt.Println("Server Working...")
	err := srv.ListenAndServe()
	log.Fatalln(err)
}
