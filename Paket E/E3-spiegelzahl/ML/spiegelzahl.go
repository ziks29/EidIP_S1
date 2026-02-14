package main

// zehnerpotenz git gibt die Zehnerpotenz von zahl zurück.
func zehnerpotenz(zahl int) int {
	result := 1
	for zahl >= 10 {
		result = result * 10
		zahl = zahl / 10
	}
	return result
}

// spiegelzahl gibt die Spiegelzahl von zahl zurück.
// zahl muss eine natürliche Zahl sein.
func spiegelzahl(zahl int) int {
	// einstellige Zahlen sind ihre Spiegelzahlen
	if zahl < 10 {
		return zahl
	} else {
		// sonst ist die Spiegelzahl die Summe aus
		// der letzten Ziffer der Zahl multipliziert mit ihrer Zehnerpotenz
		// und der Spiegelzahl ihrer anderen Ziffern.
		// Beispiel: spiegelzahl(123456) = 6 * 10^6 + spiegelzahl(12345)
		return (zahl%10)*zehnerpotenz(zahl) + spiegelzahl(zahl/10)
	}
}
