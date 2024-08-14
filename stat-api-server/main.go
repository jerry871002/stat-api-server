package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	connStr := "postgres://myuser:mypassword@db/mydatabase?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Connected to database")

	server := &StatServer{
		store: NewSqlStatStore(db),
	}

	router := mux.NewRouter()
	router.StrictSlash(true) // "/path/" and "/path" will be treated as the same path
	router.HandleFunc("/teams/", server.GetTeamsHandler).Methods("GET")
	router.HandleFunc("/batting/", server.GetBattingStatHandler).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Println("Server started at :80")
	log.Fatal(http.ListenAndServe(":80", handler))
}
