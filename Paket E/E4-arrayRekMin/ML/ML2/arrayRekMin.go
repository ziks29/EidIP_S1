package main

// arrayRekMin gibt den kleinsten Wert von arr im Indexbereich start bis inklusive ende-1 zurück.
// Dabei muss start<ende<=arrLen gelten.
func arrayRekMin(arr [arrLen]int, start, ende int) int {
	// wenn das Array nur ein Element enthält, ...
	if ende-start == 1 {
		// ... dann ist das Minimum dieses Element
		return arr[start]
	} else {
		// ... ansonsten ist das Minimum des Arrays das Minimum aus der linken und rechten
		// Hälfte des Arrays („divide-and-conquer“-Verfahren)
		mitte := (start + ende - 1) / 2
		return min(arrayRekMin(arr, start, mitte+1), arrayRekMin(arr, mitte+1, ende))
	}
}
