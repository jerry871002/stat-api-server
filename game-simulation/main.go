package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type Batter struct {
	Name       string `json:"name"`
	AtBat      int    `json:"at_bat"`
	Hit        int    `json:"hit"`
	Double     int    `json:"double"`
	Triple     int    `json:"triple"`
	HomeRun    int    `json:"home_run"`
	Walk       int    `json:"walk"`
	HitByPitch int    `json:"hit_by_pitch"`
}

func (b *Batter) Single() int {
	return b.Hit - b.Double - b.Triple - b.HomeRun
}

func (b *Batter) PlateAppearance() int {
	return b.AtBat + b.Walk + b.HitByPitch
}

func (b *Batter) OutProbability() float64 {
	return float64(b.AtBat-b.Hit) / float64(b.PlateAppearance())
}

func (b *Batter) AdvanceProbability() map[int]float64 {
	return map[int]float64{
		1: float64(b.Single()+b.Walk+b.HitByPitch) / float64(b.PlateAppearance()),
		2: float64(b.Double) / float64(b.PlateAppearance()),
		3: float64(b.Triple) / float64(b.PlateAppearance()),
		4: float64(b.HomeRun) / float64(b.PlateAppearance()),
	}
}

type BaseballGame struct {
	Inning    int
	Outs      int
	Score     int
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
	if rand.Float64() < batter.OutProbability() {
		log.Printf("Batter %s is out", batter.Name)
		g.Outs++
		if g.Outs == 3 {
			g.EndOfInning()
		}
		return
	}

	advanceBases := g.GetAdvanceBases(batter)
	if advanceBases == 4 {
		g.HandleHomeRun(batter)
	} else {
		g.HandleAdvance(batter, advanceBases)
	}
}

func (g *BaseballGame) EndOfInning() {
	if g.Inning == 9 {
		g.EndOfGame = true
		log.Printf("End of game. Final score: %d", g.Score)
	} else {
		log.Printf("End of inning %d. Score: %d", g.Inning, g.Score)
		g.Inning++
		g.Outs = 0
		g.Runners = []int{0, 0, 0}
	}
}

func (g *BaseballGame) HandleHomeRun(batter *Batter) {
	score := sum(g.Runners) + 1
	log.Printf("Batter %s hits a home run with %d RBIs", batter.Name, score)
	g.Score += score
	g.Runners = []int{0, 0, 0}
}

func (g *BaseballGame) HandleAdvance(batter *Batter, advanceBases int) {
	score := sum(g.Runners[len(g.Runners)-advanceBases:])
	log.Printf("Batter %s advances %d bases with %d RBIs", batter.Name, advanceBases, score)
	g.Score += score
	newRunners := make([]int, advanceBases-1)
	newRunners = append(newRunners, 1)
	g.Runners = append(newRunners, g.Runners[:len(g.Runners)-advanceBases]...)
}

func (g *BaseballGame) GetAdvanceBases(batter *Batter) int {
	advanceProbability := batter.AdvanceProbability()
	keys := []int{1, 2, 3, 4}
	weights := []float64{
		advanceProbability[1],
		advanceProbability[2],
		advanceProbability[3],
		advanceProbability[4],
	}
	return weightedChoice(keys, weights)
}

func weightedChoice(keys []int, weights []float64) int {
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

func simulateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var lineup []Batter
	if err := json.NewDecoder(r.Body).Decode(&lineup); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(lineup) != 9 {
		http.Error(w, "Lineup must have 9 batters", http.StatusBadRequest)
		return
	}

	game := NewBaseballGame()

	score := 0
	for i := 0; i < 1000; i++ {
		game.SimulateGame(lineup)
		score += game.Score
	}

	averageScore := float64(score) / 1000
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"average_score": averageScore})
}

func main() {
	http.HandleFunc("/simulate", simulateHandler)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}