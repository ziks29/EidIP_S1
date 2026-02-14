package main

import (
	"math/rand"
	"testing"
)

func spiegelzahlIterativ(z int) int {
	result := 0
	p := 1
	for z := z; z >= 10; {
		p = p * 10
		z = z / 10
	}
	for z > 0 {
		result = result + p*(z%10)
		p = p / 10
		z = z / 10
	}
	return result
}

func TestSpiegelzahl(t *testing.T) {
	for z := 0; z <= 1111; z++ {
		if z%10 != 0 {
			got := spiegelzahl(z)
			want := spiegelzahlIterativ(z)
			if got != want {
				t.Errorf("spiegelzahl(%v) = %v, want %v", z, got, want)
			}
		}
	}
	for i := 1; i <= 1000; i++ {
		z := rand.Intn(1000000000) + 1111
		if z%10 != 0 {
			got := spiegelzahl(z)
			want := spiegelzahlIterativ(z)
			if got != want {
				t.Errorf("spiegelzahl(%v) = %v, want %v", z, got, want)
			}
		}
	}
	got := spiegelzahl(0)
	if got != 0 {
		t.Errorf("spiegelzahl(%v) = %v, want %v", 0, got, 0)
	}
}
