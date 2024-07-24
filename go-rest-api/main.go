package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./puumat_stats.db")
	defer db.Close()
	server := &StatServer{
		store: NewSqliteStatStore(db),
	}

	router := mux.NewRouter()
	router.HandleFunc("/batting/", server.GetBattingHandler).Methods("GET")
	router.HandleFunc("/pitching/", server.GetPitchingHandler).Methods("GET")
	router.HandleFunc("/fielding/", server.GetFieldingHandler).Methods("GET")

	log.Println("Server started at :18001")
	log.Fatal(http.ListenAndServe(":18001", router))
}
