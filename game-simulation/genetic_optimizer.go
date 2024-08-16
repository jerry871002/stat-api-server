package main

import (
	"math/rand"
	"runtime"
)

type GeneticOptimizer struct {
	Optimizer
	MaxGeneration int
	PopulationSize int
}

func NewGeneticOptimizer(maxGeneration, populationSize int) *GeneticOptimizer {
	return &GeneticOptimizer{
		MaxGeneration: maxGeneration,
		PopulationSize: populationSize,
	}
}

func (g *GeneticOptimizer) Optimize(roster Roster) Lineup {
	population := make([]Lineup, g.PopulationSize)
	for i := 0; i < g.PopulationSize; i++ {
		population[i] = g.GenerateRandomLineup(roster)
	}
	infoLogger.Println("Initial population:", population)

	for gen := 0; gen < g.MaxGeneration; gen++ {
		// Evaluate population
		fitnessScores := make([]float64, g.PopulationSize)
		for i, lineup := range population {
			fitnessScores[i] = g.ComputeFitness(lineup)
		}

		infoLogger.Println("fitnessScores:", fitnessScores)

		// Select parents
		parents := g.SelectParents(population, fitnessScores)

		// Crossover and mutate
		population = g.CrossoverAndMutate(parents)
	}

	// Find the best lineup
	maxFitnessScore := 0.0
	bestLineup := Lineup{}
	for i, lineup := range population {
		fitnessScore := g.ComputeFitness(lineup)
		if fitnessScore > maxFitnessScore {
			maxFitnessScore = fitnessScore
			bestLineup = population[i]
		}
	}

	infoLogger.Println("bestLineup:", bestLineup, "with fitness score:", maxFitnessScore)

	return bestLineup
}

func (g *GeneticOptimizer) GenerateRandomLineup(roster Roster) Lineup {
	lineup := make([]Batter, 9)
	inLineup := map[string]bool{}
	for i := 0; i < 9; i++ {
		var candidate Batter
		for {
			candidate = roster[rand.Intn(len(roster))]
			if !inMap(candidate.Name, inLineup) {
				break
			}
		}
		lineup[i] = candidate
		inLineup[candidate.Name] = true
	}
	return lineup
}

func (g *GeneticOptimizer) ComputeFitness(lineup Lineup) float64 {
	numGames := 1000
	numBatches := runtime.NumCPU()
	result := simulateGamesInParallel(lineup, numGames, numBatches)
	return result["average_score"] * 100 + result["average_hits"]
}

func (g *GeneticOptimizer) SelectParents(population []Lineup, fitnessScores []float64) []Lineup {
	parents := make([]Lineup, g.PopulationSize)
	for i := 0; i < g.PopulationSize; i++ {
		a, b := rand.Intn(g.PopulationSize), rand.Intn(g.PopulationSize)
		if fitnessScores[a] > fitnessScores[b] {
			parents[i] = population[a]
		} else {
			parents[i] = population[b]
		}
	}
	return parents
}

func (g *GeneticOptimizer) CrossoverAndMutate(parents []Lineup) []Lineup {
	children := make([]Lineup, g.PopulationSize)
	for i := 0; i < len(parents) / 2; i++ {
		p1, p2 := parents[2*i], parents[2*i + 1]
		c1, c2 := g.Crossover(p1, p2)
		children[2*i], children[2*i + 1] = g.Mutate(c1, c2)
	}
	return children
}

func (g *GeneticOptimizer) Crossover(p1, p2 Lineup) (Lineup, Lineup) {
	child1, child2 := make([]Batter, 9), make([]Batter, 9)

	for _, child := range []Lineup{child1, child2} {
		standby := []Batter{}
		inChild := map[string]bool{}
		for i := 0; i < 9; i++ {
			candidate1, candidate2 := p1[i], p2[i]
			if inMap(candidate1.Name, inChild) && inMap(candidate2.Name, inChild) {
				// Both candidates are already in child, use standby list
				var standbyCandidate Batter
				for {
					standbyCandidate = standby[rand.Intn(len(standby))]
					if !inMap(standbyCandidate.Name, inChild) {
						break
					}
				}
				child[i] = standbyCandidate
				inChild[standbyCandidate.Name] = true
			} else if inMap(candidate1.Name, inChild) {
				child[i] = candidate2
				inChild[candidate2.Name] = true
				standby = append(standby, candidate1)
			} else if inMap(candidate2.Name, inChild) {
				child[i] = candidate1
				inChild[candidate1.Name] = true
				standby = append(standby, candidate2)
			} else {
				if rand.Float64() < 0.5 {
					child[i] = candidate1
					inChild[candidate1.Name] = true
					standby = append(standby, candidate2)
				} else {
					child[i] = candidate2
					inChild[candidate2.Name] = true
					standby = append(standby, candidate1)
				}
			}
		}
	}

	return child1, child2
}

func (g *GeneticOptimizer) Mutate(c1, c2 Lineup) (Lineup, Lineup) {
	for _, child := range []Lineup{c1, c2} {
		if rand.Float64() < 0.1 {
			i, j := rand.Intn(9), rand.Intn(9)
			child[i], child[j] = child[j], child[i]
		}
	}
	return c1, c2
}

func inMap[T comparable](key T, m map[T]bool) bool {
	_, ok := m[key]
	return ok
}
