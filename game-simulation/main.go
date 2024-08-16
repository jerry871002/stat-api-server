package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/rs/cors"
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	debugMode   bool
)

func weightedChoice[T any](keys []T, weights []float64) T {
	totalWeight := 0.0
	for _, weight := range weights {
		totalWeight += weight
	}
	randVal := rand.Float64() * totalWeight
	for i, weight := range weights {
		if randVal < weight {
			return keys[i]
		}
		randVal -= weight
	}
	return keys[len(keys)-1]
}

func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

func simulateBatchWorker(lineup []Batter, numGames int, scoreChan chan<- int, hitChan chan<- int) {
	startTime := time.Now()

	game := NewBaseballGame()
	scores := 0
	hits := 0
	for i := 0; i < numGames; i++ {
		game.SimulateGame(lineup)
		scores += game.Score
		hits += game.Hits
	}
	scoreChan <- scores
	hitChan <- hits

	elapsedTime := time.Since(startTime)
	infoLogger.Printf("simulateBatchWorker took %s to simulate %d games", elapsedTime, numGames)
}

func simulateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("simulateHandler is called")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Println("Invalid request method:", r.Method)
		return
	}

	var lineup []Batter
	if err := json.NewDecoder(r.Body).Decode(&lineup); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println(err)
		return
	}
	infoLogger.Printf("Received lineup: %v", lineup)

	if len(lineup) != 9 {
		http.Error(w, "Lineup must have 9 batters", http.StatusBadRequest)
		return
	}

	var numGames, numBatches int
	if debugMode {
		numGames = 10
		numBatches = 1
	} else {
		numGames = 100000
		numBatches = runtime.NumCPU()
	}

	gamePerBatch := numGames / numBatches

	scoreChan := make(chan int)
	hitChan := make(chan int)
	for i := 0; i < numBatches; i++ {
		if i == numBatches-1 {
			gamePerBatch = numGames - (gamePerBatch * (numBatches - 1))
		}
		go simulateBatchWorker(lineup, gamePerBatch, scoreChan, hitChan)
	}

	totalScore := 0
	totalHits := 0
	for i := 0; i < numBatches; i++ {
		totalScore += <-scoreChan
		totalHits += <-hitChan
	}

	averageScore := float64(totalScore) / float64(numGames)
	averageHits := float64(totalHits) / float64(numGames)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{
		"average_score": averageScore,
		"average_hits": averageHits,
	})
}

func initLogger() {
	debugLogger = log.New(
		os.Stdout,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
	)
	infoLogger = log.New(
		os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
	)
}

func configEnv() {
	initLogger()

	// debug messages are printed only when DEBUG=1 or DEBUG=true
	if os.Getenv("DEBUG") != "1" && os.Getenv("DEBUG") != "true" {
		debugMode = true
		debugLogger.SetOutput(io.Discard)
	} else {
		debugMode = false
	}
}

func main() {
	configEnv()

	http.HandleFunc("/simulate", simulateHandler)

	handler := cors.Default().Handler(http.DefaultServeMux)

	log.Println("Server is running on port 80")
	log.Fatal(http.ListenAndServe(":80", handler))
}
