package main

import (
	"fmt"
	"log"
	"math/rand"
)

type Batter struct {
	Name       string
	AtBat      int
	Hit        int
	Double     int
	Triple     int
	HomeRun    int
	Walk       int
	HitByPitch int
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
	r := rand.Float64() * totalWeight
	for i, weight := range weights {
		if r < weight {
			return keys[i]
		}
		r -= weight
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

func main() {
	jerry := Batter{"Jerry", 30, 9, 2, 1, 1, 4, 1}
	game := NewBaseballGame()

	score := 0
	for i := 0; i < 1000; i++ {
		game.SimulateGame([]Batter{jerry, jerry, jerry, jerry, jerry, jerry, jerry, jerry, jerry})
		score += game.Score
	}

	fmt.Printf("Average score: %.2f\n", float64(score)/1000)
}
