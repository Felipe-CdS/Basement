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
	Db      *sql.DB
	Testing bool
}

func NewPostgresDB(t bool) *sql.DB {

	dbStore := Store{Testing: t}

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

	if dbStore.Testing {
		log.Println("Connecting to tests database...")
		dbname = "basement_tests"
	}

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

	// CREATE DATABASE basement TEMPLATE template0;
	// SET timezone TO 'UTC';

	statement := `
	CREATE TABLE IF NOT EXISTS tags (
		"id" SERIAL PRIMARY KEY
		, "type" VARCHAR (50) NOT NULL
		, "name" VARCHAR (50) NOT NULL UNIQUE
	);

	CREATE TABLE IF NOT EXISTS activities (
		"id" SMALLSERIAL PRIMARY KEY
		, start_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
		, end_time TIMESTAMP WITH TIME ZONE
		, description TEXT
	);

	CREATE TABLE IF NOT EXISTS activities_tags (
		fk_activity_id INTEGER NOT NULL
		, fk_tag_id INTEGER NOT NULL

		, PRIMARY KEY(fk_activity_id, fk_tag_id)
		, CONSTRAINT fk_tag
			FOREIGN KEY(fk_tag_id)
			REFERENCES tags("id")
			ON UPDATE CASCADE
			ON DELETE CASCADE
		, CONSTRAINT fk_activity
			FOREIGN KEY(fk_activity_id)
			REFERENCES activities("id")
			ON UPDATE CASCADE
			ON DELETE CASCADE
	);`

	_, err := db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 1... Error: %s", err)
	}

	return nil
}

func DropTestMigrations(db *sql.DB) error {

	statement := `
		DROP TABLE activities_tags;
		DROP TABLE activities;
		DROP TABLE tags;
	`

	_, err := db.Exec(statement)

	if err != nil {

		log.Println("Error droping Migrations", err)
		return fmt.Errorf("failed to create statement 1... Error: %s", err)
	}

	log.Println("Dropped Migrations")
	return nil
}
