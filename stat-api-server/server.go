package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Team struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

type StatStore interface {
	GetTeams() ([]Team, error)
	GetBattingStat(team string, year int) ([]map[string]any, error)
}

type StatServer struct {
	store StatStore
}

func (s *StatServer) GetTeamsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetTeamsHandler is called")

	data, err := s.store.GetTeams()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (s *StatServer) GetBattingStatHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBattingStatHandler is called")

	query := r.URL.Query()
	team := query.Get("team")
	year, err := strconv.Atoi(query.Get("year"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Team:", team, "year:", year)

	data, err := s.store.GetBattingStat(team, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(data)
}
