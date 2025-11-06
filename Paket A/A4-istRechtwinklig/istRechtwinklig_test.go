package main

import (
	"testing"
)

func TestIstRechtwinklig(t *testing.T) {
	tests := [...]struct {
		// struktureller Aufbau eines Test-falls
		x    int
		y    int
		z    int
		want bool
	}{
		// Menge mit 18 Test-Fällen
		// ------------------------
		// Jeder Test-Fall ist wie folgt aufgebaut:
		// Die ersten drei Werte sind die Seitenlängen eines Dreiecks (also die Eingaben),
		// der vierte Wert gibt an, ob das Dreieck rechtwinklig ist (also das korrekte Ergebnis).
		{1, 2, 3, false},
		{3, 4, 5, true},
		{3, 5, 4, true},
		{5, 3, 4, true},
		{4, 3, 5, true},
		{4, 5, 3, true},
		{5, 4, 3, true},
		{4, 4, 2, false},
		{4, 2, 4, false},
		{2, 4, 4, false},
		{3, 3, 3, false},
		{6, 3, 4, false},
		{7, 5, 4, false},
		{16, 17, 2, false},
		{40, 41, 9, true},
		{160, 36, 164, true},
		{165, 52, 173, true},
		{168, 26, 170, true},
	}
	// alle Test-Fälle der Menge durchlaufen
	for _, tt := range tests {
		// Funktion istRechtwinklig mit den Eingabe-Daten des Test-Falls ausführen
		got := istRechtwinklig(tt.x, tt.y, tt.z)
		// berechnetes Ergebnis mit dem (korrekten) Ergebnis des Test-Falls vergleichen
		if got != tt.want {
			// Wenn ein berechnetes Ergebnis vom korrekten Ergebnis abweicht,
			// wird durch einen Aufruf der Funktion Errorf darüber informiert.
			t.Errorf("istRechtwinlig(%v, %v, %v) = %v, want %v", tt.x, tt.y, tt.z, got, tt.want)
		}
	}
	// Anmerkung: Die Ausführung einer Test-Datei gilt insgesamt als gescheitert (FAIL),
	//            wenn eine Funktion wie Errorf mindestens einmal aufgerufen wurde.
}
