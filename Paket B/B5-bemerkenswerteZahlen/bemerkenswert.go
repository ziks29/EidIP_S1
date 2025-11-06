package main

// bemerkenswert gibt die nächste natürliche Zahl größer start zurück, deren
// Quadrat die Summe der natürlichen Zahlen von 1 bis n ist (mit beliebigem n).
func bemerkenswert(start int) int {
	m := start + 1
	m2 := m * m

	for {
		// Berechne die Summe der natürlichen Zahlen von 1 bis n
		// und vergleiche sie mit m2
		sum := 0
		n := 1
		for {
			sum += n
			if sum == m2 {
				return m
			}
			if sum > m2 {
				break
			}
			n++
		}
		m++
		m2 = m * m
	}
}
