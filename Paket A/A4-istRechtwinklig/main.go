package main

import (
	"fmt"
)

// Das Programm fordert zur Eingabe von drei Seitenlängen eines Dreiecks auf
// und gibt dann aus, ob das Dreieck rechtwinklig ist.
func main() {
	// Deklaration der Seitenlängen
	var a, b, c int

	fmt.Println("Bitte geben Sie die drei ganzzahlige Seitenlängen eines Dreiecks ein:")
	_, err := fmt.Scan(&a, &b, &c)

	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}

	// prüfen, ob die Seitenlängen die Anforderungen an ein Dreieck verletzen
	// d.h. ob sie kleiner-gleich 0 sind oder die Dreiecksungleichung nicht erfüllen
	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Printf("Die Seitenlängen %d, %d und %d müssen positiv sein!\n", a, b, c)
	} else if a > (b+c) || b > (a+c) || c > (a+b) {
		fmt.Printf("Die Seitenlängen %d, %d und %d bilden kein Dreieck!\n", a, b, c)
	} else { // die Seitenlängen bilden ein Dreieck
		negation := " nicht"
		// Funktion istRechtwinklig (aus anderer Datei) aufrufen, um Rechtwinkeligkeit zu prüfen
		if istRechtwinklig(a, b, c) {
			negation = ""
		}
		fmt.Printf("Ein Dreieck mit den Seitenlängen %d, %d und %d ist%s rechtwinklig.\n", a, b, c, negation)
	}
}
