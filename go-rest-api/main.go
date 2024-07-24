package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/batting/", GetBatting).Methods("GET")
	router.HandleFunc("/pitching/", GetPitching).Methods("GET")
	router.HandleFunc("/fielding/", GetFielding).Methods("GET")

	log.Println("Server started at :18001")
	log.Fatal(http.ListenAndServe(":18001", router))
}

func GetBatting(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./puumat_stats.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM batting")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}

	var data []map[string]any
	cols, _ := rows.Columns()
	for rows.Next() {
		columns := make([]any, len(cols))
		columnPointers := make([]any, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}

		m := make(map[string]any)
		for i, colName := range cols {
			val := columnPointers[i].(*any)
			m[colName] = *val
		}
		data = append(data, m)
	}
	json.NewEncoder(w).Encode(data)
}

func GetPitching(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./puumat_stats.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM pitching")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}

	var data []map[string]any
	cols, _ := rows.Columns()
	for rows.Next() {
		columns := make([]any, len(cols))
		columnPointers := make([]any, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}

		m := make(map[string]any)
		for i, colName := range cols {
			val := columnPointers[i].(*any)
			m[colName] = *val
		}
		data = append(data, m)
	}
	json.NewEncoder(w).Encode(data)
}

func GetFielding(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./puumat_stats.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM fielding")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}

	var data []map[string]any
	cols, _ := rows.Columns()
	for rows.Next() {
		columns := make([]any, len(cols))
		columnPointers := make([]any, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}

		m := make(map[string]any)
		for i, colName := range cols {
			val := columnPointers[i].(*any)
			m[colName] = *val
		}
		data = append(data, m)
	}
	json.NewEncoder(w).Encode(data)
}