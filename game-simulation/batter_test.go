package main

import (
	"math"
	"reflect"
	"testing"
)

func TestBatterSingle(t *testing.T) {
	b := &Batter{
		Hit:    20,
		Double: 5,
		Triple: 2,
		HomeRun: 3,
	}
	expected := 10
	if got := b.Single(); got != expected {
		t.Errorf("Batter.Single() = %v, want %v", got, expected)
	}
}

func TestBatterPlateAppearance(t *testing.T) {
	b := &Batter{
		AtBat:      100,
		BallOnBase: 20,
		HitByPitch: 5,
	}
	expected := 125
	if got := b.PlateAppearance(); got != expected {
		t.Errorf("Batter.PlateAppearance() = %v, want %v", got, expected)
	}
}

func TestBatterOutProbability(t *testing.T) {
	b := &Batter{
		AtBat:  100,
		Hit:    30,
	}
	expected := 0.7
	if got := b.OutProbability(); !FloatEqual(got, expected) {
		t.Errorf("Batter.OutProbability() = %v, want %v", got, expected)
	}
}

func TestBatterBallOnBaseProbability(t *testing.T) {
	b := &Batter{
		AtBat:      100,
		BallOnBase: 20,
		HitByPitch: 5,
	}
	expected := 0.16
	if got := b.BallOnBaseProbability(); !FloatEqual(got, expected) {
		t.Errorf("Batter.BallOnBaseProbability() = %v, want %v", got, expected)
	}
}

func TestBatterHitByPitchProbability(t *testing.T) {
	b := &Batter{
		AtBat:      100,
		BallOnBase: 20,
		HitByPitch: 5,
	}
	expected := 0.04
	if got := b.HitByPitchProbability(); !FloatEqual(got, expected) {
		t.Errorf("Batter.HitByPitchProbability() = %v, want %v", got, expected)
	}
}

func TestBatterHitProbability(t *testing.T) {
	b := &Batter{
		AtBat:      100,
		Hit:		30,
		BallOnBase: 20,
		HitByPitch: 5,
	}
	expected := 0.24
	if got := b.HitProbability(); !FloatEqual(got, expected) {
		t.Errorf("Batter.HitProbability() = %v, want %v", got, expected)
	}
}

func TestBatterHitAdvanceProbability(t *testing.T) {
	b := &Batter{
		AtBat:      100,
		Hit:    30,
		Double: 5,
		Triple: 2,
		HomeRun: 3,
		BallOnBase: 20,
		HitByPitch: 5,
	}
	expected := map[int]float64{
		1: 0.16,
		2: 0.04,
		3: 0.016,
		4: 0.024,
	}
	if got := b.HitAdvanceProbability(); !reflect.DeepEqual(got, expected) {
		t.Errorf("Batter.HitAdvanceProbability() = %v, want %v", got, expected)
	}
}

func FloatEqual(a, b float64, args ...float64) bool {
	epsilon := 1e-9
	if len(args) > 0 {
		epsilon = args[0]
	}
	return math.Abs(a-b) < epsilon
}