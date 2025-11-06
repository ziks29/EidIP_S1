package main

// potenziere gibt den Potenzwert basis hoch exponent zurück.
// exponent muss >= 0 sein.
func potenziere(basis, exponent int) int {
	result := 1
	for i := 1; i <= exponent; i++ {
		result = result * basis
	}
	return result
}
