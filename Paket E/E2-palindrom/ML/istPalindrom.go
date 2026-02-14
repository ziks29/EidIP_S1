package main

import "strings"

// istPalindrom prüft, ob zeichenfolge ein Palindrom ist.
func istPalindrom(zeichenfolge string) bool {
	// Groß- oder Kleinschreibung soll keine Rolle spielen
	zeichenfolge = strings.ToUpper(zeichenfolge)
	// Leerstrings und Einzelzeichen sind Palindrome
	if len(zeichenfolge) <= 1 {
		return true
	} else {
		// ansonsten muss das erste und letzte Zeichen gleich sein
		// und zudem der Teil dazwischen ein Palindrom sein
		return zeichenfolge[:1] == zeichenfolge[len(zeichenfolge)-1:] &&
			istPalindrom(zeichenfolge[1:len(zeichenfolge)-1])
	}
}
