package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"time"

	"github.com/rs/cors"
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	numGames    int
	numBatches  int
)

type Batter struct {
	Name       string `json:"name"`
	AtBat      int    `json:"at_bat"`
	Hit        int    `json:"hit"`
	Double     int    `json:"double"`
	Triple     int    `json:"triple"`
	HomeRun    int    `json:"home_run"`
	BallOnBase int    `json:"ball_on_base"`
	HitByPitch int    `json:"hit_by_pitch"`
}

func (b *Batter) Single() int {
	return b.Hit - b.Double - b.Triple - b.HomeRun
}

func (b *Batter) PlateAppearance() int {
	return b.AtBat + b.BallOnBase + b.HitByPitch
}

func (b *Batter) OutProbability() float64 {
	return float64(b.AtBat-b.Hit) / float64(b.PlateAppearance())
}

func (b *Batter) BallOnBaseProbability() float64 {
	return float64(b.BallOnBase) / float64(b.PlateAppearance())
}

func (b *Batter) HitByPitchProbability() float64 {
	return float64(b.HitByPitch) / float64(b.PlateAppearance())
}

func (b *Batter) HitProbability() float64 {
	return float64(b.Hit) / float64(b.PlateAppearance())
}

func (b *Batter) HitAdvanceProbability() map[int]float64 {
	return map[int]float64{
		1: float64(b.Single()) / float64(b.PlateAppearance()),
		2: float64(b.Double) / float64(b.PlateAppearance()),
		3: float64(b.Triple) / float64(b.PlateAppearance()),
		4: float64(b.HomeRun) / float64(b.PlateAppearance()),
	}
}

type BaseballGame struct {
	Inning    int
	Outs      int
	Score     int
	Hits	  int
	Runners   []int
	EndOfGame bool
}

func NewBaseballGame() *BaseballGame {
	return &BaseballGame{}
}

func (g *BaseballGame) Reset() {
	g.Inning = 1
	g.Outs = 0
	g.Score = 0
	g.Hits = 0
	g.Runners = []int{0, 0, 0}
	g.EndOfGame = false
}

func (g *BaseballGame) SimulateGame(lineup []Batter) {
	if len(lineup) != 9 {
		log.Fatal("Lineup must have 9 batters")
	}

	g.Reset()
	battingOrder := lineup
	for !g.EndOfGame {
		batter := battingOrder[0]
		g.SimulateOneBatter(&batter)
		battingOrder = append(battingOrder[1:], batter)
	}
}

func (g *BaseballGame) SimulateOneBatter(batter *Batter) {
	outcome := weightedChoice(
		[]string{"out", "ball_on_base", "hit_by_pitch", "hit"},
		[]float64{batter.OutProbability(), batter.BallOnBaseProbability(), batter.HitByPitchProbability(), batter.HitProbability()},
	)

	switch outcome {
	case "out":
		g.HandleOut(batter)
	case "ball_on_base", "hit_by_pitch":
		g.HandleAwardBase(batter)
	case "hit":
		g.HandleHit(batter)
	}
}

func (g *BaseballGame) HandleOut(batter *Batter) {
	debugLogger.Printf("Batter %s is out", batter.Name)
	g.Outs++
	if g.Outs == 3 {
		g.EndOfInning()
	}
}

func (g *BaseballGame) EndOfInning() {
	if g.Inning == 9 {
		g.EndOfGame = true
		debugLogger.Printf("End of game. Final score: %d", g.Score)
	} else {
		debugLogger.Printf("End of inning %d. Score: %d", g.Inning, g.Score)
		g.Inning++
		g.Outs = 0
		g.Runners = []int{0, 0, 0}
	}
}

func (g *BaseballGame) HandleAwardBase(batter *Batter) {
	debugLogger.Printf("Batter %s is awarded to first base (BB or HBP)", batter.Name)
	if g.Runners[0] == 0 {
		g.Runners[0] = 1
	} else if reflect.DeepEqual(g.Runners, []int{1, 1, 1}) { // Bases loaded
		debugLogger.Printf("Batter %s got 1 RBI", batter.Name)
		g.Score++
	} else if sum(g.Runners) == 2 {
		g.Runners = []int{1, 1, 1}
	} else {
		g.Runners = append([]int{1}, g.Runners[:len(g.Runners)-1]...)
	}
}

func (g *BaseballGame) HandleHit(batter *Batter) {
	g.Hits++
	advanceBases := g.GetHitAdvanceBases(batter)
	if advanceBases == 4 {
		g.HandleHomeRun(batter)
	} else {
		g.HandleHitAdvance(batter, advanceBases)
	}
}

func (g *BaseballGame) HandleHomeRun(batter *Batter) {
	score := sum(g.Runners) + 1
	debugLogger.Printf("Batter %s hits a home run with %d RBIs", batter.Name, score)
	g.Score += score
	g.Runners = []int{0, 0, 0}
}

func (g *BaseballGame) HandleHitAdvance(batter *Batter, advanceBases int) {
	score := sum(g.Runners[len(g.Runners)-advanceBases:])
	var hitType string
	switch advanceBases {
	case 1:
		hitType = "single"
	case 2:
		hitType = "double"
	case 3:
		hitType = "triple"
	}
	debugLogger.Printf("Batter %s hits a %s with %d RBIs", batter.Name, hitType, score)
	g.Score += score
	newRunners := make([]int, advanceBases-1)
	newRunners = append(newRunners, 1)
	g.Runners = append(newRunners, g.Runners[:len(g.Runners)-advanceBases]...)
}

func (g *BaseballGame) GetHitAdvanceBases(batter *Batter) int {
	advanceProbability := batter.HitAdvanceProbability()
	keys := []int{1, 2, 3, 4}
	weights := []float64{
		advanceProbability[1],
		advanceProbability[2],
		advanceProbability[3],
		advanceProbability[4],
	}
	return weightedChoice(keys, weights)
}

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

func configureEnvironment() {
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// debug messages are printed only when DEBUG=1 or DEBUG=true
	if os.Getenv("DEBUG") != "1" && os.Getenv("DEBUG") != "true" {
		debugLogger.SetOutput(io.Discard)
		numGames = 100000
		numBatches = runtime.NumCPU()
	} else {
		numGames = 10
		numBatches = 1
	}
}

func main() {
	configureEnvironment()

	http.HandleFunc("/simulate", simulateHandler)

	handler := cors.Default().Handler(http.DefaultServeMux)

	log.Println("Server is running on port 80")
	log.Fatal(http.ListenAndServe(":80", handler))
}
