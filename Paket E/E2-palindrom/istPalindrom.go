package main

import (
	"strings"
	"unicode"
)

// istPalindrom prüft, ob zeichenfolge ein Palindrom ist.
func istPalindrom(zeichenfolge string) bool {
	// Leerzeichen und Sonderzeichen entfernen, zu Kleinbuchstaben konvertieren
	cleaned := ""
	for _, char := range zeichenfolge {
		if unicode.IsLetter(char) {
			cleaned += strings.ToLower(string(char))
		}
	}
	// Prüfe ob cleaned ein Palindrom ist
	for i := 0; i < len(cleaned)/2; i++ {
		if cleaned[i] != cleaned[len(cleaned)-1-i] {
			return false
		}
	}
	return true
}
