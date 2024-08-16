package main

import (
	"reflect"
	"testing"
)

func TestHandleHitAdvance(t *testing.T) {
	g := &BaseballGame{
		Runners: []int{0, 1, 1},
		Score:   2,
	}

	b := &Batter{
		Name: "John Doe",
	}

	initLogger()

	g.HandleHitAdvance(b, 2)

	expectedScore := 4
	if g.Score != expectedScore {
		t.Errorf("HandleHitAdvance() score = %d, want %d", g.Score, expectedScore)
	}

	expectedRunners := []int{0, 1, 0}
	if !reflect.DeepEqual(g.Runners, expectedRunners) {
		t.Errorf("HandleHitAdvance() runners = %v, want %v", g.Runners, expectedRunners)
	}
}

func TestHandleHomeRun(t *testing.T) {
	// Runner on 2nd and 3rd
	g := &BaseballGame{
		Runners: []int{0, 1, 1},
		Score:   2,
	}

	b := &Batter{
		Name: "John Doe",
	}

	g.HandleHomeRun(b)

	expectedScore := 5
	if g.Score != expectedScore {
		t.Errorf("HandleHitAdvance() score = %d, want %d", g.Score, expectedScore)
	}

	expectedRunners := []int{0, 0, 0}
	if !reflect.DeepEqual(g.Runners, expectedRunners) {
		t.Errorf("HandleHitAdvance() runners = %v, want %v", g.Runners, expectedRunners)
	}

	// Bases empty
	g = &BaseballGame{
		Runners: []int{0, 0, 0},
		Score:   0,
	}

	b = &Batter{
		Name: "John Doe",
	}

	g.HandleHomeRun(b)

	expectedScore = 1
	if g.Score != expectedScore {
		t.Errorf("HandleHitAdvance() score = %d, want %d", g.Score, expectedScore)
	}

	expectedRunners = []int{0, 0, 0}
	if !reflect.DeepEqual(g.Runners, expectedRunners) {
		t.Errorf("HandleHitAdvance() runners = %v, want %v", g.Runners, expectedRunners)
	}
}