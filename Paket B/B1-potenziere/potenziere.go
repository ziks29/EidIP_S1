package main

// potenziere gibt den Potenzwert basis hoch exponent zurück.
// exponent muss >= 0 sein.
func potenziere(basis, exponent int) int {
	// if exponent == 0 {
	// 	return 1
	// }
	// if exponent > 0 {
	// 	return basis * potenziere(basis, exponent-1)
	// }

	// With for loop
	result := 1
	for i := 1; i <= exponent; i++ {
		result = result * basis
	}
	return result

}
