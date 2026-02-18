package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// arrayRekMin gibt den kleinsten Wert von arr im Indexbereich start bis inklusive ende-1 zurück.
// Dabei muss start<ende<=arrLen gelten.
func arrayRekMin(arr [arrLen]int, start, ende int) int {
	if start == ende-1 {
		return arr[start]
	}
	return min(arr[start], arrayRekMin(arr, start+1, ende))
}
