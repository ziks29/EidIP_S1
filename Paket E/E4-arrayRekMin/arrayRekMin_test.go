package main

import (
	"math/rand"
	"testing"
)

func arrayMin(arr [arrLen]int, start, ende int) int {
	result := arr[start]
	for i := start + 1; i < ende; i++ {
		result = min(result, arr[i])
	}
	return result
}

func TestArrayRekMin(t *testing.T) {
	var arr [arrLen]int
	for i := 1; i <= 10000; i++ {
		tatsLen := rand.Intn(arrLen-1)+1
		for j := 0; j < tatsLen; j++ {
			arr[j] = rand.Intn(1000) - rand.Intn(500) + rand.Intn(500)
		}
		got := arrayRekMin(arr, 0, tatsLen)
		want := arrayMin(arr, 0, tatsLen)
		if got != want {
			t.Errorf("arrayRekMin(%v) = %v, want %v", arr[0:tatsLen], got, want)
		}
	}
}
