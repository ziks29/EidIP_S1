package main

import "fmt"

// Das Programm implementiert ein Buchstabensortierspiel gemäß
// der Spielregeln aus der Aufgabenstellung.
func main() {
	var original, eingabe string

	// zufälliges Wort aus Liste wählen
	original = rndWort()

	fmt.Println("Wir präsentieren Ihnen ein Wort.\n" +
		"Bitte geben Sie dieses Wort alphabetisch sortiert ein.\n" +
		"Geben Sie dabei in dem Wort mehrfach auftretende Buchstaben genauso oft wieder ein\n" +
		"und geben Sie Großbuchstaben klein ein.")
	fmt.Println()
	fmt.Println("Das Wort lautet:")
	fmt.Println(original)
	_, err := fmt.Scan(&eingabe)
	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}
	fmt.Println()

	lösung := buchstabenSortieren(original)
	if lösung == eingabe {
		fmt.Println("Gratulation! Ihre Eingabe war richtig! :-)")
	} else {
		fmt.Println("Leider falsch! :-(")
		fmt.Println("präsentiertes Wort:", original)
		fmt.Println("Ihre Eingabe:      ", eingabe)
		fmt.Println("richtige Lösung:   ", lösung)
	}
}
