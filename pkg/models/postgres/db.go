package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Store struct {
	Db *sql.DB
}

func NewPostgresDB() *sql.DB {

	dbStore := Store{}

	if err := dbStore.getConnection(); err != nil {
		log.Fatalf("failed to connect to the database... Error: %s", err)
	}

	if err := createMigrations(dbStore.Db); err != nil {
		log.Fatalln(err)
	}

	return dbStore.Db
}

func (dbStore *Store) getConnection() error {

	if dbStore.Db != nil {
		return nil
	}

	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "basement"

	if os.Getenv("APP_ENV") == "TESTING" {
		port = 8001
	}

	if os.Getenv("APP_ENV") == "PROD" {
		host = os.Getenv("PGHOST")
		user = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname = os.Getenv("PGDATABASE")
		port, _ = strconv.Atoi(os.Getenv("PGPORT"))
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	dbStore.Db = db
	log.Printf("Connected successfully to the database | ENV: %s | host: %s | dbname: %s\n", os.Getenv("APP_ENV"), host, dbname)

	return nil
}

func createMigrations(db *sql.DB) error {

	statement := `CREATE TABLE IF NOT EXISTS activities (
		id INTEGER PRIMARY KEY
		, activity_type VARCHAR
		, start_time TIMESTAMP WITH TIME ZONE NOT NULL
		, end_time TIMESTAMP WITH TIME ZONE
	)`

	_, err := db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 1... Error: %s", err)
	}

	statement = `CREATE TABLE IF NOT EXISTS day_stats (
		id UUID PRIMARY KEY
		, date DATE NOT NULL
		, study INTEGER NOT NULL DEFAULT 0
		, programming_work INTEGER NOT NULL DEFAULT 0
		, programming_hobby INTEGER NOT NULL DEFAULT 0
		, read_study INTEGER NOT NULL DEFAULT 0
		, read_fun INTEGER NOT NULL DEFAULT 0
		, garbage INTEGER NOT NULL DEFAULT 0
	)`

	_, err = db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 2... Error: %s", err)
	}

	return nil
}
