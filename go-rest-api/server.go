package main

import (
	"encoding/json"
	"net/http"
)

type StatStore interface {
	GetPitching() ([]map[string]any, error)
	GetBatting() ([]map[string]any, error)
	GetFielding() ([]map[string]any, error)
}

type StatServer struct {
	store StatStore
}

func (s *StatServer) GetPitchingHandler(w http.ResponseWriter, r *http.Request) {
	data, err := s.store.GetPitching()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (s *StatServer) GetBattingHandler(w http.ResponseWriter, r *http.Request) {
	data, err := s.store.GetBatting()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (s *StatServer) GetFieldingHandler(w http.ResponseWriter, r *http.Request) {
	data, err := s.store.GetFielding()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}
