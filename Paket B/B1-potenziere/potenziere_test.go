package main

import (
	"math"
	"math/rand"
	"testing"
)

func TestPotenziereZufällig(t *testing.T) {
	for i := 1; i <= 10000; i++ {
		basis := rand.Intn(99)
		exponent := rand.Intn(9)
		want := int(math.Pow(float64(basis), float64(exponent)))
		got := potenziere(basis, exponent)
		if got != want {
			t.Errorf("potenziere(%v,%v)=%v, want %v", basis, exponent, got, want)
		}
	}
}

func TestPotenziereSpezWerte(t *testing.T) {
	tests := [...]struct {
		basis    int
		exponent int
		want     int
	}{
		{0, 0, 1},
		{0, 1, 0},
		{0, 99, 0},
		{1, 0, 1},
		{1, 1, 1},
		{1, 98, 1},
	}
	for _, tt := range tests {
		got := potenziere(tt.basis, tt.exponent)
		if got != tt.want {
			t.Errorf("potentiere(%v,%v) = %v, want %v", tt.basis, tt.exponent, got, tt.want)
		}
	}
}
