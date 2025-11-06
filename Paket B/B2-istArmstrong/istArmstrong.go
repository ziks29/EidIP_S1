package main

// istArmstrong bestimmt, ob zahl eine Armstrong-Zahl ist.
// zahl muss eine natürliche Zahl sein.
func istArmstrong(zahl int) bool {
	digits := 0  // Zähle die Anzahl der Ziffern in zahl
	temp := zahl // Hilfsvariable
	for temp > 0 {
		digits++         // Zähle eine Ziffer
		temp = temp / 10 // Entferne die letzte Ziffer
	}

	sum := 0       // Berechne die Summe der Ziffern hoch digits
	temp = zahl    // Setze Hilfsvariable zurück
	for temp > 0 { // Solange noch Ziffern übrig sind
		digit := temp % 10               // Letzte Ziffer extrahieren
		sum += potenziere(digit, digits) // Ziffer hoch digits zur Summe hinzufügen
		temp = temp / 10                 // Letzte Ziffer entfernen
	}

	return sum == zahl // Vergleiche die Summe mit der ursprünglichen Zahl
}
